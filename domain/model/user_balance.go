package model

import (
	"time"

	"github.com/mirzaakhena/oms/shared"
)

type UserBalance struct {
	User        *User
	Date        time.Time
	Amount      float64
	Balance     float64
	Description string
}

func NewDeductedUserBalance(req DeductedUserBalanceRequest) (*UserBalance, error) {

	if req.Amount <= 0 {
		return nil, shared.AmountMustGreaterThanZeroError
	}

	newBalance := req.LastUserBalance.Balance - req.Amount
	if newBalance < 0 {
		return nil, shared.BalanceIsNotEnoughError
	}

	return &UserBalance{
		Date:        req.Date,
		User:        req.User,
		Amount:      req.Amount,
		Balance:     newBalance,
		Description: req.Description,
	}, nil
}

type DeductedUserBalanceRequest struct {
	LastUserBalance *UserBalance
	User            *User
	Date            time.Time
	Amount          float64
	Description     string
}

func (u *UserBalance) ValidatePaymentBalanceIsEnough(amount float64) error {
	if u.Balance-amount < 0 {
		return shared.BalanceIsNotEnoughError
	}
	return nil
}
