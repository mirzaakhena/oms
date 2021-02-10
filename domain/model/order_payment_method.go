package model

import (
	"strings"

	"github.com/mirzaakhena/oms/shared"
)

type PaymentMethod string

const (
	DANAPaymentMethod  = PaymentMethod("DNA")
	GOPAYPaymentMethod = PaymentMethod("GPY")
	OVOPaymentMethod   = PaymentMethod("OVO")
)

type PaymentMethodStructure struct {
	Code string
}

var enum = map[PaymentMethod]PaymentMethodStructure{
	DANAPaymentMethod:  {Code: "N"},
	GOPAYPaymentMethod: {Code: "G"},
	OVOPaymentMethod:   {Code: "B"},
}

func NewPaymentMethod(name string) (PaymentMethod, error) {
	name = strings.ToUpper(name)

	_, exist := enum[PaymentMethod(name)]
	if !exist {
		return "", shared.UnrecognizedPaymentMethod
	}

	return PaymentMethod(name), nil
}

func (r PaymentMethod) OrderIDCode() string {
	return enum[r].Code
}
