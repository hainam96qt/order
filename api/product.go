package api

import (
	"encoding/json"
	"net/http"
	"order-gokomodo/common"
	"order-gokomodo/middleware"
	"order-gokomodo/pkg/entities"
)

type CreateProductRequest struct {
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}

type CreateProductResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	SellerID    int    `json:"seller_id"`
}

func (a *APIv1) CreateProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input CreateProductRequest
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		middleware.WriteError(w, r, err, http.StatusBadRequest)
		return
	}
	userID, err := common.GetUserIDFromContext(ctx)
	if err != nil {
		middleware.WriteError(w, r, err, http.StatusBadRequest)
		return
	}

	result, err := a.ProductService.CreateProduct(ctx, &entities.CreateProductRequest{
		ProductName: input.Name,
		Description: input.Description,
		Price:       input.Price,
		SellerID:    userID,
	})
	if err != nil {
		middleware.WriteError(w, r, err, http.StatusInternalServerError)
		return
	}
	middleware.WriteResponse(ctx, w, r, &CreateProductResponse{
		ID:          result.ID,
		Name:        result.ProductName,
		Price:       result.Price,
		Description: result.Description,
		SellerID:    result.SellerID,
	})
}

type ListProductRequest struct {
	SellerID int `json:"seller_id"`
	PerPage  int `json:"per_page"`
	Page     int `json:"page"`
}

type ListProductResponse struct {
	Products   []Product `json:"products"`
	TotalCount int       `json:"total_count"`
}

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	SellerID    int    `json:"seller_id"`
}

func (a *APIv1) GetProducts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input ListProductRequest
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		middleware.WriteError(w, r, err, http.StatusBadRequest)
		return
	}

	result, err := a.ProductService.ListProduct(ctx, &entities.ListProductRequest{
		SellerID: input.SellerID,
		PerPage:  input.PerPage,
		Page:     input.Page,
	})
	if err != nil {
		middleware.WriteError(w, r, err, http.StatusInternalServerError)
		return
	}
	var products []Product
	for _, v := range result.Products {
		products = append(products, Product{
			ID:          v.ID,
			Name:        v.ProductName,
			Price:       v.Price,
			Description: v.Description,
			SellerID:    v.SellerID,
		})
	}
	middleware.WriteResponse(ctx, w, r, &ListProductResponse{
		Products:   products,
		TotalCount: len(products),
	})
}
