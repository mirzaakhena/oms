package model

import (
	"time"

	"github.com/mirzaakhena/oms/domain"
)

type PaymentStatusType string

const (
	WaitingPaymentStatus = "WAITING"
	PaidPaymentStatus    = "PAID"
	ExpiredPaymentStatus = "EXPIRED"
	FailPaymentStatus    = "FAIL"
)

type PaymentStatus struct {
	Payment *Payment          //
	Status  PaymentStatusType // WAITING | PAID | FAIL | EXPIRED
	Date    time.Time         //
}

type PaymentStatusRequest struct {
	Payment *Payment
	Status  string
}

const (
	PaymentStatusMustNotEmptyError         = domain.ErrorType("Status Must Not Empty")
	NotAllowedPaymentStatusTransitionError = domain.ErrorType("Not Allowed Payment Status Transition")
)
