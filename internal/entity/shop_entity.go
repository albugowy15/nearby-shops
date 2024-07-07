package entity

import (
	"database/sql"
	"time"
)

type Shop struct {
	CreatedAt   time.Time      `db:"created_at"`
	Name        string         `db:"name"`
	City        string         `db:"city"`
	Location    string         `db:"location"`
	Description sql.NullString `db:"description"`
	ID          int64          `db:"id"`
}

func (e *Shop) TableName() string {
	return "shops"
}
