package errs

import (
	"fmt"
	"net/http"
)

// HTTPError is an error that will be rendered to the client.
type HTTPError struct {
	Code int

	Message          string
	ValidationErrors []ValidationError `json:"validation_errors"`
}

// New initializes new HTTPError
func New(err error, code int) *HTTPError {
	if code <= 0 {
		code = http.StatusInternalServerError
	}

	e := &HTTPError{
		Code: code,
	}

	if err != nil {
		e.Message = err.Error()
	}

	return e
}

// Location sets HTTPError Location
func (e *HTTPError) Msg(msg string) *HTTPError {
	e.Message = msg
	return e
}

// Error implements the error interface
func (e *HTTPError) Error() string {
	return fmt.Sprintf("handler: %v: %v ", e.Code, e.Message)
}

// Validations appends validation errors to HTTPError.ValidationErrors
func (e *HTTPError) Validations(validations ...ValidationError) *HTTPError {
	e.ValidationErrors = append(e.ValidationErrors, validations...)
	return e
}
