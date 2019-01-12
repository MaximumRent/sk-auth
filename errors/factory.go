package errors

// Here described factory methods for custom errors.

func GetMongoOperationError(cause, operationType string) error {
	mongoOperationError := new(MongoOperationError)
	mongoOperationError.cause = cause
	mongoOperationError.operationType = operationType
	return mongoOperationError
}

func GetEmptyValidationError() error {
	return new(EntityValidationError)
}
