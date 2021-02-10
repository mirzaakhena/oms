package model

type OrderItem struct {
	MenuItemCode string
	Quantity     int
}

type OrderItemRequest struct {
	MenuItemCode string
	Quantity     int
}
