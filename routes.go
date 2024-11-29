package main

import (
	"github.com/gorilla/mux"
)

func setupRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/products", createProductHandler).Methods("POST")
	router.HandleFunc("/order", orderKIZHandler).Methods("POST")

	return router
}
