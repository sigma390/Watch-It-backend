package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"example.com/internal/models"
)

//handler Function

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hellow")
}

// new Route Handler
func (app *application) HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page , hello from %s", app.Domain)

}

func (app *application) Statuss(w http.ResponseWriter, r *http.Request) {
	//kya send krna hai Frontend ko
	var payload = struct {
		Status  string `json:"status"`
		Version string `json:"version"`
		Message string `json:"msg"`
	}{
		Status:  "success",
		Version: "1.0.0",
		Message: "API is running fine",
	}
	//convert to json using Marshal

	out, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}
	//set headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//send json
	w.Write(out)
	// end of function
	// return nil
}

//movie Interface

func (app *application) Movies(w http.ResponseWriter, r *http.Request) {

	//get all movies
	var movies []models.Movie = []models.Movie{
		{
			ID:          1,
			Title:       "The Dark Knight",
			Description: "A movie about a batman",
			ReleaseDate: time.Now(),
			Duration:    152,
			Genre:       "Action",
			Rating:      8.9,
		},
		{
			ID:          2,
			Title:       "The Dark Knight Rises",
			Description: "A movie about a batman",
			ReleaseDate: time.Now(),
			Duration:    152,
			Genre:       "Action",
			Rating:      8.9,
		},
		{
			ID:          3,
			Title:       "The Dark Knight Rises",
			Description: "A movie about a batman",
			ReleaseDate: time.Now(),
			Duration:    152,
			Genre:       "Action",
			Rating:      8.9,
		},
	}

	//convert to json
	out, err := json.Marshal(movies)
	if err != nil {
		log.Fatal(err)
	}

	//set headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	//send json
	w.Write(out)

}
