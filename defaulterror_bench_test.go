package commonserrors

import (
	"fmt"
	"testing"
)

func BenchmarkCreateDefaultError(b *testing.B) {
	if got := CreateDefaultError("this is an error message"); got == nil {
		b.Errorf("unable to create default error")
	}
}

func BenchmarkMakeDefaultError(b *testing.B) {
	if got := MakeDefaultError(fmt.Errorf("this is an error message")); got == nil {
		b.Errorf("unable to create default error")
	}
}

func BenchmarkStatus(b *testing.B) {
	if got := MakeDefaultError(fmt.Errorf("this is an error message")).Status(500); got == nil {
		b.Errorf("unable to create default error")
	}
}
