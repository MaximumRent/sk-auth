package errors

// Here described factory methods for custom errors.

func GetMongoOperationError(cause, operation string) error {
	return nil
}

func GetValidationError() error {
	return nil
}
