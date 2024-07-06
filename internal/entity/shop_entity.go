package entity

type Shop struct {
	ID          int64  `db:"id"`
	Name        int64  `db:"name"`
	Description int64  `db:"description"`
	Location    string `db:"location"`
	CreatedAt   int64  `db:"created_at"`
}

func (e *Shop) TableName() string {
	return "shops"
}
