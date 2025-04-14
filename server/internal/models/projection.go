package models

import "time"

type Projection struct {
	ID       int64     `db:"id"       json:"id"`
	MovieID  int64     `db:"movie_id" json:"movie_id"`
	Cinema   string    `db:"cinema"   json:"cinema"`
	Showtime time.Time `db:"showtime" json:"showtime"`
}
