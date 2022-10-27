package cerr

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	errorMessage string = "the error message"
	fieldname    string = "fieldname"
)

func TestCreateValidationError(t *testing.T) {
	assert := assert.New(t)

	field := fieldname
	message := errorMessage
	err := CreateValidationError(field, message)
	assert.NotNil(err)

	assert.Equal(err.Field, field)
	assert.Equal(err.Message, message)
	assert.Equal(err.Code, 400)
}

func TestValidationError_Error(t *testing.T) {
	assert := assert.New(t)

	field := fieldname
	message := errorMessage

	err := CreateValidationError(field, message)
	assert.NotNil(err)

	strError := err.Error()
	assert.NotEmpty(strError)
	assert.Equal(strError, fmt.Sprintf("%s: %s", field, message))
}

func TestValidationError_Status(t *testing.T) {
	assert := assert.New(t)

	field := fieldname
	message := errorMessage

	err := CreateValidationError(field, message)
	assert.NotNil(err)

	retError := err.Status(500)

	assert.NotNil(retError)
	assert.Equal(retError, err)

	assert.Equal(err.Code, 500)
}
