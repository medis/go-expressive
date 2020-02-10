package Response

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

// Server error helper.
//func ServerError(logger *log.Logger, w http.ResponseWriter, err error) {
//	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
//	logger.Output(2, trace)
//
//	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
//}
//
//// Client error response helper.
//func ClientError(logger *log.Logger, w http.ResponseWriter, status int) {
//	http.Error(w, http.StatusText(status), status)
//}
//
//// Not found response helper.
//func NotFound(logger *log.Logger, w http.ResponseWriter) {
//	ClientError(logger, w, http.StatusNotFound)
//}

func ErrorResponse(bytes []byte, w http.ResponseWriter, statusCode int, headers map[string]string) {
	w.Header().Set("X-Content-Type-Options", "nosniff")

	// Set headers.
	for k, v := range headers {
		w.Header().Set(k, v)
	}

	// Set return status.
	w.WriteHeader(statusCode)

	_, err := w.Write(bytes)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	//logger.Output(2, trace)
	log.Output(2, trace)

	headers := make(map[string]string)
	ErrorResponse([]byte(http.StatusText(http.StatusInternalServerError)), w, http.StatusInternalServerError, headers)
}

// Client error response helper.
func ClientError(w http.ResponseWriter, status int) {
	headers := make(map[string]string)
	ErrorResponse([]byte(http.StatusText(status)), w, status, headers)
}

// Not found response helper.
func NotFound(w http.ResponseWriter) {
	ClientError(w, http.StatusNotFound)
}
