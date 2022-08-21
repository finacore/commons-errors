package commonserrors

import (
	"errors"
	"reflect"
	"testing"
)

func TestCreateDefaultError(t *testing.T) {
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
			if got := CreateDefaultError(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateDefaultError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeDefaultError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want *DefaultError
	}{
		{
			name: "basic-error",
			args: args{err: errors.New("just some test")},
			want: &DefaultError{Message: "just some test"},
		},
		{
			name: "no-message",
			args: args{err: errors.New("")},
			want: &DefaultError{Message: ""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeDefaultError(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeDefaultError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultError_Error(t *testing.T) {
	type fields struct {
		Message string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "basic-error",
			fields: fields{Message: "This is just a test"},
			want:   "This is just a test",
		},
		{
			name:   "no-message",
			fields: fields{Message: ""},
			want:   "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &DefaultError{
				Message: tt.fields.Message,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("DefaultError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
