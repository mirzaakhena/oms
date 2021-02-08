package model

import (
	"time"

	"github.com/mirzaakhena/oms/domain"
)

type Payment struct {
	ID                   string
	Date                 time.Time
	PhoneNumber          string
	OrderID              string
	TotalAmount          float64
	OrderFinishNotifyURL string
	PaymentStatuses      []*PaymentStatus
}

func NewPayment(req PaymentRequest) (*Payment, error) {

	if req.PhoneNumber == "" {
		return nil, PhoneNumberMustNotEmptyError
	}

	if req.OrderID == "" {
		return nil, PhoneNumberMustNotEmptyError
	}

	if req.TotalAmount < 0 {
		return nil, PhoneNumberMustNotEmptyError
	}

	payment := &Payment{
		ID:                   req.ID,
		Date:                 req.Date,
		PhoneNumber:          req.PhoneNumber,
		OrderID:              req.OrderID,
		TotalAmount:          req.TotalAmount,
		OrderFinishNotifyURL: req.OrderFinishNotifyURL,
	}

	err := payment.AddPaymentStatus(WaitingPaymentStatus)

	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (u *Payment) AddPaymentStatus(newStatus PaymentStatus) error {

	lastPaymentStatus := u.GetCurrentPaymentStatus()
	if lastPaymentStatus != nil {

		err := lastPaymentStatus.ValidateNextPaymentStatus(PaymentStatusRequest{NewStatus: newStatus})

		if err != nil {
			return err
		}
	}

	u.PaymentStatuses = append(u.PaymentStatuses, &newStatus)

	return nil
}

type PaymentRequest struct {
	ID                   string
	Date                 time.Time
	PhoneNumber          string
	OrderID              string
	TotalAmount          float64
	OrderFinishNotifyURL string
}

func (u *Payment) GetCurrentPaymentStatus() *PaymentStatus {
	lenPaymentStatus := len(u.PaymentStatuses)
	if lenPaymentStatus > 0 {
		return u.PaymentStatuses[lenPaymentStatus-1]
	}
	return nil
}

const (
	PhoneNumberMustNotEmptyError = domain.ErrorType("Phone Number Must Not Empty")
	OrderIDMustNotEmptyError     = domain.ErrorType("Order ID Must Not Empty")
)
