package cerr

import (
	"fmt"
	"testing"
)

func Benchmark_CreateValidationError(b *testing.B) {
	if got := CreateValidationError("some-field", "somme error"); got == nil {
		b.Errorf("the object should not be nil")
	}
}

func Benchmark_Set_ValidationError_Status(b *testing.B) {
	if got := MakeDefaultError(fmt.Errorf("this is an error message")).Status(500); got == nil {
		b.Errorf("unable to create default error")
	}
}
