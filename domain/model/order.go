package model

import (
	"fmt"
	"strings"
	"time"

	"github.com/mirzaakhena/oms/domain"
)

type Order struct {
	ID            string
	Date          time.Time
	OutletCode    string
	PhoneNumber   string
	TableNumber   string
	PaymentMethod string
	OrderLine     []*OrderItem
	OrderStatuses []*OrderStatus
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

	var order Order
	order.Date = time.Now()
	order.ID = req.OrderID
	order.OutletCode = req.OutletCode
	order.PaymentMethod = req.PaymentMethod
	order.PhoneNumber = req.PhoneNumber
	order.TableNumber = req.TableNumber

	err := order.AddOrderStatus(InitOrderStatus)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (o *Order) ValidateOrderItem() error {
	if len(o.OrderLine) == 0 {
		return OrderlineMustNotEmptyError
	}
	return nil
}

func (o *Order) AddOrderItem(req OrderItemRequest) error {

	if req.MenuItemCode == "" {
		return MenuItemCodeMustNotEmptyError
	}

	if req.Quantity <= 0 {
		return MenuItemCodeMustNotEmptyError
	}

	o.OrderLine = append(o.OrderLine, &OrderItem{
		Order:        o,
		MenuItemCode: req.MenuItemCode,
		Quantity:     req.Quantity,
	})

	return nil
}

type OrderRequest struct {
	OrderID       string
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

func (o *Order) AddOrderStatus(newStatus OrderStatusType) error {

	if newStatus == "" {
		return OrderStatusMustNotEmptyError
	}

	lastOrderStatus := o.GetCurrentOrderStatus()
	if lastOrderStatus != nil {

		toPaid := lastOrderStatus.Status == "WAIT" && newStatus == "PAID"
		toExpired := lastOrderStatus.Status == "WAIT" && newStatus == "EXPIRED"
		toFail := lastOrderStatus.Status == "WAIT" && newStatus == "FAIL"

		if !toPaid && !toExpired && !toFail {
			return NotAllowedOrderStatusTransitionError
		}
	}

	o.OrderStatuses = append(o.OrderStatuses, &OrderStatus{
		Order:  o,
		Date:   time.Now(),
		Status: string(newStatus),
	})
	return nil
}

func (o *Order) GetCurrentOrderStatus() *OrderStatus {

	lenOrderStatus := len(o.OrderStatuses)
	if lenOrderStatus > 0 {
		return o.OrderStatuses[lenOrderStatus-1]
	}

	return nil
}

const (
	MenuItemCodeMustNotEmptyError = domain.ErrorType("Menu Item Code Must Not Empty")
	QuantityMustNotEmptyError     = domain.ErrorType("Quantity Must Not Empty")
)
