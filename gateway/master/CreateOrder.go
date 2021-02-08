package master

import (
	"context"

	"github.com/mirzaakhena/oms/domain/model"
	"github.com/mirzaakhena/oms/infrastructure/log"
	"github.com/mirzaakhena/oms/infrastructure/util"
	"github.com/mirzaakhena/oms/usecase/createorder/port"
	"gorm.io/gorm"
)

type createOrder struct {
	DB *gorm.DB
}

// NewCreateOrderGateway ...
func NewCreateOrderGateway(DB *gorm.DB) port.CreateOrderOutport {

	return &createOrder{
		DB: DB,
	}
}

// SaveOrder ...
func (r *createOrder) SaveOrder(ctx context.Context, req port.SaveOrderRequest) (*port.SaveOrderResponse, error) {
	log.Info(ctx, "Request  %v", util.ToJSON(req))

	{
		err := r.DB.Create(&model.Order{
			ID:            req.Order.ID,
			OutletCode:    req.Order.OutletCode,
			PhoneNumber:   req.Order.PhoneNumber,
			TableNumber:   req.Order.TableNumber,
			PaymentMethod: req.Order.PaymentMethod,
		}).Error

		if err != nil {
			return nil, err
		}
	}

	{
		// orderItems := []model.OrderItem{}
		// for _, orderItem := range req.OrderLine {
		// 	orderItems = append(orderItems, model.OrderItem{
		// 		ID:           uuid.NewString(),
		// 		OrderID:      req.OrderID,
		// 		MenuItemCode: orderItem.MenuItemCode,
		// 		Quantity:     orderItem.Quantity,
		// 	})
		// }

		// err := r.DB.Create(&orderItems).Error
		// if err != nil {
		// 	return nil, err
		// }
	}

	var res port.SaveOrderResponse

	log.Info(ctx, "Response %v", util.ToJSON(res))
	return &res, nil
}

// CreatePayment ...
func (r *createOrder) CreatePayment(ctx context.Context, req port.CreatePaymentRequest) (*port.CreatePaymentResponse, error) {
	log.Info(ctx, "Request  %v", util.ToJSON(req))

	var res port.CreatePaymentResponse
	// res.PaymentID = uuid.NewString()

	// err := r.DB.Create(&model.Payment{
	// 	ID:          res.PaymentID,
	// 	PhoneNumber: req.PhoneNumber,
	// 	OrderID:     req.OrderID,
	// 	TotalAmount: req.TotalAmount,
	// }).Error

	// if err != nil {
	// 	return nil, err
	// }

	log.Info(ctx, "Response %v", util.ToJSON(res))
	return &res, nil
}

// GetAllMenuItemPrice ...
func (r *createOrder) GetAllMenuItemPrice(ctx context.Context, req port.GetAllMenuItemPriceRequest) (*port.GetAllMenuItemPriceResponse, error) {
	log.Info(ctx, "Request  %v", util.ToJSON(req))

	var menus []model.Menu
	r.DB.Where("menu_item_code IN (?)", req.MenuItemCodes).Find(&menus)

	menuItemCodeWithPrices := map[string]float64{}
	for _, menu := range menus {
		menuItemCodeWithPrices[menu.MenuItemCode] = menu.Price
	}

	var res port.GetAllMenuItemPriceResponse
	res.MenuItemWithPrices = menuItemCodeWithPrices

	log.Info(ctx, "Response %v", util.ToJSON(res))
	return &res, nil
}

// GetLatestIndexID ...
func (r *createOrder) GetLatestIndexID(ctx context.Context, req port.GetLatestIndexIDRequest) (*port.GetLatestIndexIDResponse, error) {
	log.Info(ctx, "Request  %v", util.ToJSON(req))

	var res port.GetLatestIndexIDResponse
	res.Index = 2

	log.Info(ctx, "Response %v", util.ToJSON(res))
	return &res, nil
}

// GetOrderFinishNotifyURL ...
func (r *createOrder) GetOrderFinishNotifyURL(ctx context.Context, req port.GetOrderFinishNotifyURLRequest) (*port.GetOrderFinishNotifyURLResponse, error) {
	log.Info(ctx, "Request  %v", util.ToJSON(req))

	var res port.GetOrderFinishNotifyURLResponse
	res.OrderFinishNotifyURL = "https://pleasenotifymefromhere"

	log.Info(ctx, "Response %v", util.ToJSON(res))
	return &res, nil
}
