package entity

import "gopkg.in/go-playground/validator.v9"

type LogoutUserInfo struct {
	Nickname   string `json:"nickname" validate:"required"`
	Email      string `json:"email" validate:"required"`
	Token      string `json:"token" validate:"required"`
	AuthDevice string `json:"authDevice" validate:"required"`
}

func (self *LogoutUserInfo) SelfValidate() error {
	err := validator.New().Struct(self)
	return err
}
