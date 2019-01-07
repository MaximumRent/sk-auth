package errors

// Custom error for mongo operations.
// Has cause and operation type.
type MongoOperationError struct {
	cause         string
	operationType string
}

func (error *MongoOperationError) Error() string {
	return nil
}
