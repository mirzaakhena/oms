package model

import (
	"fmt"
	"strings"
	"time"

	"github.com/mirzaakhena/oms/shared"
)

type Order struct {
	ID            OrderID
	Date          time.Time
	OutletCode    string
	PhoneNumber   string
	TableNumber   string
	PaymentMethod PaymentMethod
	OrderLine     []*OrderItem
	OrderStates   []*OrderState
}

func NewOrder(req OrderRequest) (*Order, error) {

	if strings.TrimSpace(req.OutletCode) == "" {
		return nil, fmt.Errorf("OutletCode must not blank")
	}

	if strings.TrimSpace(req.TableNumber) == "" {
		return nil, fmt.Errorf("TableNumber must not blank")
	}

	if strings.TrimSpace(req.PhoneNumber) == "" {
		return nil, fmt.Errorf("PhoneNumber must not blank")
	}

	if strings.TrimSpace(req.PaymentMethod) == "" {
		return nil, fmt.Errorf("PaymentMethod must not blank")
	}

	var resultOrderID OrderID
	{

		orderID, err := NewOrderID(OrderIDRequest{
			OutletCode:    req.OutletCode,
			Date:          req.Date,
			PaymentMethod: PaymentMethod(req.PaymentMethod),
			Sequence:      req.SequenceIndex,
		})
		if err != nil {
			return nil, err
		}
		resultOrderID = orderID
	}

	var resultPaymentMethod PaymentMethod
	{
		paymentMethod, err := NewPaymentMethod(req.PaymentMethod)
		if err != nil {
			return nil, err
		}
		resultPaymentMethod = paymentMethod
	}

	var order Order
	order.Date = req.Date
	order.ID = resultOrderID
	order.OutletCode = req.OutletCode
	order.PaymentMethod = resultPaymentMethod
	order.PhoneNumber = req.PhoneNumber
	order.TableNumber = req.TableNumber

	err := order.AddOrderStatus(InitOrderState)

	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (o *Order) ValidateOrderItem() error {
	if len(o.OrderLine) == 0 {
		return shared.OrderlineMustNotEmptyError
	}
	return nil
}

func (o *Order) AddOrderItem(req OrderItemRequest) error {

	if req.MenuItemCode == "" {
		return shared.MenuItemCodeMustNotEmptyError
	}

	if req.Quantity <= 0 {
		return shared.MenuItemCodeMustNotEmptyError
	}

	o.OrderLine = append(o.OrderLine, &OrderItem{
		MenuItemCode: req.MenuItemCode,
		Quantity:     req.Quantity,
	})

	return nil
}

type OrderRequest struct {
	SequenceIndex int
	Date          time.Time
	OutletCode    string
	PhoneNumber   string
	TableNumber   string
	PaymentMethod string
}

func (o *Order) GetTotalPrice(pricePerMenu func(menuItemCode string) float64) float64 {
	totalAmount := 0.0
	for _, orderItem := range o.OrderLine {
		price := pricePerMenu(orderItem.MenuItemCode) * float64(orderItem.Quantity)
		totalAmount += price
	}
	return totalAmount
}

func (o *Order) AddOrderStatus(newState OrderStateType) error {

	if newState == "" {
		return shared.OrderStateMustNotEmptyError
	}

	lastOrderState := o.GetCurrentOrderState()
	if lastOrderState != nil {

		toPaid := lastOrderState.State == "WAIT" && newState == "PAID"
		toExpired := lastOrderState.State == "WAIT" && newState == "EXPIRED"
		toFail := lastOrderState.State == "WAIT" && newState == "FAIL"

		if !toPaid && !toExpired && !toFail {
			return shared.NotAllowedOrderStateTransitionError
		}
	}

	o.OrderStates = append(o.OrderStates, &OrderState{
		State: newState,
	})
	return nil
}

func (o *Order) GetCurrentOrderState() *OrderState {

	lenOrderStatus := len(o.OrderStates)
	if lenOrderStatus > 0 {
		return o.OrderStates[lenOrderStatus-1]
	}

	return nil
}
