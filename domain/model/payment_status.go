package model

import (
	"github.com/mirzaakhena/oms/domain"
)

type PaymentStatusType string

const (
	WaitingPaymentStatus = PaymentStatusType("WAITING")
	PaidPaymentStatus    = PaymentStatusType("PAID")
	ExpiredPaymentStatus = PaymentStatusType("EXPIRED")
	FailPaymentStatus    = PaymentStatusType("FAIL")
)

type PaymentStatus struct {
	Status PaymentStatusType // WAITING | PAID | FAIL | EXPIRED
}

type PaymentStatusRequest struct {
	Status string
}

const (
	PaymentStatusMustNotEmptyError         = domain.ErrorType("Status Must Not Empty")
	NotAllowedPaymentStatusTransitionError = domain.ErrorType("Not Allowed Payment Status Transition")
)
