package model

import (
	"github.com/mirzaakhena/oms/domain"
)

type PaymentStatus string

type validateTransitionRequest struct {
	ToStatus PaymentStatus
}

type validateTransition func(toStatus validateTransitionRequest) bool

var paymentStatusRule = map[PaymentStatus]validateTransition{
	WaitingPaymentStatus: func(req validateTransitionRequest) bool {
		if req.ToStatus == PaidPaymentStatus {
			return true
		}
		if req.ToStatus == ExpiredPaymentStatus {
			return true
		}
		if req.ToStatus == FailPaymentStatus {
			return true
		}
		return false
	},

	PaidPaymentStatus: func(req validateTransitionRequest) bool {
		if req.ToStatus == CancelledPaymentStatus {
			return true
		}
		return false
	},

	ExpiredPaymentStatus: func(req validateTransitionRequest) bool {
		return false
	},

	FailPaymentStatus: func(req validateTransitionRequest) bool {
		return false
	},

	CancelledPaymentStatus: func(req validateTransitionRequest) bool {
		return false
	},
}

const (
	WaitingPaymentStatus   = PaymentStatus("WAITING")
	PaidPaymentStatus      = PaymentStatus("PAID")
	ExpiredPaymentStatus   = PaymentStatus("EXPIRED")
	FailPaymentStatus      = PaymentStatus("FAIL")
	CancelledPaymentStatus = PaymentStatus("CANCELLED")
)

func (r PaymentStatus) ValidateNextPaymentStatus(newPaymentStatus PaymentStatusRequest) error {

	validateFunc := paymentStatusRule[r]

	isAllowed := validateFunc(validateTransitionRequest{
		ToStatus: newPaymentStatus.NewStatus,
	})

	if !isAllowed {
		return NotAllowedOrderStatusTransitionError
	}

	return nil
}

type PaymentStatusRequest struct {
	NewStatus PaymentStatus
}

const (
	PaymentStatusMustNotEmptyError         = domain.ErrorType("Status Must Not Empty")
	NotAllowedPaymentStatusTransitionError = domain.ErrorType("Not Allowed Payment Status Transition")
)
