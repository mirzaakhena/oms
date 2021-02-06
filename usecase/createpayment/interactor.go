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
		err := user.ValidateUserAbilityToPay()

		if err != nil {
			return nil, err
		}
	}

	var latestBalance model.UserBalance
	{
		outportRes, err := r.gateway.GetLatestUserBalance(ctx, port.GetLatestUserBalanceRequest{ //
			PhoneNumber: req.PhoneNumber,
		})

		if err != nil {
			return nil, err
		}

		latestBalance = outportRes.UserBalance
	}

	{
		err := latestBalance.ValidatePaymentBalanceIsEnough(req.TotalAmount)
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

	var paymentToStored *model.Payment
	{
		payment, err := model.NewPayment(model.PaymentRequest{
			ID:          res.PaymentID,
			PhoneNumber: req.PhoneNumber,
			OrderID:     req.OrderID,
			TotalAmount: req.TotalAmount,
		})

		if err != nil {
			return nil, err
		}

		paymentToStored = payment
	}

	{
		err := paymentToStored.AddPaymentStatus(model.WaitingPaymentStatus)

		if err != nil {
			return nil, err
		}
	}

	{
		_, err := r.gateway.SavePayment(ctx, port.SavePaymentRequest{
			Payment: paymentToStored,
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
