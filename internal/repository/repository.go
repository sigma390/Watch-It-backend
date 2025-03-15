package repository

import (
	"database/sql"

	"example.com/internal/models"
)

// DBRepository defines the interface for database operations
// By using an interface, we can:
// - Swap between different database types (PostgreSQL, MySQL, etc.)
// - Create mock implementations for testing
// - Change database technology without changing application code
type DBRepository interface {
	// AllMovies retrieves all movies from the database
	// Returns a slice of movie pointers and any error encountered
	Connection() *sql.DB
	AllMovies() ([]*models.Movie, error)
}
