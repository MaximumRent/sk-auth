package entity

import "gopkg.in/go-playground/validator.v9"

// Here, login is email or nickname
type LoginUserInfo struct {
	Login      string `json:"login" validate:"required"`
	Password   string `json:"password" validate:"required"`
	AuthDevice string `json:"authDevice" validate:"required"`
}

func (self *LoginUserInfo) SelfValidate() error {
	err := validator.New().Struct(self)
	return err
}
