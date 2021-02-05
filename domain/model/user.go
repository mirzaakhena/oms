package model

import "github.com/mirzaakhena/oms/domain"

type User struct {
	Name        string
	PhoneNumber string
	Pin         string
	UserType    string // NORMAL | PREMIUM
	Status      string // NOT_VERIFIED | ACTIVE | SUSPENDED
}

func (u *User) ValidateActivation() error {
	if u.Status != "ACTIVE" {
		return UserIsNotActive
	}

	if u.UserType != "PREMIUM" {
		return UserIsNotPremium
	}

	return nil
}

const (
	UserIsNotActive  = domain.ErrorType("User is not active")
	UserIsNotPremium = domain.ErrorType("User is not premium")
)
