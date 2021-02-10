package port

import (
	"context"

	"github.com/mirzaakhena/oms/domain/model"
)

// CreatePaymentOutport ...
type CreatePaymentOutport interface {
	GetUser(ctx context.Context, req GetUserRequest) (*GetUserResponse, error)
	GenerateID(ctx context.Context, req GenerateIDRequest) (*GenerateIDResponse, error)
	SavePayment(ctx context.Context, req SavePaymentRequest) (*SavePaymentResponse, error)
	GetLastPayment(ctx context.Context, req GetLastPaymentRequest) (*GetLastPaymentResponse, error)
}

// GetUserRequest ...
type GetUserRequest struct {
	PhoneNumber string
}

// GetUserResponse ...
type GetUserResponse struct {
	User model.User
}

// GenerateIDRequest ...
type GenerateIDRequest struct {
}

// GenerateIDResponse ...
type GenerateIDResponse struct {
	PaymentID string
}

// SavePaymentRequest ...
type SavePaymentRequest struct {
	Payment *model.Payment
}

// SavePaymentResponse ...
type SavePaymentResponse struct {
}

// GetLastPaymentRequest ...
type GetLastPaymentRequest struct {
	PhoneNumber string
}

// GetLastPaymentResponse ...
type GetLastPaymentResponse struct {
	LastPayment *model.Payment
}
