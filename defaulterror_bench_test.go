package commonserrors

import (
	"errors"
	"reflect"
	"testing"
)

func BenchmarkCreateDefaultError(b *testing.B) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
		want *DefaultError
	}{
		{
			name: "with-message",
			args: args{message: "just some test"},
			want: &DefaultError{Message: "just some test"},
		},
		{
			name: "without-message",
			args: args{message: ""},
			want: &DefaultError{Message: ""},
		},
	}
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			if got := CreateDefaultError(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				b.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkMakeDefaultError(b *testing.B) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want *DefaultError
	}{
		{
			name: "with-message",
			args: args{err: errors.New("just some test")},
			want: &DefaultError{Message: "just some test"},
		},
		{
			name: "without-message",
			args: args{err: errors.New("")},
			want: &DefaultError{Message: ""},
		},
	}
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			if got := MakeDefaultError(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				b.Errorf("MakeDefaultError() = %v, want %v", got, tt.want)
			}
		})
	}
}
