package errs

import "net/http"

type ErrMessage interface {
	Message() string
	Status() int
	Error() string
}

type errMessage struct {
	ErrMessage string `json:"message"`
	ErrStatus  int    `json:"status"`
	ErrError   string `json:"error"`
}

func (e *errMessage) Message() string {
	return e.ErrMessage
}

func (e *errMessage) Status() int {
	return e.ErrStatus
}

func (e *errMessage) Error() string {
	return e.ErrError
}

func NewBadRequest(message string) *errMessage {
	return &errMessage{
		ErrMessage: message,
		ErrStatus:  http.StatusBadRequest,
		ErrError: "BAD_REQUEST",
	}
}

func NewInternalServerError(message string) *errMessage {
	return &errMessage{
		ErrMessage: message,
		ErrStatus: http.StatusInternalServerError,
		ErrError: "INTERNAL_SERVER_ERROR",
	}
}

func NewUnprocessableEntityError(message string) *errMessage {
	return &errMessage{
		ErrMessage: message,
		ErrStatus: http.StatusUnprocessableEntity,
		ErrError: "INVALID_REQUEST_BODY",
	}
}
