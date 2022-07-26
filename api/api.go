package api

import (
	"log"
	"order-gokomodo/configs"
	"order-gokomodo/middleware"
	"order-gokomodo/pkg/usecase/service/auth"
	"order-gokomodo/pkg/usecase/service/order"
	product "order-gokomodo/pkg/usecase/service/product"

	"github.com/gorilla/mux"
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

	v1Router := router.PathPrefix("/v1").Subrouter()
	v1Router.Use(middleware.Authentication)
	v1Router.HandleFunc("/createProduct", a.CreateProduct)
	v1Router.HandleFunc("/getProducts", a.GetProducts)

	v1Router.Use(middleware.Authentication)
	v1Router.HandleFunc("/createOrder", a.CreateOrder)
	v1Router.HandleFunc("/getOrders", a.ListOrder)
	v1Router.HandleFunc("/acceptOrder", a.AcceptOrder)

	return router
}
