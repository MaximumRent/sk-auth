package validation

// Interface for self validatable entities.
// Self validation checks that entity has valid field value and so on.
type SelfValidatable interface {
	SelfValidate() bool
}
