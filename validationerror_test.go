package commonserrors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateValidationError(t *testing.T) {
	assert := assert.New(t)

	field := "fieldname"
	message := "the error message"

	err := CreateValidationError(field, message)
	assert.NotNil(err)

	assert.Equal(err.Field, field)
	assert.Equal(err.Message, message)
	assert.Equal(err.Code, 400)
}

func TestValidationError_Error(t *testing.T) {
	assert := assert.New(t)

	field := "fieldname"
	message := "the error message"

	err := CreateValidationError(field, message)
	assert.NotNil(err)

	strError := err.Error()
	assert.NotEmpty(strError)
	assert.Equal(strError, fmt.Sprintf("%s: %s", field, message))
}

func TestValidationError_Status(t *testing.T) {
	assert := assert.New(t)

	field := "fieldname"
	message := "the error message"

	err := CreateValidationError(field, message)
	assert.NotNil(err)

	err.Status(500)
	assert.Equal(err.Code, 500)
}

func Benchmark_Set_ValidationError_Status(b *testing.B) {
	if got := MakeDefaultError(fmt.Errorf("this is an error message")).Status(500); got == nil {
		b.Errorf("unable to create default error")
	}
}
