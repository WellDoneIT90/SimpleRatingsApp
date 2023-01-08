package models

import (
	"time"

	"github.com/google/uuid"
)

type Rating struct {
	ID            uuid.UUID `db:"ID" json:"id" validate:"required,uuid"`
	CreatedAt     time.Time `db:"CREATED_AT" json:"createdat"`
	Company       string    `db:"COMPANY" json:"company" validate:"required,lte=255"`
	Author        string    `db:"AUTHOR" json:"author" validate:"required,lte=255"`
	AuthorRole    string    `db:"AUTHOR_ROLE" json:"authorrole" validate:"required,lte=255"`
	CompanyRating int       `db:"COMPANY_RATING" json:"rating" validate:"required,min=1,max=5"`
	Description   string    `db:"DESCRIPTION" json:"description" validate:"required,lte=255"`
}
