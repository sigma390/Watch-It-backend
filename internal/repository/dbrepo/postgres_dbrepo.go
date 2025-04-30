package dbrepo

import (
	"context"
	"database/sql"
	"time"

	"example.com/internal/models"
)

// PostgresDBRepo handles all database operations for PostgreSQL
type PostgresDBRepo struct {
	DB *sql.DB // Connection pool to the database
}

// Maximum time allowed for database operations
var DbTimeout = time.Second * 3

// Connection returns a connection pool to the database
func (m *PostgresDBRepo) Connection() *sql.DB {
	return m.DB
}

// AllMovies fetches all movies from the database, sorted by title
func (m *PostgresDBRepo) AllMovies() ([]*models.Movie, error) {
	// Create a timeout context to prevent long-running queries
	ctx, cancel := context.WithTimeout(context.Background(), DbTimeout)
	defer cancel() // Ensure resources are released when done

	// SQL query to get all movie information
	query := `
	SELECT 
	     id, title,release_date, runtime,  mpaa_rating, description,
		 coalesce(image, '') as image,
		 created_at, updated_at
	FROM 
	     movies
	ORDER BY 
	    title
	`

	// Execute the query
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Always close the result set

	// Slice to store our movie results
	var movies []*models.Movie

	// Loop through each row returned from database
	for rows.Next() {
		var movie models.Movie
		// Copy database values into our movie struct
		err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.ReleaseDate,
			&movie.RunTime,
			&movie.MPAARating,
			&movie.Description,
			&movie.Image,
			&movie.CreatedAt,
			&movie.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		// Add this movie to our results
		movies = append(movies, &movie)
	}

	// Return all movies found
	return movies, nil
}

func (m *PostgresDBRepo) GetUserByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DbTimeout)
	defer cancel()

	query := `select id, email, first_name, last_name, password, created_at, updated_at
	from users
	where email = $1
	`

	var user models.User
	row := m.DB.QueryRowContext(ctx, query, email)
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
