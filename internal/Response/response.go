package Response

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

// Server error helper.
func ServerError(logger *log.Logger, w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	logger.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// Client error response helper.
func ClientError(logger *log.Logger, w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// Not found response helper.
func NotFound(logger *log.Logger, w http.ResponseWriter) {
	ClientError(logger, w, http.StatusNotFound)
}
