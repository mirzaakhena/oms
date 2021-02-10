package model

import (
	"fmt"
	"time"

	"github.com/mirzaakhena/oms/shared"
)

type OrderID string

func NewOrderID(req OrderIDRequest) (OrderID, error) {

	if len(req.OutletCode) != 4 {
		return "", shared.OrderIDLengthMust4Char
	}

	if req.Sequence == 0 {
		return "", shared.SequenceMustGreaterThanZero
	}

	if req.Sequence > 9999 {
		return "", shared.SequenceOutOfBound
	}

	s := fmt.Sprintf("%s%s%s%s%s%04d",
		req.OutletCode,
		req.Date.Format("01"),
		req.Date.Format("02"),
		req.PaymentMethod.OrderIDCode(),
		req.Date.Format("06"),
		req.Sequence,
	)

	return OrderID(s), nil
}

func (o OrderID) String() string {
	return string(o)
}

type OrderIDRequest struct {
	OutletCode    string
	Date          time.Time
	PaymentMethod PaymentMethodEnum
	Sequence      int
}
