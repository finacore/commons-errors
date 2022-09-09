// Package commonserrors encapsulates some methods and data structure responsibles to organize
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
//	go get github.com/finacore/commons-errors
//
// To avoid the conflicts between packages during the imnport phase, is strongly recommended
// select other name for this package in the moment of import like shown below:
//
//	import (
//		commonserrors "github.com/finacore/commons-errors"
//	)
//
// Once imported you can use the package as you prefer. So read the function and methods
// documentation to know how to use this packagage.
package commonserrors

// DefaultError is a representaton of the most basic kind of error in this package. It's
// composed by a message description and an error code. By definition, the error code is
// not rendered when this structure is coverted to JSON.
type DefaultError struct {
	Message string `json:"message"`
	Code    int    `json:"-"`
}

// CreateDefaultError returns a DefaultError object that contains the error message and the
// error code. By default the error code associated to this error is 500, but it's can be
// changed calling the method Status(int)
//
// Usage:
//
//	err := commonserrors.CreateDefaultError("the error message goes here")
func CreateDefaultError(message string) *DefaultError {
	err := &DefaultError{
		Message: message,
	}

	return err.Status(500)
}

// MakeDefaultError returns a DefaultError object that contains the error message and the
// error code. This method uses a preexistent error  in order to capture their error message
// and calls the CreateDefaultError method.
//
// As the CreateDefaultError method, this one also set the default status as 500,  but it's
// can be changed calling the method Status(int)
//
// Usage:
//
//	err := commonserrors.MakeDefaultError(previousError)
func MakeDefaultError(err error) *DefaultError {
	return CreateDefaultError(err.Error())
}

// Error method that returns a string literal containing the error message associated to the
// data structure.
//
// Usage:
//
//	err := commonserrors.MakeDefaultError(previousError)
//	str := err.Error()
func (e *DefaultError) Error() string {
	return e.Message
}

// Status method is responsable to overwrite the default error code that has set to the data
// structure and returns the own changed data structure.
//
// Usage:
//
//	err := commonserrors.MakeDefaultError(previousError).Status(200)
func (e *DefaultError) Status(status int) *DefaultError {
	e.Code = status
	return e
}
