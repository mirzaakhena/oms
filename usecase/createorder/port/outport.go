package port

import (
	"context"
)

// CreateOrderOutport ...
type CreateOrderOutport interface {
	SaveOrder(ctx context.Context, req SaveOrderRequest) (*SaveOrderResponse, error)
	CreatePayment(ctx context.Context, req CreatePaymentRequest) (*CreatePaymentResponse, error)
	GetAllMenuItemPrice(ctx context.Context, req GetAllMenuItemPriceRequest) (*GetAllMenuItemPriceResponse, error)
	GenerateOrderID(ctx context.Context, req GenerateOrderIDRequest) (*GenerateOrderIDResponse, error)
}

// SaveOrderRequest ...
type SaveOrderRequest struct {
	OrderID       string
	OutletCode    string
	PhoneNumber   string
	TableNumber   string
	PaymentMethod string
	OrderLine     []OrderItem
}

// SaveOrderResponse ...
type SaveOrderResponse struct {
}

// CreatePaymentRequest ...
type CreatePaymentRequest struct {
	PhoneNumber string
	OrderID     string
	TotalAmount float64
}

// CreatePaymentResponse ...
type CreatePaymentResponse struct {
	PaymentID string
}

// GetAllMenuItemPriceRequest ...
type GetAllMenuItemPriceRequest struct {
	MenuItemCodes []string
}

// GetAllMenuItemPriceResponse ...
type GetAllMenuItemPriceResponse struct {
	MenuItemWithPrices map[string]float64
}

// GenerateOrderIDRequest ...
type GenerateOrderIDRequest struct {
	OutletCode string
}

// GenerateOrderIDResponse ...
type GenerateOrderIDResponse struct {
	OrderID string
}
