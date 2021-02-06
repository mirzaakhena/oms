package port

import (
	"context"
	"time"
)

// CreatePaymentInport ...
type CreatePaymentInport interface {
	Execute(ctx context.Context, req CreatePaymentRequest) (*CreatePaymentResponse, error)
}

// CreatePaymentRequest ...
type CreatePaymentRequest struct {
	PhoneNumber          string
	OrderID              string
	TotalAmount          float64
	Date                 time.Time
	OrderFinishNotifyURL string
}

// CreatePaymentResponse ...
type CreatePaymentResponse struct {
	PaymentID string
}
