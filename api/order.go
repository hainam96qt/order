package api

import (
	"encoding/json"
	"net/http"
	"order-gokomodo/middleware"
	"order-gokomodo/pkg/entities"
	"time"
)

type CreateOrderRequest struct {
	OrderItems []OrderItem `json:"order_items"`
}

type OrderItem struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type OrderProduct struct {
	OrderProductID int    `json:"order_product_id"`
	ProductID      int    `json:"product_id"`
	ProductName    string `json:"product_name"`
	Quantity       int    `json:"quantity"`
}

type CreateOrderResponse struct {
	OrderID                    int
	DeliverySourceAddress      string
	DeliveryDestinationAddress string
	BuyerID                    int
	SellerID                   int
	CreatedAt                  time.Time
	UpdatedAt                  time.Time
	TotalPrice                 int
	Status                     string
	Products                   []OrderProduct
}

func (a *APIv1) CreateOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input CreateOrderRequest
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		middleware.WriteError(w, r, err, http.StatusBadRequest)
		return
	}

	var orderItems []entities.OrderItem
	for _, v := range input.OrderItems {
		orderItems = append(orderItems, entities.OrderItem{
			ProductID: v.ProductID,
			Quantity:  v.Quantity,
		})
	}
	result, err := a.OrderService.CreateOrder(ctx, &entities.CreateOrderRequest{OrderItems: orderItems})
	if err != nil {
		middleware.WriteError(w, r, err, http.StatusInternalServerError)
		return
	}
	var orderProduct []OrderProduct
	for _, v := range result.Products {
		orderProduct = append(orderProduct, OrderProduct{
			OrderProductID: v.OrderProductID,
			ProductID:      v.ProductID,
			ProductName:    v.ProductName,
			Quantity:       v.Quantity,
		})
	}
	middleware.WriteResponse(ctx, w, r, &CreateOrderResponse{
		OrderID:                    result.OrderID,
		DeliverySourceAddress:      result.DeliverySourceAddress,
		DeliveryDestinationAddress: result.DeliveryDestinationAddress,
		BuyerID:                    result.BuyerID,
		SellerID:                   result.SellerID,
		CreatedAt:                  result.CreatedAt,
		UpdatedAt:                  result.UpdatedAt,
		TotalPrice:                 result.TotalPrice,
		Status:                     result.Status,
		Products:                   orderProduct,
	})
}

// ListOrder
type ListOrderRequest struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
}

type ListOrderResponse struct {
	Orders []Order `json:"orders"`
}

type Order struct {
	OrderID                    int            `json:"order_id"`
	DeliverySourceAddress      string         `json:"delivery_source_address"`
	DeliveryDestinationAddress string         `json:"delivery_destination_address"`
	BuyerID                    int            `json:"buyer_id"`
	SellerID                   int            `json:"seller_id"`
	CreatedAt                  time.Time      `json:"created_at"`
	UpdatedAt                  time.Time      `json:"updated_at"`
	TotalPrice                 int            `json:"total_price"`
	Status                     string         `json:"status"`
	Products                   []OrderProduct `json:"products"`
}

func (a *APIv1) ListOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input ListOrderRequest
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		middleware.WriteError(w, r, err, http.StatusBadRequest)
		return
	}
	result, err := a.OrderService.ListOrder(ctx, &entities.ListOrderRequest{
		Page:    input.Page,
		PerPage: input.PerPage,
	})
	if err != nil {
		middleware.WriteError(w, r, err, http.StatusInternalServerError)
		return
	}
	var orders []Order
	for _, order := range result.Orders {
		var orderProducts []OrderProduct
		for _, orderProduct := range order.Products {
			orderProducts = append(orderProducts, OrderProduct{
				OrderProductID: orderProduct.OrderProductID,
				ProductID:      orderProduct.ProductID,
				ProductName:    orderProduct.ProductName,
				Quantity:       orderProduct.Quantity,
			})
		}
		orders = append(orders, Order{
			OrderID:                    order.OrderID,
			DeliverySourceAddress:      order.DeliverySourceAddress,
			DeliveryDestinationAddress: order.DeliveryDestinationAddress,
			BuyerID:                    order.BuyerID,
			SellerID:                   order.SellerID,
			CreatedAt:                  order.CreatedAt,
			UpdatedAt:                  order.UpdatedAt,
			TotalPrice:                 order.TotalPrice,
			Status:                     order.Status,
			Products:                   orderProducts,
		})
	}

	middleware.WriteResponse(ctx, w, r, &ListOrderResponse{
		Orders: orders,
	})
}

type AcceptOrderRequest struct {
	OrderID int `json:"order_id"`
}

type AcceptOrderResponse struct {
	Success bool `json:"success"`
}

func (a *APIv1) AcceptOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input AcceptOrderRequest
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		middleware.WriteError(w, r, err, http.StatusBadRequest)
		return
	}
	result, err := a.OrderService.AcceptOrder(ctx, input.OrderID)
	if err != nil {
		middleware.WriteError(w, r, err, http.StatusInternalServerError)
		return
	}
	middleware.WriteResponse(ctx, w, r, &AcceptOrderResponse{
		Success: result,
	})
}
