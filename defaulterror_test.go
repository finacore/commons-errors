package commonserrors

import (
	"errors"
	"reflect"
	"testing"
)

const justSomeTest = "just some test"

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
			name: "basic-error-create",
			args: args{message: justSomeTest},
			want: &DefaultError{Message: justSomeTest},
		},
		{
			name: "no-message-create",
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
			name: "basic-error-make",
			args: args{err: errors.New(justSomeTest)},
			want: &DefaultError{Message: justSomeTest},
		},
		{
			name: "no-message-make",
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
			name:   "basic-error-error",
			fields: fields{Message: justSomeTest},
			want:   justSomeTest,
		},
		{
			name:   "no-message-error",
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

func TestDefaultError_Status(t *testing.T) {
	type fields struct {
		Message string
		Code    int
	}
	type args struct {
		status int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		status int
	}{
		{
			name: "default-usage",
			fields: fields{
				Message: "message not found",
				Code:    404,
			},
			args:   args{status: 404},
			status: 404,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := CreateDefaultError(tt.fields.Message).Status(tt.args.status)

			if e.Code != tt.status {
				t.Errorf("the error error: got=%d, want=%d", e.Code, tt.args.status)
			}
		})
	}
}
