package Response

import (
	"encoding/json"
	"net/http"
)

// w io.Writer
func JsonResponse(value interface{}, w http.ResponseWriter, statusCode int, headers map[string]string) {
	// Set default headers.
	w.Header().Set("Content-Type", "application/json")

	// Set headers.
	for k, v := range headers {
		w.Header().Set(k, v)
	}

	// Set return status.
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(value)
	if err != nil {
		ServerError(w, err)
	}
}
