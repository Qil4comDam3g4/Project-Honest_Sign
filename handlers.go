package main

import (
	"encoding/json"
	"net/http"
)

func createProductHandler(w http.ResponseWriter, r *http.Request) {
	var product Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Логика создания карточки товара в НацКаталоге Честного знака

	w.WriteHeader(http.StatusCreated)
}

func orderKIZHandler(w http.ResponseWriter, r *http.Request) {
	var order Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Логика заказа КИЗов по API Честного знака

	w.WriteHeader(http.StatusOK)
}
