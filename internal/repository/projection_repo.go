package repository

import (
	"moviereservationsystem/internal/models"

	"github.com/jmoiron/sqlx"
)

type ProjectionRepository interface {
	GetProjections() ([]models.Projection, error)
	CreateProjection(proj *models.Projection) error
	UpdateProjection(proj *models.Projection) error
	DeleteProjection(id int64) error
}

type projectionRepo struct {
	DB *sqlx.DB
}

func NewProjectionRepository(db *sqlx.DB) ProjectionRepository {
	return &projectionRepo{DB: db}
}

func (r *projectionRepo) GetProjections() ([]models.Projection, error) {
	var projections []models.Projection
	query := `SELECT id, movie_id, cinema, showtime FROM projections`
	err := r.DB.Select(&projections, query)
	return projections, err
}

func (r *projectionRepo) CreateProjection(proj *models.Projection) error {
	query := `INSERT INTO projections (movie_id, cinema, showtime)
	          VALUES ($1, $2, $3) RETURNING id`
	return r.DB.QueryRow(query, proj.MovieID, proj.Cinema, proj.Showtime).Scan(&proj.ID)
}

func (r *projectionRepo) UpdateProjection(proj *models.Projection) error {
	query := `UPDATE projections SET movie_id=$1, cinema=$2, showtime=$3 WHERE id=$4`
	_, err := r.DB.Exec(query, proj.MovieID, proj.Cinema, proj.Showtime, proj.ID)
	return err
}

func (r *projectionRepo) DeleteProjection(id int64) error {
	query := `DELETE FROM projections WHERE id=$1`
	_, err := r.DB.Exec(query, id)
	return err
}
