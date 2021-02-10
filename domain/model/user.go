package model

import "github.com/mirzaakhena/oms/shared"

type User struct {
	Name        string
	PhoneNumber string
	Pin         string
	UserType    string // NORMAL | PREMIUM
	Status      string // NOT_VERIFIED | ACTIVE | SUSPENDED
}

func (u *User) ValidateUserStatus() error {
	if u.Status != "ACTIVE" {
		return shared.UserIsNotActive
	}

	if u.UserType != "PREMIUM" {
		return shared.UserIsNotPremium
	}

	return nil
}

const ()
