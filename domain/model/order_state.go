package model

type OrderState struct {
	State OrderStateType // INIT | PAID | PREPARE | CANCEL
}

type OrderStateType string

const (
	InitOrderState    = OrderStateType("INIT")
	PaidOrderState    = OrderStateType("PAID")
	ExpiredOrderState = OrderStateType("EXPIRED")
	FailOrderState    = OrderStateType("FAIL")
)
