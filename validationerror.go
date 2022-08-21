package commonserrors

import "fmt"

// ValidationError data struct that possibilite to easy convert the error in a JSON object.
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// CreateValidationError function to build a new ValidatonError object based on the name of field
// and the error messge passed by paramenter.
func CreateValidationError(field string, message string) *ValidationError {
	return &ValidationError{
		Field:   field,
		Message: message,
	}
}

// Error function to convert ValidationError into a string
func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}
