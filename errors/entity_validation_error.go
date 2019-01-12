package errors

import (
	"fmt"
	"strings"
)

// Custom error for entity validation error.
// Has cause and array of invalid fields.
type EntityValidationError struct {
	invalidFields []string
}

func (error *EntityValidationError) AddInvalidField(invalidField string) {
	error.invalidFields = append(error.invalidFields, invalidField)
}

func (error *EntityValidationError) HasInvalidFields() bool {
	return ((error.invalidFields != nil) && (len(error.invalidFields) > 0))
}

func (error *EntityValidationError) Error() string {
	return fmt.Sprintf("Next fields not validated: [%s].", strings.Join(error.invalidFields[:], ","))
}
