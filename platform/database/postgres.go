package database

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
)

func PostgreSQLConnection() (*sqlx.DB, error) {
	// Define Database connection settings
	maxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	maxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	maxLifetimeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))

	// Define database connection for PostgreSQL
	db, err := sqlx.Connect("pgx", os.Getenv("DB_SERVER_URL"))
	if err != nil {
		return nil, fmt.Errorf("error, not connected to database, %w", err)
	}

	// Set database connection settings
	db.SetMaxOpenConns(maxConn)
	db.SetMaxIdleConns(maxIdleConn)
	db.SetConnMaxLifetime(time.Duration(maxLifetimeConn))

	// Try to ping database
	if err := db.Ping(); err != nil {
		defer db.Close()
		return nil, fmt.Errorf("error, sending ping to database was not possible %w", err)
	}

	return db, nil
}
