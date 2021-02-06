package port

import (
	"context"
	"time"
)

// CreateOrderInport ...
type CreateOrderInport interface {
	Execute(ctx context.Context, req CreateOrderRequest) (*CreateOrderResponse, error)
}

// CreateOrderRequest ...
type CreateOrderRequest struct {
	Date          time.Time
	OutletCode    string
	PhoneNumber   string
	TableNumber   string
	PaymentMethod string
	OrderLine     []OrderItem
}

// CreateOrderResponse ...
type CreateOrderResponse struct {
	PaymentID string
	OrderID   string
}

type OrderItem struct {
	MenuItemCode string
	Quantity     int
}
