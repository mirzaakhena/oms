package createpayment

import (
	"context"

	"github.com/mirzaakhena/oms/domain/model"
	"github.com/mirzaakhena/oms/usecase/createpayment/port"
)

//go:generate mockery --dir port/ --name CreatePaymentOutport -output mocks/

// NewCreatePaymentUsecase ...
func NewCreatePaymentUsecase(outputPort port.CreatePaymentOutport) port.CreatePaymentInport {
	return &createPaymentInteractor{
		gateway: outputPort,
	}
}

type createPaymentInteractor struct {
	gateway port.CreatePaymentOutport
}

// Execute ...
func (r *createPaymentInteractor) Execute(ctx context.Context, req port.CreatePaymentRequest) (*port.CreatePaymentResponse, error) {

	var res port.CreatePaymentResponse

	// var latestBalance model.UserBalance
	// {
	// 	outportRes, err := r.gateway.GetLatestUserBalance(ctx, port.GetLatestUserBalanceRequest{ //
	// 		PhoneNumber: req.PhoneNumber,
	// 	})

	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	latestBalance = outportRes.UserBalance
	// }

	// {
	// 	err := latestBalance.ValidatePaymentBalanceIsEnough(req.TotalAmount)

	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }

	var user model.User
	{
		outportRes, err := r.gateway.GetUser(ctx, port.GetUserRequest{
			PhoneNumber: req.PhoneNumber,
		})

		if err != nil {
			return nil, err
		}

		user = outportRes.User
	}

	{
		err := user.ValidateUserStatus()

		if err != nil {
			return nil, err
		}
	}

	{
		outportRes, err := r.gateway.GenerateID(ctx, port.GenerateIDRequest{})

		if err != nil {
			return nil, err
		}

		res.PaymentID = outportRes.PaymentID
	}

	var lastPayment *model.Payment
	{
		outportRes, err := r.gateway.GetLastPayment(ctx, port.GetLastPaymentRequest{ //
			PhoneNumber: req.PhoneNumber,
		})

		if err != nil {
			return nil, err
		}

		lastPayment = outportRes.LastPayment
	}

	var paymentToSaved *model.Payment
	{
		payment, err := model.NewPayment(model.PaymentRequest{
			LastPayment:          lastPayment,
			ID:                   res.PaymentID,
			Date:                 req.Date,
			PhoneNumber:          req.PhoneNumber,
			OrderID:              req.OrderID,
			TotalAmount:          req.TotalAmount,
			OrderFinishNotifyURL: req.OrderFinishNotifyURL,
		})

		if err != nil {
			return nil, err
		}

		paymentToSaved = payment
	}

	{
		_, err := r.gateway.SavePayment(ctx, port.SavePaymentRequest{
			Payment: paymentToSaved,
		})

		if err != nil {
			return nil, err
		}
	}

	// var newUserBalance *model.UserBalance
	// {
	// 	userBalance, err := model.NewDeductedUserBalance(model.DeductedUserBalanceRequest{
	// 		LastUserBalance: &latestBalance,
	// 	})
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	newUserBalance = userBalance
	// }

	// {
	// 	_, err := r.gateway.SaveNewBalance(ctx, port.SaveNewBalanceRequest{
	// 		UserBalance: newUserBalance,
	// 	})
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }

	return &res, nil
}
