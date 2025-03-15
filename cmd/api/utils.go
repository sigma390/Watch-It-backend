package main

import (
	"encoding/json"
	"errors"
	"io"
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

//==================> read Json <================

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	//max Bytes Json must not over 1mb
	maxBytes := 1024 * 1024

	//READ SIZE OF BODY
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	//decode File
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(data)
	if err != nil {
		return err
	}

	//validate Single File Only
	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only contain a single JSON value")
	}

	return nil
}

//====================> error Json <====================

// errorJSON is a helper method that generates a standardized JSON error response
// It takes an error, converts it to a JSONResponse, and sends it to the client
// Optionally accepts a custom HTTP status code (defaults to 400 Bad Request if not provided)
func (app *application) errorJSON(w http.ResponseWriter, err error, status ...int) error {
	// Default to HTTP 400 Bad Request if no status code is provided
	statusCode := http.StatusBadRequest

	// If a custom status code is provided as a variadic parameter, use it instead
	if len(status) > 0 {
		statusCode = status[0]
	}

	// Create the error response payload
	var payLoad JSONResponse
	payLoad.Error = true          // Mark this as an error response
	payLoad.Message = err.Error() // Use the error's message as the response message

	// Use the writeJSON helper to send the error response to the client
	// This handles JSON marshaling, setting headers, and writing to the response
	return app.writeJSON(w, statusCode, payLoad)
}
