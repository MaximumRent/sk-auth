package entity

import (
	"gopkg.in/go-playground/validator.v9"
)

// Short variation of user.
// Defined only with email, nickname and token
type ShortUser struct {
	Email string 	`json:"email" validate:"required"`
	Token string 	`json:"token" validate:"required"`
	Nickname string `json:"nickname" validate:"required"`
}

func (self *ShortUser) SelfValidate() error {
	err := validator.New().Struct(self)
	return err
}

// Here, login is email or nickname
type LoginUserInfo struct {
	Login 		string 	`json:"login" validate:"required"`
	Password 	string	`json:"password" validate:"required"`
	AuthDevice 	string	`json:"authDevice" validate:"required"`
}

func (self *LoginUserInfo) SelfValidate() error {
	err := validator.New().Struct(self)
	return err
}
