package commonserrors

import (
	"fmt"
	"testing"
)

const (
	message            string = "unable to create default error"
	thisIsErrorMessage string = "this is an error message"
)

func Benchmark_CreateDefaultError(b *testing.B) {
	if got := CreateDefaultError(thisIsErrorMessage); got == nil {
		b.Errorf(message)
	}
}

func Benchmark_MakeDefaultError(b *testing.B) {
	if got := MakeDefaultError(fmt.Errorf(thisIsErrorMessage)); got == nil {
		b.Errorf(message)
	}
}

func Benchmark_Set_DefaultError_Status(b *testing.B) {
	if got := MakeDefaultError(fmt.Errorf(thisIsErrorMessage)).Status(500); got == nil {
		b.Errorf(message)
	}
}
