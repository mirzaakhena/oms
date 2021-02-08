package port

import (
	"context"
	"time"

	"github.com/mirzaakhena/oms/domain/model"
)

// CreateOrderOutport ...
type CreateOrderOutport interface {
	SaveOrder(ctx context.Context, req SaveOrderRequest) (*SaveOrderResponse, error)
	CreatePayment(ctx context.Context, req CreatePaymentRequest) (*CreatePaymentResponse, error)
	GetAllMenuItemPrice(ctx context.Context, req GetAllMenuItemPriceRequest) (*GetAllMenuItemPriceResponse, error)
	GetLatestIndexID(ctx context.Context, req GetLatestIndexIDRequest) (*GetLatestIndexIDResponse, error)
	GetOrderFinishNotifyURL(ctx context.Context, req GetOrderFinishNotifyURLRequest) (*GetOrderFinishNotifyURLResponse, error)
}

// SaveOrderRequest ...
type SaveOrderRequest struct {
	Order *model.Order
}

// SaveOrderResponse ...
type SaveOrderResponse struct {
}

// CreatePaymentRequest ...
type CreatePaymentRequest struct {
	PhoneNumber          string
	OrderID              string
	TotalAmount          float64
	OrderFinishNotifyURL string
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

// GetLatestIndexIDRequest ...
type GetLatestIndexIDRequest struct {
	OutletCode string
	Date       time.Time
}

// GetLatestIndexIDResponse ...
type GetLatestIndexIDResponse struct {
	Index int
}

// GetOrderFinishNotifyURLRequest ...
type GetOrderFinishNotifyURLRequest struct {
}

// GetOrderFinishNotifyURLResponse ...
type GetOrderFinishNotifyURLResponse struct {
	OrderFinishNotifyURL string
}
