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

	err := payment.AddPaymentStatus(AddPaymentStateRequest{
		NewState: WaitingPaymentState,
		Date:     req.Date,
	})

	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (u *Payment) AddPaymentStatus(req AddPaymentStateRequest) error {

	lastPaymentStatus := u.GetCurrentPaymentStatus()
	if lastPaymentStatus != nil {

		err := lastPaymentStatus.ValidateNextPaymentState(PaymentStateRequest{NewState: req.NewState})

		if err != nil {
			return err
		}
	}

	u.PaymentStates = append(u.PaymentStates, &PaymentState{
		State: req.NewState,
		Date:  req.Date,
	})

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

type AddPaymentStateRequest struct {
	NewState PaymentStateEnum
	Date     time.Time
}
