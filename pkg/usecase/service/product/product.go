package product

import (
	"context"
	"database/sql"
	"order-gokomodo/configs"
	"order-gokomodo/pkg/db/mysql_db"
	"order-gokomodo/pkg/entities"
	sql_model "order-gokomodo/pkg/usecase/model"
)

var _ entities.ProductService = &ProductService{}

type ProductService struct {
	DatabaseConn *sql.DB
	Query        *sql_model.Queries
}

func NewProductService(cfg *configs.Config) (*ProductService, error) {
	databaseConn, err := mysql_db.ConnectDatabase(cfg.Mysqldb)
	if err != nil {
		return nil, err
	}
	query := sql_model.New(databaseConn)
	return &ProductService{
		DatabaseConn: databaseConn,
		Query:        query,
	}, nil
}

func (p ProductService) CreateProduct(ctx context.Context, request *entities.CreateProductRequest) (*entities.CreateProductResponse, error) {
	tx, err := p.DatabaseConn.Begin()
	if err != nil {
		return nil, err
	}

	result, err := p.Query.WithTx(tx).CreateProduct(ctx, sql_model.CreateProductParams{
		Name: request.ProductName,
		Description: sql.NullString{
			String: request.Description,
			Valid:  true,
		},
		Price:    int64(request.Price),
		SellerID: int32(request.SellerID),
	})
	newProductID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	newProduct, err := p.Query.WithTx(tx).GetProductByID(ctx, int32(newProductID))
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	return &entities.CreateProductResponse{
		ID:          int(newProduct.ID),
		ProductName: newProduct.Name,
		Description: newProduct.Description.String,
		Price:       int(newProduct.Price),
		SellerID:    int(newProduct.SellerID),
	}, tx.Commit()
}

func (p ProductService) ListProduct(ctx context.Context, request *entities.ListProductRequest) (*entities.ListProductResponse, error) {
	products, err := p.Query.GetProductBySeller(ctx, sql_model.GetProductBySellerParams{
		SellerID: int32(request.SellerID),
		Limit:    int32(request.PerPage),
		Offset:   int32(request.PerPage * (request.Page - 1)),
	})
	if err != nil {
		return nil, err
	}
	var productsResponse []entities.Product
	for _, v := range products {
		productsResponse = append(productsResponse, entities.Product{
			ID:          int(v.ID),
			ProductName: v.Name,
			Description: v.Description.String,
			Price:       int(v.Price),
			SellerID:    int(v.SellerID),
		})
	}
	return &entities.ListProductResponse{
		Products:   productsResponse,
		TotalCount: len(productsResponse),
	}, nil
}
