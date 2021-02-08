package createorder

import (
	"context"
	"testing"
	"time"

	"github.com/mirzaakhena/oms/domain/model"
	"github.com/mirzaakhena/oms/usecase/createorder/mocks"
	"github.com/mirzaakhena/oms/usecase/createorder/port"
	"github.com/stretchr/testify/assert"
)

func Test_CreateOrder_Normal(t *testing.T) {

	ctx := context.Background()

	date := time.Date(2021, time.December, 11, 12, 12, 12, 12, time.UTC)

	outputPort := mocks.CreateOrderOutport{}
	{
		call := outputPort.On("SaveOrder", ctx, port.SaveOrderRequest{ //
			Order: &model.Order{
				ID:            "02081211N210002",
				Date:          date,
				OutletCode:    "0208",
				PhoneNumber:   "08123",
				TableNumber:   "B32",
				PaymentMethod: "DNA",
				OrderLine: []*model.OrderItem{
					{MenuItemCode: "101", Quantity: 2},
					{MenuItemCode: "102", Quantity: 3},
				},
				OrderStatuses: []*model.OrderStatus{
					{Status: model.InitOrderStatus},
				},
			},
		})
		call.Return(&port.SaveOrderResponse{ //
		}, nil)
	}
	{
		call := outputPort.On("CreatePayment", ctx, port.CreatePaymentRequest{ //
			PhoneNumber:          "08123",
			OrderID:              "02081211N210002",
			TotalAmount:          16000,
			OrderFinishNotifyURL: "http://notifyme.com",
		})
		call.Return(&port.CreatePaymentResponse{ //
			PaymentID: "12345",
		}, nil)
	}
	{
		call := outputPort.On("GetAllMenuItemPrice", ctx, port.GetAllMenuItemPriceRequest{ //
			MenuItemCodes: []string{"101", "102"},
		})
		call.Return(&port.GetAllMenuItemPriceResponse{ //
			MenuItemWithPrices: map[string]float64{"101": 2000, "102": 4000},
		}, nil)
	}
	{
		call := outputPort.On("GetLatestIndexID", ctx, port.GetLatestIndexIDRequest{ //
			OutletCode: "0208",
			Date:       date,
		})
		call.Return(&port.GetLatestIndexIDResponse{ //
			Index: 2,
		}, nil)
	}
	{
		call := outputPort.On("GetOrderFinishNotifyURL", ctx, port.GetOrderFinishNotifyURLRequest{ //
		})
		call.Return(&port.GetOrderFinishNotifyURLResponse{ //
			OrderFinishNotifyURL: "http://notifyme.com",
		}, nil)
	}

	res, err := NewCreateOrderUsecase(&outputPort).Execute(ctx, port.CreateOrderRequest{ //
		Date:          date,
		OutletCode:    "0208",
		PhoneNumber:   "08123",
		TableNumber:   "B32",
		PaymentMethod: "DNA",
		OrderLine: []port.OrderItem{
			{MenuItemCode: "101", Quantity: 2},
			{MenuItemCode: "102", Quantity: 3},
		},
	})

	assert.Nil(t, err)

	assert.Equal(t, &port.CreateOrderResponse{ //
		PaymentID: "12345",
		OrderID:   "02081211N210002",
	}, res)

}
