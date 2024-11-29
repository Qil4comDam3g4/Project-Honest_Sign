package main

import (
	"fmt"
)

func createProductInNationalCatalog(product Product) error {
	// Взаимодействие с API Честного знака для создания карточки товара
	fmt.Printf("Creating product in National Catalog: %s\n", product.GTIN)
	return nil
}

func orderKIZs(order Order) error {
	// Логика заказа КИЗов через API
	fmt.Printf("Ordering KIZs for user: %s\n", order.UserID)
	return nil
}
