package api

import (
	"encoding/json"
	"net/http"
	"order-gokomodo/middleware"
	"order-gokomodo/pkg/entities"
	"time"
)

type CreateOrderRequest struct {
	OrderItems []OrderItem `json:"order_products"`
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
