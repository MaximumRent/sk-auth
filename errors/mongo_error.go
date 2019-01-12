package errors

import (
	"fmt"
)

// Custom error for mongo operations.
// Has cause and operation type.
type MongoOperationError struct {
	cause         string
	operationType string
}

func (error *MongoOperationError) Error() string {
	return fmt.Sprintf("Error on operation [%s]. Cause: %s.", error.operationType, error.cause)
}
