package commonserrors

import (
	"testing"
)

func Benchmark_CreateValidationError(b *testing.B) {
	if got := CreateValidationError("some-field", "somme error"); got == nil {
		b.Errorf("the object should not be nil")
	}
}
