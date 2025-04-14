package repository

import (
	"supertickets/internal/models"
	"time"

	"github.com/jmoiron/sqlx"
)

type ReservationRepository interface {
	CreateReservation(resv *models.Reservation) error
	GetReservationsByUser(userID int64) ([]models.Reservation, error)
	DeleteReservation(id int64) error
}

type reservationRepo struct {
	DB *sqlx.DB
}

func NewReservationRepository(db *sqlx.DB) ReservationRepository {
	return &reservationRepo{DB: db}
}

func (r *reservationRepo) CreateReservation(resv *models.Reservation) error {
	query := `INSERT INTO reservations (user_id, projection_id, seats, created_at)
	          VALUES ($1, $2, $3, $4) RETURNING id`
	return r.DB.QueryRow(query, resv.UserID, resv.ProjectionID, resv.Seats, time.Now()).
		Scan(&resv.ID)
}

func (r *reservationRepo) GetReservationsByUser(userID int64) ([]models.Reservation, error) {
	var reservations []models.Reservation
	query := `SELECT id, user_id, projection_id, seats FROM reservations WHERE user_id=$1`
	err := r.DB.Select(&reservations, query, userID)
	return reservations, err
}

func (r *reservationRepo) DeleteReservation(id int64) error {
	query := `DELETE FROM reservations WHERE id=$1`
	_, err := r.DB.Exec(query, id)
	return err
}
