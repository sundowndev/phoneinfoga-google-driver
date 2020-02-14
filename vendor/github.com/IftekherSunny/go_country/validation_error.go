package country

// ValidationError struct
type ValidationError struct {
	message string
}

// implementing error method
func (v ValidationError) Error() string {
	return v.message
}

// NewValidationError will create a new validation error class
func NewValidationError(code string) error {
	ve := &ValidationError{}
	ve.message = "The country ISO 3166-1 Alpha-2 code [ " + code + " ] does not exists."

	return ve
}
