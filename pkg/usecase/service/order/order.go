package order

import (
	"context"
	"database/sql"
	"errors"
	"order-gokomodo/common"
	"order-gokomodo/configs"
	"order-gokomodo/pkg/db/mysql_db"
	"order-gokomodo/pkg/entities"
	sql_model "order-gokomodo/pkg/usecase/model"
)

var _ entities.OrderService = &OrderService{}

type OrderService struct {
	DatabaseConn *sql.DB
	Query        *sql_model.Queries
}

func NewOrderService(cfg *configs.Config) (*OrderService, error) {
	databaseConn, err := mysql_db.ConnectDatabase(cfg.Mysqldb)
	if err != nil {
		return nil, err
	}
	query := sql_model.New(databaseConn)
	return &OrderService{
		DatabaseConn: databaseConn,
		Query:        query,
	}, nil
}

func (o OrderService) CreateOrder(ctx context.Context, request *entities.CreateOrderRequest) (*entities.CreateOrderResponse, error) {
	tx, err := o.DatabaseConn.Begin()
	if err != nil {
		return nil, err
	}

	// get info product
	sellerID := 0
	totalPrice := 0
	var productByIDs = make(map[int]entities.Product)
	var productIDs []int
	for _, v := range request.OrderItems {
		productIDs = append(productIDs, v.ProductID)
	}
	productDbs, err := o.Query.WithTx(tx).GetProductsByIDs(ctx, productIDs)
	if err != nil {
		return nil, err
	}
	for _, v := range productDbs {
		sellerID = int(v.SellerID)
		productByIDs[int(v.ID)] = entities.Product{
			ID:          int(v.ID),
			ProductName: v.Name,
			Description: v.Description.String,
			Price:       int(v.Price),
			SellerID:    int(v.SellerID),
		}
	}
	for _, v := range request.OrderItems {
		totalPrice += productByIDs[v.ProductID].Price * v.Quantity
	}

	// verify product
	isDifference := false
	for _, v := range productByIDs {
		if v.SellerID != sellerID {
			isDifference = true
		}
	}
	if isDifference {
		return nil, errors.New("Only one seller in one order")
	}

	if len(productByIDs) != len(request.OrderItems) {
		return nil, errors.New("Some product not permit to order")
	}

	userID, err := common.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	// get info user
	buyer, err := o.Query.WithTx(tx).GetUserByID(ctx, int32(userID))
	if err != nil {
		return nil, err
	}

	seller, err := o.Query.WithTx(tx).GetUserByID(ctx, int32(userID))
	if err != nil {
		return nil, err
	}

	// create order
	result, err := o.Query.WithTx(tx).CreateOrder(ctx, sql_model.CreateOrderParams{
		DeliverySourceAddress:      seller.Address,
		DeliveryDestinationAddress: buyer.Address,
		BuyerID:                    int32(userID),
		SellerID:                   int32(sellerID),
		Status:                     "pending",
		TotalPrice:                 int64(totalPrice),
	})
	if err != nil {
		return nil, err
	}
	newOrderID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	newOrder, err := o.Query.WithTx(tx).GetOrderByID(ctx, int32(newOrderID))
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}
	// create order product
	var products []entities.OrderProduct
	for _, v := range request.OrderItems {
		_, err := o.Query.WithTx(tx).CreateOrderProduct(ctx, sql_model.CreateOrderProductParams{
			OrderID:   int32(newOrderID),
			ProductID: int32(v.ProductID),
			Quantity:  int32(v.Quantity),
		})
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		products = append(products, entities.OrderProduct{
			OrderProductID: int(newOrderID),
			ProductID:      v.ProductID,
			ProductName:    productByIDs[int(v.ProductID)].ProductName,
			Quantity:       v.Quantity,
		})
	}

	// get them back
	return &entities.CreateOrderResponse{
		OrderID:                    int(newOrder.ID),
		DeliverySourceAddress:      newOrder.DeliverySourceAddress,
		DeliveryDestinationAddress: newOrder.DeliveryDestinationAddress,
		BuyerID:                    int(newOrder.BuyerID),
		SellerID:                   int(newOrder.SellerID),
		CreatedAt:                  newOrder.CreatedAt.Time,
		UpdatedAt:                  newOrder.UpdatedAt.Time,
		TotalPrice:                 int(newOrder.TotalPrice),
		Status:                     string(newOrder.Status),
		Products:                   products,
	}, tx.Commit()
}

func (o OrderService) ListOrder(ctx context.Context, request *entities.ListOrderRequest) (*entities.ListOrderResponse, error) {
	userID, err := common.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	role, err := common.GetUserRoleFromContext(ctx)
	if err != nil {
		return nil, err
	}
	var orders []sql_model.Order
	if role == "seller" {
		orders, err = o.Query.GetOrderBySellerID(ctx, sql_model.GetOrderBySellerIDParams{
			SellerID: int32(userID),
			Limit:    int32(request.PerPage),
			Offset:   int32(request.PerPage * (request.Page - 1)),
		})
		if err != nil {
			return nil, err
		}
	} else {
		orders, err = o.Query.GetOrderByBuyerID(ctx, sql_model.GetOrderByBuyerIDParams{
			BuyerID: int32(userID),
			Limit:   int32(request.PerPage),
			Offset:  int32(request.PerPage * (request.Page - 1)),
		})
		if err != nil {
			return nil, err
		}
	}
	var ordersIDs []int
	for _, v := range orders {
		ordersIDs = append(ordersIDs, int(v.ID))
	}
	orderProducts, err := o.Query.GetOrderProductByOrderIDs(ctx, ordersIDs)
	if err != nil {
		return nil, err
	}

	var productIDs []int
	for _, v := range orderProducts {
		productIDs = append(productIDs, int(v.ProductID))
	}
	products, err := o.Query.GetProductsByIDs(ctx, productIDs)
	if err != nil {
		return nil, err
	}
	var mapProductBuyIDs = make(map[int]entities.Product)
	for _, v := range products {
		mapProductBuyIDs[int(v.ID)] = entities.Product{
			ID:          int(v.ID),
			ProductName: v.Name,
			Description: v.Description.String,
			Price:       int(v.Price),
			SellerID:    int(v.SellerID),
		}
	}
	var mapOrderProductsBuyOrderID = make(map[int][]entities.OrderProduct)
	for _, v := range orderProducts {
		mapOrderProductsBuyOrderID[int(v.OrderID)] = append(mapOrderProductsBuyOrderID[int(v.OrderID)], entities.OrderProduct{
			OrderProductID: int(v.ID),
			ProductID:      int(v.ProductID),
			ProductName:    mapProductBuyIDs[int(v.ProductID)].ProductName,
			Quantity:       int(v.Quantity),
		})
	}
	var result entities.ListOrderResponse
	for _, v := range orders {
		result.Orders = append(result.Orders, entities.Order{
			OrderID:                    int(v.ID),
			DeliverySourceAddress:      v.DeliverySourceAddress,
			DeliveryDestinationAddress: v.DeliveryDestinationAddress,
			BuyerID:                    int(v.BuyerID),
			SellerID:                   int(v.SellerID),
			CreatedAt:                  v.CreatedAt.Time,
			UpdatedAt:                  v.UpdatedAt.Time,
			TotalPrice:                 int(v.TotalPrice),
			Status:                     string(v.Status),
			Products:                   mapOrderProductsBuyOrderID[int(v.ID)],
		})
	}
	return &result, nil
}

func (o OrderService) AcceptOrder(ctx context.Context, orderID int) (bool, error) {
	userID, err := common.GetUserIDFromContext(ctx)
	if err != nil {
		return false, err
	}
	role, err := common.GetUserRoleFromContext(ctx)
	if err != nil {
		return false, err
	}
	if role != "seller" {
		return false, errors.New("Only seller can accept order")
	}
	err = o.Query.UpdateOrderStatus(ctx, sql_model.UpdateOrderStatusParams{
		SellerID: int32(userID),
		ID:       int32(orderID),
	})
	if err != nil {
		return false, err
	}
	return true, nil
}
