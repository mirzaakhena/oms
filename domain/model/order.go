package model

import "time"

type Order struct {
	ID            string
	Date          time.Time
	OutletCode    string
	PhoneNumber   string
	TableNumber   string
	PaymentMethod string
	OrderLine     []OrderItem
}

type OrderItem struct {
	ID           string
	OrderID      string
	MenuItemCode string
	Quantity     int
}

type OrderStatus struct {
	ID      string
	OrderID string
	Date    time.Time
	Status  string
}
