// Package cerr encapsulates some methods and data structure responsibles to organize
// and manage the erros obtained duruing the program execution.
//
// Among the characteristics, the one that stand out are the capability to convert the data
// structure to an string, following the error interface implemented by the golang error, As
// well the agregation of an integer error code that simplify the error recognitions less the
// necessity of make reflections in the received object.
//
// The first step to use this package is to get it through the command below, then you will be
// able to import it in your code.
//
//	go get github.com/gsdenys/cerr
//
// To avoid the conflicts between packages during the imnport phase, is strongly recommended
// select other name for this package in the moment of import like shown below:
//
//	import (
//		"github.com/gsdenys/cerr"
//	)
//
// Once imported you can use the package as you prefer. So read the function and methods
// documentation to know how to use this packagage.
package cerr

import "fmt"

// ValidationError is a representaton of the error that coming from the validation library like
// go-playground/validator. It's composed by the DefaultError fields inheritance and the field
// name
type ValidationError struct {
	Field string `json:"field"`
	DefaultError
}

// CreateValidationError function to build a new ValidatonError object based on the name of field
// and the error messge passed by paramenter.

// CreateValidationError returns a CreateValidationError object that contains the field name, the
// error message and the error code. By default the error code associated to this error is 400,
// but it's can be changed calling the method Status(int)
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

// Error method that returns a string literal containing the composition of the field name and the
// message separated by ': '
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

// Status method is responsable to overwrite the default error code that has set to the data
// structure and returns the own changed data structure.
//
// Usage:
//
//	err := cerr.CreateValidationError("name", "the name should not be empty").Status(200)
func (e *ValidationError) Status(status int) *ValidationError {
	e.Code = status
	return e
}
