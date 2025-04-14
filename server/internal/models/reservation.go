package models

type Reservation struct {
	ID           int64 `db:"id"            json:"id"`
	UserID       int64 `db:"user_id"       json:"user_id"`
	ProjectionID int64 `db:"projection_id" json:"projection_id"`
	Seats        int   `db:"seats"         json:"seats"`
}
