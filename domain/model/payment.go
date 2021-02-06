package model

import (
	"time"

	"github.com/mirzaakhena/oms/domain"
)

type Payment struct {
	ID                   string
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

func (u *Payment) AddPaymentStatus(newStatus PaymentStatusType) error {

	if newStatus == "" {
		return PaymentStatusMustNotEmptyError
	}

	lastPaymentStatus := u.GetCurrentPaymentStatus()
	if lastPaymentStatus != nil {

		toPaid := lastPaymentStatus.Status == "WAIT" && newStatus == "PAID"
		toExpired := lastPaymentStatus.Status == "WAIT" && newStatus == "EXPIRED"
		toFail := lastPaymentStatus.Status == "WAIT" && newStatus == "FAIL"

		if !toPaid && !toExpired && !toFail {
			return NotAllowedPaymentStatusTransitionError
		}
	}

	u.PaymentStatuses = append(u.PaymentStatuses, &PaymentStatus{
		Payment: u,
		Status:  PaymentStatusType(newStatus),
		Date:    time.Now(),
	})

	return nil
}

type PaymentRequest struct {
	ID                   string
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
