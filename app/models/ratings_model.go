package models

import (
	"time"

	"github.com/google/uuid"
)

type Ratings struct {
	ID            uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
	Company       string    `db:"company" json:"company" validate:"required,lte=255"`
	Author        string    `db:"author" json:"author" validate:"required,lte=255"`
	AuthorRole    string    `db:"author_role" json:"authorrole" validate:"required,lte=255"`
	CompanyRating int       `db:"company_rating" json:"rating" validate:"required,min=1,max=5"`
	Description   string    `db:"description" json:"description" validate:"required,lte=255"`
}
