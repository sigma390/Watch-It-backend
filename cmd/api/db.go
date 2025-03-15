package main

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// openDB initializes a database connection pool using the provided DSN
// DSN (Data Source Name) is a connection string containing database credentials and connection details
// Example DSN: "postgres://username:password@localhost:5432/database_name"
func openDB(dsn string) (*sql.DB, error) {
	// Create a new connection pool using the pgx PostgreSQL driver
	// Note: This doesn't establish connections yet, just sets up the pool
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		// Return error if pool creation fails
		return nil, err
	}

	// Verify connectivity by establishing an actual connection
	// This confirms the DSN is valid and database is accessible
	err = db.Ping()
	if err != nil {
		// Return error if connection test fails
		return nil, err
	}

	// Return the verified connection pool ready for use
	return db, nil
}

// connectToDB establishes a connection to the database using the application's DSN
// and returns the database connection pool
func (app *application) connectToDB() (*sql.DB, error) {
	connection, err := openDB(app.DSN)

	if err != nil {
		return nil, err
	}
	log.Println("Connected to Postgres DB")

	return connection, nil
}
