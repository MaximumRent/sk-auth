package errors

// Custom error for entity validation error.
// Has cause and array of invalid fields.
type EntityValidationError struct {
	cause         string
	invalidFields []string
}

func (error *EntityValidationError) AddInvalidField(invalidField string) {
	error.invalidFields = append(error.invalidFields, invalidField)
}

func (error *EntityValidationError) Error() string {
	return ""
}
