package entity

import "gopkg.in/go-playground/validator.v9"

type AddRoleRequest struct {
	RoleName	string	`json:"role_name" validate:"required"`
}

func (self *AddRoleRequest) SelfValidate() error {
	err := validator.New().Struct(self)
	return err
}
