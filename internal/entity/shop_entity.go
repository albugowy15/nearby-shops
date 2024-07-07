package entity

import (
	"database/sql"
	"time"
)

type Shop struct {
	ID          int64          `db:"id"`
	Name        string         `db:"name"`
	Description sql.NullString `db:"description"`
	City        string         `db:"city"`
	Location    string         `db:"location"`
	CreatedAt   time.Time      `db:"created_at"`
}

func (e *Shop) TableName() string {
	return "shops"
}
