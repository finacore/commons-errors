// This package contains encapsulate the basica errors types to helps programers to write
// low code and at the some time make a good error handling, providing a good log events
// as needed
package commonserrors

import "fmt"

// ValidationError data struct that possibilite to easy convert the error in a JSON object.
type ValidationError struct {
	Field string `json:"field"`
	DefaultError
}

// CreateValidationError function to build a new ValidatonError object based on the name of field
// and the error messge passed by paramenter.
func CreateValidationError(field string, message string) *ValidationError {
	val := &ValidationError{}

	val.Field = field
	val.Message = message
	val.Code = 400

	return val
}

// Error function to convert ValidationError into a string
func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// Status function to set the status code and return the actual stauts
func (e *ValidationError) Status(status int) *ValidationError {
	e.Code = status
	return e
}
