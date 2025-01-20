package errorz

import "fmt"

type APIError struct {
	StatusCode int
	Message    string
}

func (e *APIError) Error() string {
	return e.Message
}

func NewAPIError(statusCode int, message string) *APIError {
	return &APIError{
		StatusCode: statusCode,
		Message:    message,
	}
}

func (a *APIError) JoinError(message string) error {
	if a == nil {
		return nil
	}
	if a.Error() == "" {
		return fmt.Errorf("%v%w", message, a)
	}
	return fmt.Errorf("%v %w", message, a)
}
