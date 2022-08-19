package baseerror

type DefaultError struct {
	Message string `json:"message"`
}

//CreateDefaultError function to generate new error message
func Default(message string) *DefaultError {
	return &DefaultError{
		Message: message,
	}
}
