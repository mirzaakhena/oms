package model

import (
	"strings"

	"github.com/mirzaakhena/oms/shared"
)

type PaymentMethodEnum string

const (
	DANAPaymentMethodEnum  PaymentMethodEnum = "DNA"
	GOPAYPaymentMethodEnum PaymentMethodEnum = "GPY"
	OVOPaymentMethodEnum   PaymentMethodEnum = "OVO"
)

type enumDetail struct {
	OrderIDCode string
}

var enumMapStructure = map[PaymentMethodEnum]enumDetail{
	DANAPaymentMethodEnum:  {OrderIDCode: "N"},
	GOPAYPaymentMethodEnum: {OrderIDCode: "G"},
	OVOPaymentMethodEnum:   {OrderIDCode: "B"},
}

func NewPaymentMethodEnum(name string) (PaymentMethodEnum, error) {
	name = strings.ToUpper(name)

	_, exist := enumMapStructure[PaymentMethodEnum(name)]
	if !exist {
		return "", shared.UnrecognizedPaymentMethod
	}

	return PaymentMethodEnum(name), nil
}

func (r PaymentMethodEnum) OrderIDCode() string {
	return enumMapStructure[r].OrderIDCode
}
