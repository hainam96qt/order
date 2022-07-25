package entities

import "context"

type ProductService interface {
	CreateProduct(context.Context, *CreateProductRequest) (*CreateProductResponse, error)
	ListProduct(context.Context, *ListProductRequest) (*ListProductResponse, error)
}

type Product struct {
	ID          int
	ProductName string
	Description string
	Price       int
	SellerID    int
}

type CreateProductRequest struct {
	ProductName string
	Description string
	Price       int
	SellerID    int
}

type CreateProductResponse struct {
	ID          int
	ProductName string
	Description string
	Price       int
	SellerID    int
}

type ListProductRequest struct {
	SellerID int
	PerPage  int
	Page     int
}

type ListProductResponse struct {
	Products   []Product
	TotalCount int
}
