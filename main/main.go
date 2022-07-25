package main

import (
	"log"
	"net/http"
	"order-gokomodo/api"
	"order-gokomodo/configs"

	"github.com/gorilla/mux"
)

func main() {
	cfg, err := configs.NewConfig()
	if err != nil {
		log.Print(err)
	}

	router := mux.NewRouter()
	apiv1 := api.NewAPIv1(cfg)

	apiv1.AttachHandlers(router)

	_ = http.ListenAndServe(":8081", router)
}
