package main

type User struct {
	ID      string
	Name    string
	Inn     string
	IsAdmin bool
}

type Product struct {
	GTIN     string
	Quantity int
}

type Order struct {
	UserID   string
	Products []Product
}
