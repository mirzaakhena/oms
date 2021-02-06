package model

import "github.com/mirzaakhena/oms/domain"

type OrderItem struct {
	Order        *Order
	MenuItemCode string
	Quantity     int
}

type OrderItemRequest struct {
	MenuItemCode string
	Quantity     int
}

const (
	OrderlineMustNotEmptyError = domain.ErrorType("Orderline must not empty")
)
