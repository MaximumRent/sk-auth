package entity

import "gopkg.in/go-playground/validator.v9"

type AccessRequest struct {
	Path	string	`json:"path" validate:"required"`
}

func (self *AccessRequest) SelfValidate() error {
	err := validator.New().Struct(self)
	return err
}