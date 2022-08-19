package errors

import (
	"reflect"
	"testing"
)

func BenchmarkValidationError(b *testing.B) {
	type args struct {
		field   string
		message string
	}
	tests := []struct {
		name string
		args args
		want *ValidationError
	}{
		{
			name: "with-data",
			args: args{field: "test", message: "just some test"},
			want: &ValidationError{Field: "test", Message: "just some test"},
		},
		{
			name: "without-data",
			args: args{},
			want: &ValidationError{Field: "", Message: ""},
		},
	}
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			if got := Validation(tt.args.field, tt.args.message); !reflect.DeepEqual(got, tt.want) {
				b.Errorf("Validation() = %v, want %v", got, tt.want)
			}
		})
	}
}
