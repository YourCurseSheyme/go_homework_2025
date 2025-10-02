package endpoints

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/YourCurseSheyme/go_homework_2025/hw2/json"
)

type HardOpDTORequest struct {
	Code  int `json:"code"`
	Slept int `json:"slept"`
}

func HardOp(min, max time.Duration, errPct int) Handler {
	rand.Seed(time.Now().UnixNano())
	return func(writer http.ResponseWriter, request *http.Request) error {
		ms := int((max - min) / time.Millisecond)
		delay := min + time.Duration(rand.Intn(ms+1))*time.Millisecond
		select {
		case <-time.After(delay):
		case <-request.Context().Done():
			return &HttpError{Code: http.StatusRequestTimeout, Message: "request canceled"}
		}
		if errPct > 0 && rand.Intn(100) < errPct {
			return &HttpError{Code: http.StatusInternalServerError, Message: "internal error"}
		}
		json.WriteJSON(writer, http.StatusOK, HardOpDTORequest{
			Code:  http.StatusOK,
			Slept: int(delay / time.Second),
		})
		return nil
	}
}

func HardOpEndpoint(min, max time.Duration, errPct int) http.HandlerFunc {
	return MakeEndpoint(HardOp(min, max, errPct), http.MethodGet)
}
