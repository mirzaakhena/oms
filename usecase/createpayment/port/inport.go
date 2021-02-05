package port

import (
	"context"
)

// CreatePaymentInport ...
type CreatePaymentInport interface {
	Execute(ctx context.Context, req CreatePaymentRequest) (*CreatePaymentResponse, error)
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
