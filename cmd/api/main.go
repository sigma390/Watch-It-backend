package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"example.com/internal/repository/dbrepo"
)

const port = 8000

// application struct holds the application-wide dependencies and configuration
// This provides a clean way to share these resources across different parts of the app
type application struct {
	Domain string                // Server domain name
	DSN    string                // Database connection string
	DB     dbrepo.PostgresDBRepo // Repository for database operations
}

func main() {
	// Initialize the application struct
	var app application

	// Parse command line flags
	// DSN (Data Source Name) contains all PostgreSQL connection details
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable connect_timeout=5", "Postgres Connection String")
	flag.Parse()

	// Set application domain
	app.Domain = "localhost"

	// Connect to the database
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err) // Exit if database connection fails
	}

	// Initialize the PostgreSQL repository with our database connection
	app.DB = dbrepo.PostgresDBRepo{DB: conn}
	defer app.DB.Connection().Close() // Ensure connection pool is closed when app exits

	// Set up HTTP routes
	// Old way (simple handler function)
	http.HandleFunc("/", Hello)

	// New way (using application's route method that returns a router)
	app.routes()

	// Start the HTTP server
	fmt.Println("Starting Server on port", port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err) // Exit if server fails to start
	}
}
