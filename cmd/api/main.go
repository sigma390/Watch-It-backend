package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const port = 8000

// type For App Config basiacally it Stores All
type application struct {
	Domain string
	DSN    string
	DB     *sql.DB
}

func main() {
	// app Config
	var app application

	//read From Command Line (DSN me Store kr rhe hai yeh Data )
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable connect_timeout=5", "Postgres Connection String")
	flag.Parse()

	//use of that App variable
	app.Domain = "localhost"
	//db conenction
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	app.DB = conn
	defer app.DB.Close() //Closing Connection Pool

	//handler
	//===========> OLD WAY <===================
	http.HandleFunc("/", Hello)

	//============> NEW WAY <==================
	app.routes()
	//starting the Server
	//===============? OLD WAY <=====================
	// err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	//================> NEW WAY <=======================
	fmt.Println("Starting Server on port", port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}

}
