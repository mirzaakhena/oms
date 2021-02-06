package model

import (
	"github.com/mirzaakhena/oms/domain"
)

type OrderStatus struct {
	Status OrderStatusType // INIT | PAID | PREPARE | CANCEL
}

type OrderStatusType string

const (
	InitOrderStatus    = OrderStatusType("INIT")
	PaidOrderStatus    = OrderStatusType("PAID")
	ExpiredOrderStatus = OrderStatusType("EXPIRED")
	FailOrderStatus    = OrderStatusType("FAIL")
)

const (
	OrderStatusMustNotEmptyError         = domain.ErrorType("Order Status Must Not Empty")
	NotAllowedOrderStatusTransitionError = domain.ErrorType("Not Allowed Order Status Transition")
)
