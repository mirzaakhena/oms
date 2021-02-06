package createpayment

import (
	"context"
	"testing"
	"time"

	"github.com/mirzaakhena/oms/domain/model"
	"github.com/mirzaakhena/oms/usecase/createpayment/mocks"
	"github.com/mirzaakhena/oms/usecase/createpayment/port"
	"github.com/stretchr/testify/assert"
)

func Test_CreatePayment_Normal(t *testing.T) {

	ctx := context.Background()

	user := model.User{
		Name:        "Mirza",
		PhoneNumber: "081321",
		Pin:         "123456",
		UserType:    "PREMIUM",
		Status:      "ACTIVE",
	}

	date := time.Now()

	outputPort := mocks.CreatePaymentOutport{}
	{
		call := outputPort.On("GetUser", ctx, port.GetUserRequest{ //
			PhoneNumber: "08123",
		})
		call.Return(&port.GetUserResponse{ //
			User: user,
		}, nil)
	}
	{
		call := outputPort.On("GetLatestUserBalance", ctx, port.GetLatestUserBalanceRequest{ //
			PhoneNumber: "08123",
		})
		call.Return(&port.GetLatestUserBalanceResponse{ //
			UserBalance: model.UserBalance{
				User:        &user,
				Date:        date,
				Amount:      10000,
				Balance:     50000,
				Description: "12345",
			},
		}, nil)
	}
	{
		call := outputPort.On("GenerateID", ctx, port.GenerateIDRequest{ //
		})
		call.Return(&port.GenerateIDResponse{ //
			PaymentID: "4567",
		}, nil)
	}
	{
		call := outputPort.On("SavePayment", ctx, port.SavePaymentRequest{ //
			Payment: &model.Payment{
				ID:                   "4567",
				Date:                 date,
				PhoneNumber:          "08123",
				OrderID:              "12345",
				TotalAmount:          500,
				OrderFinishNotifyURL: "https://notifyme.com",
				PaymentStatuses: []*model.PaymentStatus{
					{Status: model.WaitingPaymentStatus},
				},
			},
		})
		call.Return(&port.SavePaymentResponse{ //
		}, nil)
	}

	res, err := NewCreatePaymentUsecase(&outputPort).Execute(ctx, port.CreatePaymentRequest{ //
		PhoneNumber:          "08123",
		OrderID:              "12345",
		TotalAmount:          500,
		Date:                 date,
		OrderFinishNotifyURL: "https://notifyme.com",
	})

	assert.Nil(t, err)

	assert.Equal(t, &port.CreatePaymentResponse{ //
		PaymentID: "4567",
	}, res)

}
