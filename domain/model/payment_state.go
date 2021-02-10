package model

import (
	"time"

	"github.com/mirzaakhena/oms/shared"
)

type PaymentStateEnum string

const (
	WaitingPaymentState   = PaymentStateEnum("WAITING")
	PaidPaymentState      = PaymentStateEnum("PAID")
	ExpiredPaymentState   = PaymentStateEnum("EXPIRED")
	FailPaymentState      = PaymentStateEnum("FAIL")
	CancelledPaymentState = PaymentStateEnum("CANCELLED")
)

type PaymentState struct {
	State PaymentStateEnum
	Date  time.Time
}

type validateTransitionRequest struct {
	ToState PaymentStateEnum
}

type validateTransitionType func(toState validateTransitionRequest) bool

var paymentStateRule = map[PaymentStateEnum]validateTransitionType{

	WaitingPaymentState: func(req validateTransitionRequest) bool {

		switch req.ToState {

		case PaidPaymentState:
			return true

		case ExpiredPaymentState:
			return true

		case FailPaymentState:
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

func NewPaymentState(req PaymentStateRequest) *PaymentState {
	return &PaymentState{
		State: req.NewState,
		Date:  req.Date,
	}
}

func (r PaymentState) TransitTo(req PaymentStateRequest) (*PaymentState, error) {

	validateTransitionFunc := paymentStateRule[r.State]

	isAllowed := validateTransitionFunc(validateTransitionRequest{
		ToState: req.NewState,
	})

	if !isAllowed {
		return nil, shared.NotAllowedOrderStateTransitionError.Var(r, req.NewState)
	}

	if r.Date.After(req.Date) {
		return nil, shared.InvalidDateOrderStateTransitionError.Var(r, req.NewState)
	}

	return NewPaymentState(req), nil
}

type PaymentStateRequest struct {
	NewState PaymentStateEnum
	Date     time.Time
}
