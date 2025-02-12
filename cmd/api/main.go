package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8000

// type For App Config basiacally it Stores All
type application struct {
	Domain string
}

func main() {
	// app Config
	var app application
	//use of that App variable
	app.Domain = "localhost"
	//db conenction

	//handler
	http.HandleFunc("/", Hello)
	//starting the Server
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal(err)
	}

}
