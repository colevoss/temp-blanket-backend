package response

import "net/http"

type ErrorResponse struct {
	Message string      `json:"message"`
	Code    int         `json:"-"`
	Err     error       `json:"-"`
	Data    interface{} `json:"data"`
}

func (err *ErrorResponse) Error() string {
	return err.Message
}

func NewError(message string, code int, data interface{}, err error) error {
	return &ErrorResponse{
		Message: message,
		Code:    code,
		Err:     err,
	}
}

func BadRequest(message string, data interface{}, err error) error {
	return NewError(message, http.StatusBadRequest, data, err)
}

func NotFound(msg string, data interface{}, err error) error {
	return NewError(msg, http.StatusNotFound, data, err)
}
