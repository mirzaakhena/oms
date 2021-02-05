package createorder

import (
	"context"
	"fmt"
	"strings"

	"github.com/mirzaakhena/oms/usecase/createorder/port"
)

//go:generate mockery --dir port/ --name CreateOrderOutport -output mocks/

// NewCreateOrderUsecase ...
func NewCreateOrderUsecase(outputPort port.CreateOrderOutport) port.CreateOrderInport {
	return &createOrderInteractor{
		gateway: outputPort,
	}
}

type createOrderInteractor struct {
	gateway port.CreateOrderOutport
}

// Execute ...
func (r *createOrderInteractor) Execute(ctx context.Context, req port.CreateOrderRequest) (*port.CreateOrderResponse, error) {

	if strings.TrimSpace(req.OutletCode) == "" {
		return nil, fmt.Errorf("OutletCode must not blank")
	}

	if strings.TrimSpace(req.PhoneNumber) == "" {
		return nil, fmt.Errorf("PhoneNumber must not blank")
	}

	if strings.TrimSpace(req.PaymentMethod) == "" {
		return nil, fmt.Errorf("PaymentMethod must not blank")
	}

	if len(req.OrderLine) == 0 {
		return nil, fmt.Errorf("OrderLine must not empty")
	}

	var res port.CreateOrderResponse

	orderID := ""
	{
		resOutport, err := r.gateway.GenerateOrderID(ctx, port.GenerateOrderIDRequest{ //
			OutletCode: req.OutletCode,
		})

		if err != nil {
			return nil, err
		}

		res.OrderID = resOutport.OrderID

	}

	{
		_, err := r.gateway.SaveOrder(ctx, port.SaveOrderRequest{ //
			OrderID:       orderID,
			OutletCode:    req.OutletCode,
			PhoneNumber:   req.PhoneNumber,
			TableNumber:   req.TableNumber,
			PaymentMethod: req.PaymentMethod,
			OrderLine:     req.OrderLine,
		})

		if err != nil {
			return nil, err
		}

	}

	menuItemCodes := []string{}
	for _, orderItem := range req.OrderLine {
		menuItemCodes = append(menuItemCodes, orderItem.MenuItemCode)
	}

	var menuItemCodeWithPrices map[string]float64

	{
		resOutport, err := r.gateway.GetAllMenuItemPrice(ctx, port.GetAllMenuItemPriceRequest{ //
			MenuItemCodes: menuItemCodes,
		})

		if err != nil {
			return nil, err
		}

		menuItemCodeWithPrices = resOutport.MenuItemWithPrices
	}

	totalAmount := 0.0
	for _, orderItem := range req.OrderLine {
		price := menuItemCodeWithPrices[orderItem.MenuItemCode] * float64(orderItem.Quantity)
		totalAmount += price
	}

	{
		resOutport, err := r.gateway.CreatePayment(ctx, port.CreatePaymentRequest{ //
			PhoneNumber: req.PhoneNumber,
			TotalAmount: totalAmount,
			OrderID:     res.OrderID,
		})

		if err != nil {
			return nil, err
		}

		res.PaymentID = resOutport.PaymentID
	}

	return &res, nil
}
