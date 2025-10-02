package endpoints

import (
	"errors"
	"net/http"

	"github.com/YourCurseSheyme/go_homework_2025/hw2/json"
)

type HttpError struct {
	Code    int
	Message string
}

func (e *HttpError) Error() string { return e.Message }

type Handler func(writer http.ResponseWriter, request *http.Request) error

func MakeEndpoint(handler Handler, method string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != method {
			json.WriteError(writer, http.StatusMethodNotAllowed, "method "+method+" not allowed")
			return
		}
		if err := handler(writer, request); err != nil {
			var yaErr *HttpError
			if errors.As(err, &yaErr) {
				json.WriteError(writer, yaErr.Code, yaErr.Message)
			} else {
				json.WriteError(writer, http.StatusInternalServerError, err.Error())
			}
		}
	}
}
