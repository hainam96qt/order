package api

import (
	"challenge/pkg/mod/github.com/gorilla/mux"
	"log"
	"order-gokomodo/configs"
	"order-gokomodo/middleware"
	"order-gokomodo/pkg/usecase/service/auth"
	"order-gokomodo/pkg/usecase/service/order"
	product "order-gokomodo/pkg/usecase/service/product"
)

type APIv1 struct {
	AuthService    *auth.AuthenticationService
	ProductService *product.ProductService
	OrderService   *order.OrderService
}

func NewAPIv1(cfg *configs.Config) *APIv1 {
	authService, err := auth.NewAuthenticationService(cfg)
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}

	productService, err := product.NewProductService(cfg)
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}

	orderService, err := order.NewOrderService(cfg)
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}
	return &APIv1{
		AuthService:    authService,
		ProductService: productService,
		OrderService:   orderService,
	}
}

func (a *APIv1) AttachHandlers(router *mux.Router) *mux.Router {
	router.HandleFunc("/login", a.Login)
	router.HandleFunc("/register", a.Register)

	productRouter := router.PathPrefix("/product").Subrouter()
	productRouter.Use(middleware.Authentication)
	productRouter.HandleFunc("/create", a.CreateProduct)
	productRouter.HandleFunc("/gets", a.GetProducts)

	orderRouter := router.PathPrefix("/order").Subrouter()
	orderRouter.Use(middleware.Authentication)
	orderRouter.HandleFunc("/create", a.CreateProduct)
	orderRouter.HandleFunc("/gets", a.CreateProduct)
	orderRouter.HandleFunc("/accept", a.CreateProduct)

	return router
}
