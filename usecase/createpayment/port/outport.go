package port

import (
	"context"

	"github.com/mirzaakhena/oms/domain/model"
)

// CreatePaymentOutport ...
type CreatePaymentOutport interface {
	GetUser(ctx context.Context, req GetUserRequest) (*GetUserResponse, error)
	GetLatestUserBalance(ctx context.Context, req GetLatestUserBalanceRequest) (*GetLatestUserBalanceResponse, error)
	GenerateID(ctx context.Context, req GenerateIDRequest) (*GenerateIDResponse, error)
	SavePayment(ctx context.Context, req SavePaymentRequest) (*SavePaymentResponse, error)
}

// GetUserRequest ...
type GetUserRequest struct {
	PhoneNumber string
}

// GetUserResponse ...
type GetUserResponse struct {
	User model.User
}

// GetLatestUserBalanceRequest ...
type GetLatestUserBalanceRequest struct {
	PhoneNumber string
}

// GetLatestUserBalanceResponse ...
type GetLatestUserBalanceResponse struct {
	UserBalance model.UserBalance
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
