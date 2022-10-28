package cerr

import "fmt"

// ValidationError is a representation of the error that coming from the
// validation library like go-playground/validator. It's composed by the
// DefaultError fields inheritance and the field name
type ValidationError struct {
	Field string `json:"field"`
	DefaultError
}

// CreateValidationError function to build a new ValidationError object based on
// the name of field and the error message passed by parameter.

// CreateValidationError returns a CreateValidationError object that contains the
// field name, the error message and the error code. By default the error code
// associated to this error is 400, but it's can be changed calling the method
// Status(int)
//
// Usage:
//
//	err := cerr.CreateValidationError("name", "the name should not be empty")
func CreateValidationError(field string, message string) *ValidationError {
	val := &ValidationError{}

	val.Field = field
	val.Message = message
	val.Code = 400

	return val
}

// Error method that returns a string literal containing the composition of the
// field name and the message separated by ':'
//
// e.g. name: the name should not be empty
//
// Usage:
//
//	err := cerr.CreateValidationError("name", "the name should not be empty")
//	str := err.Error()
func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// Status method is responsable to overwrite the default error code that has set
// to the data structure and returns the own changed data structure.
//
// Usage:
//
//	err := cerr.CreateValidationError("name", "the name should not be empty")
//	err.Status(200)
func (e *ValidationError) Status(status int) *ValidationError {
	e.Code = status
	return e
}
