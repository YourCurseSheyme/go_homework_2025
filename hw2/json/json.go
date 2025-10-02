package json

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
)

const MaxBodySize int64 = 1 << 20

func WriteJSON(writer http.ResponseWriter, status int, value any) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	writer.WriteHeader(status)
	_ = json.NewEncoder(writer).Encode(value)
}

func WriteError(writer http.ResponseWriter, status int, message string) {
	type errorResponse struct {
		Error string `json:"error"`
	}
	WriteJSON(writer, status, errorResponse{Error: message})
}

func ReadJSON(request *http.Request, destination any) error {
	if request.Body == nil {
		return http.ErrBodyNotAllowed
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error closing body:", err)
		}
	}(request.Body)
	request.Body = http.MaxBytesReader(nil, request.Body, MaxBodySize)
	dec := json.NewDecoder(request.Body)
	return dec.Decode(destination)
}

func PrintJSON(writer io.Writer, data []byte) {
	var anyJSON any
	if err := json.Unmarshal(data, &anyJSON); err != nil {
		_, _ = fmt.Fprintln(writer, "Error:", err)
		return
	}
	printMap := func(m map[string]any) {
		keys := make([]string, 0, len(m))
		for key := range m {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			_, _ = fmt.Fprintf(writer, "%s: %v\n", key, m[key])
		}
	}
	switch v := anyJSON.(type) {
	case map[string]any:
		printMap(v)
	case []any:
		for i, it := range v {
			_, _ = fmt.Fprintf(writer, "[%d]\n", i)
			if m, ok := it.(map[string]any); ok {
				printMap(m)
			} else {
				_, _ = fmt.Fprintf(writer, "%v\n", it)
			}
		}
	default:
		_, _ = fmt.Fprintf(writer, "%v\n", v)
	}
}
