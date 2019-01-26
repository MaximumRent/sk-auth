package errors

import "fmt"

type RabbitMqError struct {
	message string
}

func GetRabbitMqError(message string) *RabbitMqError {
	return &RabbitMqError{message: message}
}

func (error *RabbitMqError) Error() string {
	return fmt.Sprintf("%s", error.message)
}
