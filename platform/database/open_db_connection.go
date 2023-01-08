package database

import "github.com/WellDoneIT90/SimpleRatingsApp/app/queries"

type Queries struct {
	*queries.RatingQueries
}

// OpenDBConnection func
func OpenDBConnection() (*Queries, error) {
	// Define new PostgreSQL connection
	db, err := PostgreSQLConnection()
	if err != nil {
		return nil, err
	}

	return &Queries{
		RatingQueries: &queries.RatingQueries{DB: db},
	}, nil
}
