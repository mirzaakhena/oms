package model

import (
	"strings"

	"github.com/mirzaakhena/oms/domain"
)

type PaymentMethod string

var enum = map[string]string{
	"DNA": "N",
	"GPY": "G",
	"OVO": "B",
}

func NewPaymentMethod(name string) (PaymentMethod, error) {
	name = strings.ToUpper(name)

	_, exist := enum[name]
	if !exist {
		return "", UnrecognizedPaymentMethod
	}

	return PaymentMethod(name), nil
}

func (r PaymentMethod) OrderIDCode() string {
	return enum[string(r)]
}

const (
	UnrecognizedPaymentMethod = domain.ErrorType("Unrecognized payment method")
)
