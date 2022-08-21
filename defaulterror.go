package commonserrors

// DefaultError ç
type DefaultError struct {
	Message string `json:"message"`
}

// CreateDefaultError function to build a new DefaultError object based on a simple string
func CreateDefaultError(message string) *DefaultError {
	return &DefaultError{
		Message: message,
	}
}

// MakeDefaultError function to build a new DefaultError object based on a pre existend error.
func MakeDefaultError(err error) *DefaultError {
	return CreateDefaultError(err.Error())
}

// Error function to convert DefaultError into a string
func (e *DefaultError) Error() string {
	return e.Message
}
