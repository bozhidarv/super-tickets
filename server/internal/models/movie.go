package models

type Movie struct {
	ID          int64  `db:"id"          json:"id"`
	Title       string `db:"title"       json:"title"`
	Description string `db:"description" json:"description"`
	Duration    int    `db:"duration"    json:"duration"`
}
