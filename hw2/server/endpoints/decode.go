package endpoints

import (
	"encoding/base64"
	"net/http"

	"github.com/YourCurseSheyme/go_homework_2025/hw2/json"
)

type DTODecodeRequest struct {
	InputString string `json:"inputString"`
}

type DTODecodeResponse struct {
	OutputString string `json:"outputString"`
}

func DecodeString() Handler {
	return func(writer http.ResponseWriter, request *http.Request) error {
		var jsonRequest DTODecodeRequest
		if err := json.ReadJSON(request, &jsonRequest); err != nil {
			return err
		}
		input := jsonRequest.InputString
		if input == "" {
			return &HttpError{Code: http.StatusBadRequest, Message: "input cannot be empty"}
		}
		raw, err := base64.StdEncoding.DecodeString(input)
		if err != nil {
			return &HttpError{Code: http.StatusBadRequest, Message: err.Error()}
		}
		json.WriteJSON(writer, http.StatusOK, DTODecodeResponse{string(raw)})
		return nil
	}
}

func DecodeEndpoint() http.HandlerFunc {
	return MakeEndpoint(DecodeString(), http.MethodPost)
}
