package entity

import (
	"gopkg.in/go-playground/validator.v9"
)

// Short variation of user.
// Defined only with email, nickname and token
type ShortUser struct {
	Email    string `json:"email" validate:"required"`
	Token    string `json:"token" validate:"required"`
	Nickname string `json:"nickname" validate:"required"`
}

func (self *ShortUser) SelfValidate() error {
	err := validator.New().Struct(self)
	return err
}
