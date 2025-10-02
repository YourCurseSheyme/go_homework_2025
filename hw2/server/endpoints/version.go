package endpoints

import (
	"net/http"

	"github.com/YourCurseSheyme/go_homework_2025/hw2/json"
)

var Version = "v1.0.0"

type DTOVersionResponse struct {
	Version string `json:"version"`
}

func GetVersion() Handler {
	return func(writer http.ResponseWriter, request *http.Request) error {
		json.WriteJSON(writer, http.StatusOK, DTOVersionResponse{Version: Version})
		return nil
	}
}

func VersionEndpoint() http.HandlerFunc {
	return MakeEndpoint(GetVersion(), http.MethodGet)
}
