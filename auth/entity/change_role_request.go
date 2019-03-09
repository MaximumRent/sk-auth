package entity

import "gopkg.in/go-playground/validator.v9"

type ChangeRoleRequest struct {
	RoleName	string	`json:"role_name" validate:"required"`
}

func (self *ChangeRoleRequest) SelfValidate() error {
	err := validator.New().Struct(self)
	return err
}
