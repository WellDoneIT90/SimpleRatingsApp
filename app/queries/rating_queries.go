package queries

import (
	"github.com/WellDoneIT90/SimpleRatingsApp/app/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// RatingQueries struct for queries from Ratings model
type RatingQueries struct {
	*sqlx.DB
}

// GetRatings method for getting all ratings.
func (q *RatingQueries) GetRatings() ([]models.Ratings, error) {
	// Define ratings variable
	ratings := []models.Ratings{}

	//define query string
	query := `SELECT * FROM ratings`

	// send query to database
	err := q.Get(&ratings, query)
	if err != nil {
		return ratings, err
	}

	return ratings, nil
}

// GetRating method for getting one rating by given ID.
func (q *RatingQueries) GetRating(id uuid.UUID) (models.Ratings, error) {
	// Define rating variable
	rating := models.Ratings{}

	// Define query string
	query := `SELECT * FROM ratings WHERE id = $1`

	// Send query to database
	err := q.Get(&rating, query, id)
	if err != nil {
		return rating, err
	}

	return rating, nil
}

// CreateRating method to create new rating by given rating object.
func (q *RatingQueries) CreateRating(r *models.Ratings) error {
	// Define query string
	query := `INSERT INTO ratings VALUES ($1, $2, $3, $4, $5, $6, $7)`

	// Send query to database
	_, err := q.Exec(query, r.ID, r.CreatedAt, r.Company, r.Author, r.AuthorRole, r.CompanyRating, r.Description)
	if err != nil {
		return err
	}

	return nil
}

// DeleteRating method to delete rating by given ID.
func (q *RatingQueries) DeleteRating(id uuid.UUID) error {
	// Define query string
	query := `DELETE FROM ratings WHERE id = $1`

	// Send query to database
	_, err := q.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
