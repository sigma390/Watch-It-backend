package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"example.com/internal/repository/dbrepo"
)

const port = 8000

// application struct holds the application-wide dependencies and configuration
// This provides a clean way to share these resources across different parts of the app
type application struct {
	Domain       string                // Server domain name
	DSN          string                // Database connection string
	DB           dbrepo.PostgresDBRepo // Repository for database operations
	auth         Auth                  // Authentication configuration
	JWTSecret    string                // JWT Secret
	JWTIssuer    string                // JWT Issuer
	JWTAudience  string                // JWT Audience
	CookieDomain string                // Cookie Domain
}

func main() {
	// Initialize the application struct
	var app application

	// Parse command line flags
	// DSN (Data Source Name) contains all PostgreSQL connection details
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=movies sslmode=disable connect_timeout=5", "Postgres Connection String")
	flag.StringVar(&app.JWTSecret, "jwt-secret", "secret", "JWT Secret")
	flag.StringVar(&app.JWTIssuer, "jwt-issuer", "example.com", "JWT Issuer")
	flag.StringVar(&app.JWTAudience, "jwt-audience", "example.com", "JWT Audience")
	flag.StringVar(&app.CookieDomain, "cookie-domain", "localhost", "Cookie Domain")
	flag.Parse()

	// Set application domain
	app.Domain = "localhost"

	// Connect to the database
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err) // Exit if database connection fails
	}

	//Auth Object Initialisation
	app.auth = Auth{
		Issuer:        app.JWTIssuer,          // The entity issuing the JWT tokens (typically your API name)
		Audience:      app.JWTAudience,        // The intended recipient of the token (typically your frontend app)
		Secret:        app.JWTSecret,          // Secret key used to sign JWT tokens
		TokenExpiry:   24 * time.Minute,       // How long access tokens remain valid
		RefreshExpiry: 24 * time.Minute,       // How long refresh tokens remain valid
		CookieDomain:  app.CookieDomain,       // Domain for which cookies are valid
		CookiePath:    "/",                    // Path restriction for cookies
		CookieName:    "__Host-refresh_token", //
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
