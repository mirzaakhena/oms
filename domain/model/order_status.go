package model

import (
	"time"

	"github.com/mirzaakhena/oms/domain"
)

type OrderStatus struct {
	Order  *Order
	Date   time.Time
	Status string // INIT | PAID | PREPARE | CANCEL
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
