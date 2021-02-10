package model

import (
	"time"

	"github.com/mirzaakhena/oms/shared"
)

type PaymentStateEnum string

type PaymentState struct {
	State PaymentStateEnum
	Date  time.Time
}

type validateTransitionRequest struct {
	ToState PaymentStateEnum
}

type validateTransition func(toState validateTransitionRequest) bool

var paymentStateRule = map[PaymentStateEnum]validateTransition{
	WaitingPaymentState: func(req validateTransitionRequest) bool {
		if req.ToState == PaidPaymentState {
			return true
		}
		if req.ToState == ExpiredPaymentState {
			return true
		}
		if req.ToState == FailPaymentState {
			return true
		}
		return false
	},

	PaidPaymentState: func(req validateTransitionRequest) bool {
		return req.ToState == CancelledPaymentState
	},

	ExpiredPaymentState: func(req validateTransitionRequest) bool {
		return false
	},

	FailPaymentState: func(req validateTransitionRequest) bool {
		return false
	},

	CancelledPaymentState: func(req validateTransitionRequest) bool {
		return false
	},
}

const (
	WaitingPaymentState   = PaymentStateEnum("WAITING")
	PaidPaymentState      = PaymentStateEnum("PAID")
	ExpiredPaymentState   = PaymentStateEnum("EXPIRED")
	FailPaymentState      = PaymentStateEnum("FAIL")
	CancelledPaymentState = PaymentStateEnum("CANCELLED")
)

func (r PaymentState) ValidateNextPaymentState(newPaymentState PaymentStateRequest) error {

	validateTransitionFunc := paymentStateRule[r.State]

	isAllowed := validateTransitionFunc(validateTransitionRequest{
		ToState: newPaymentState.NewState,
	})

	if !isAllowed {
		return shared.NotAllowedOrderStateTransitionError.Var(r, newPaymentState.NewState)
	}

	return nil
}

type PaymentStateRequest struct {
	NewState PaymentStateEnum
}
