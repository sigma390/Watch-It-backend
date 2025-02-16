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
	//===========> OLD WAY <===================
	http.HandleFunc("/", Hello)

	//============> NEW WAY <==================
	app.routes()
	//starting the Server
	//===============? OLD WAY <=====================
	// err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	//================> NEW WAY <=======================
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}

}
