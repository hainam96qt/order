package entities

import (
	"context"
	"time"
)

type OrderService interface {
	ListOrder(context.Context, *ListOrderRequest) (*ListOrderResponse, error)
	CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error)
	AcceptOrder(ctx context.Context, orderID int) (bool, error)
}

type ListOrderRequest struct {
	Page    int
	PerPage int
}

type ListOrderResponse struct {
	Orders []Order
}

type Order struct {
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

type CreateOrderRequest struct {
	OrderItems []OrderItem
}

type OrderItem struct {
	ProductID int
	Quantity  int
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

type OrderProduct struct {
	OrderProductID int
	ProductID      int
	ProductName    string
	Quantity       int
}

type AcceptOrderRequest struct {
	OrderID int
}
