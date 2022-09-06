package commonserrors

import (
	"fmt"
	"testing"
)

func Benchmark_CreateDefaultError(b *testing.B) {
	if got := CreateDefaultError("this is an error message"); got == nil {
		b.Errorf("unable to create default error")
	}
}

func Benchmark_MakeDefaultError(b *testing.B) {
	if got := MakeDefaultError(fmt.Errorf("this is an error message")); got == nil {
		b.Errorf("unable to create default error")
	}
}

func Benchmark_Status(b *testing.B) {
	if got := MakeDefaultError(fmt.Errorf("this is an error message")).Status(500); got == nil {
		b.Errorf("unable to create default error")
	}
}
