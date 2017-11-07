package errors

import "net/http"
import "errors"

//HTTPError Base http error
type HTTPError struct {
	Code    int
	Message string
	Name    string
}

func (e HTTPError) Error() string {
	return e.Message
}

// New init a std error
func New(text string) error {
	return errors.New(text)
}

//NewHTTPError init a new http error
func NewHTTPError(name string, code int, message string) *HTTPError {
	return &HTTPError{code, message, name}
}

//BadRequestError 400 Bad Request
func BadRequestError() *HTTPError {
	return NewHTTPError("BadRequest", http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
}

//UnauthorizedError 401 Unauthorized
func UnauthorizedError() *HTTPError {
	return NewHTTPError("Unauthorized", http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
}

//InternalServerError 500 Internal Server Error
func InternalServerError(err error) *HTTPError {
	msg := http.StatusText(http.StatusInternalServerError)
	if err != nil {
		msg = err.Error()
	}
	return NewHTTPError("InternalServer", http.StatusInternalServerError, msg)
}
