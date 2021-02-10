package model

import (
	"time"

	"github.com/mirzaakhena/oms/shared"
)

type Payment struct {
	ID                   string
	Date                 time.Time
	PhoneNumber          string
	OrderID              string
	TotalAmount          float64
	OrderFinishNotifyURL string
	PaymentStates        []*PaymentState
}

func NewPayment(req PaymentRequest) (*Payment, error) {

	if req.PhoneNumber == "" {
		return nil, shared.PhoneNumberMustNotEmptyError
	}

	if req.OrderID == "" {
		return nil, shared.PhoneNumberMustNotEmptyError
	}

	if req.Date.After(time.Now()) {
		return nil, shared.InvalidDateError
	}

	if req.LastPayment != nil {
		if req.LastPayment.Date.After(req.Date) {
			return nil, shared.InvalidDateError
		}
	}

	if req.TotalAmount < 0 {
		return nil, shared.PhoneNumberMustNotEmptyError
	}

	payment := &Payment{
		ID:                   req.ID,
		Date:                 req.Date,
		PhoneNumber:          req.PhoneNumber,
		OrderID:              req.OrderID,
		TotalAmount:          req.TotalAmount,
		OrderFinishNotifyURL: req.OrderFinishNotifyURL,
	}

	payment.PaymentStates = append(payment.PaymentStates, &PaymentState{
		State: WaitingPaymentState,
		Date:  req.Date,
	})

	return payment, nil
}

func (u *Payment) AddPaymentState(req PaymentStateRequest) error {

	newPaymentState, err := u.GetCurrentPaymentStatus().TransitTo(req)

	if err != nil {
		return err
	}

	u.PaymentStates = append(u.PaymentStates, newPaymentState)

	return nil
}

type PaymentRequest struct {
	LastPayment          *Payment
	ID                   string
	Date                 time.Time
	PhoneNumber          string
	OrderID              string
	TotalAmount          float64
	OrderFinishNotifyURL string
}

func (u *Payment) GetCurrentPaymentStatus() *PaymentState {
	lenPaymentStatus := len(u.PaymentStates)
	if lenPaymentStatus > 0 {
		return u.PaymentStates[lenPaymentStatus-1]
	}
	return nil
}
