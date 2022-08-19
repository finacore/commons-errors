package baseerror

import (
	"reflect"
	"testing"
)

func TestDefaultError(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
		want *DefaultError
	}{
		{
			name: "basic-error",
			args: args{message: "just some test"},
			want: &DefaultError{Message: "just some test"},
		},
		{
			name: "no-message",
			args: args{message: ""},
			want: &DefaultError{Message: ""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Default(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
