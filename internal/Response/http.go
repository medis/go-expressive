package Response

import (
	"bytes"
	"net/http"
)

// w io.Writer
func HttpResponse(bytesBuffer *bytes.Buffer, w http.ResponseWriter, statusCode int, headers map[string]string) {
	// Set default headers.
	w.Header().Set("Content-Type", "text/html")

	// Set headers.
	for k, v := range headers {
		w.Header().Set(k, v)
	}

	// Set return status.
	w.WriteHeader(statusCode)

	_, err := bytesBuffer.WriteTo(w)
	if err != nil {
		ServerError(w, err)
	}
}
