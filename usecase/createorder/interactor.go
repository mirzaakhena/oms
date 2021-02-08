package createorder

import (
	"context"

	"github.com/mirzaakhena/oms/domain/model"
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

	var res port.CreateOrderResponse

	var sequenceIndex int
	{
		resOutport, err := r.gateway.GetLatestIndexID(ctx, port.GetLatestIndexIDRequest{ //
			OutletCode: req.OutletCode,
			Date:       req.Date,
		})

		if err != nil {
			return nil, err
		}

		sequenceIndex = resOutport.Index
	}

	var orderToSave *model.Order
	{
		order, err := model.NewOrder(model.OrderRequest{
			Date:          req.Date,
			SequenceIndex: sequenceIndex,
			OutletCode:    req.OutletCode,
			PhoneNumber:   req.PhoneNumber,
			TableNumber:   req.TableNumber,
			PaymentMethod: req.PaymentMethod,
		})

		if err != nil {
			return nil, err
		}

		orderToSave = order
		res.OrderID = order.ID.String()
	}

	{
		for _, orderItem := range req.OrderLine {
			err := orderToSave.AddOrderItem(model.OrderItemRequest{
				MenuItemCode: orderItem.MenuItemCode,
				Quantity:     orderItem.Quantity,
			})

			if err != nil {
				return nil, err
			}

		}
	}

	{
		err := orderToSave.ValidateOrderItem()

		if err != nil {
			return nil, err
		}
	}

	var menuItemCodeWithPrices map[string]float64
	{

		menuItemCodes := []string{}
		for _, orderItem := range req.OrderLine {
			menuItemCodes = append(menuItemCodes, orderItem.MenuItemCode)
		}

		resOutport, err := r.gateway.GetAllMenuItemPrice(ctx, port.GetAllMenuItemPriceRequest{ //
			MenuItemCodes: menuItemCodes,
		})

		if err != nil {
			return nil, err
		}

		menuItemCodeWithPrices = resOutport.MenuItemWithPrices
	}

	totalPrice := orderToSave.GetTotalPrice(func(menuItemCode string) float64 {
		return menuItemCodeWithPrices[menuItemCode]
	})

	{
		_, err := r.gateway.SaveOrder(ctx, port.SaveOrderRequest{ //
			Order: orderToSave,
		})

		if err != nil {
			return nil, err
		}
	}

	var orderFinishNotifyURL string
	{
		outportRes, err := r.gateway.GetOrderFinishNotifyURL(ctx, port.GetOrderFinishNotifyURLRequest{})

		if err != nil {
			return nil, err
		}

		orderFinishNotifyURL = outportRes.OrderFinishNotifyURL
	}

	{
		resOutport, err := r.gateway.CreatePayment(ctx, port.CreatePaymentRequest{ //
			PhoneNumber:          req.PhoneNumber,
			TotalAmount:          totalPrice,
			OrderID:              res.OrderID,
			OrderFinishNotifyURL: orderFinishNotifyURL,
		})

		if err != nil {
			return nil, err
		}

		res.PaymentID = resOutport.PaymentID
	}

	return &res, nil
}
