package baseerror

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func Validation(field string, message string) *ValidationError {
	return &ValidationError{
		Field:   field,
		Message: message,
	}
}
