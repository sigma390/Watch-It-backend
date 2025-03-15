package main

import (
	"encoding/json"
	"net/http"
)

// JSONResponse defines the structure for all JSON responses from our API
// It provides a consistent format with error status, message, and optional data
type JSONResponse struct {
	Error   bool        `json:"error"`          // Indicates if the response contains an error
	Message string      `json:"message"`        // A human-readable message about the response
	Data    interface{} `json:"data,omitempty"` // Optional data payload, omitted when empty
}

// writeJSON marshals data to JSON and writes it to the http.ResponseWriter
// It handles setting appropriate headers and status codes
func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	// Convert the data to JSON
	out, err := json.Marshal(data)
	if err != nil {
		return err // Return error if JSON marshaling fails
	}

	// Add any custom headers if provided
	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	// Set content type to application/json
	w.Header().Set("Content-Type", "application/json")

	// Set the HTTP status code
	w.WriteHeader(status)

	// Write the JSON data to the response
	_, err = w.Write(out)
	if err != nil {
		return err // Return error if writing to response fails
	}

	return nil // Return nil on success
}
