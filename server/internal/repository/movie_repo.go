package repository

import (
	"supertickets/internal/models"

	"github.com/jmoiron/sqlx"
)

type MovieRepository interface {
	GetMovies() ([]models.Movie, error)
	CreateMovie(movie *models.Movie) error
	UpdateMovie(movie *models.Movie) error
	DeleteMovie(id int64) error
}

type movieRepo struct {
	DB *sqlx.DB
}

func NewMovieRepository(db *sqlx.DB) MovieRepository {
	return &movieRepo{DB: db}
}

func (r *movieRepo) GetMovies() ([]models.Movie, error) {
	var movies []models.Movie
	query := `SELECT id, title, description, duration FROM movies`
	err := r.DB.Select(&movies, query)
	return movies, err
}

func (r *movieRepo) CreateMovie(movie *models.Movie) error {
	query := `INSERT INTO movies (title, description, duration)
	          VALUES ($1, $2, $3) RETURNING id`
	return r.DB.QueryRow(query, movie.Title, movie.Description, movie.Duration).Scan(&movie.ID)
}

func (r *movieRepo) UpdateMovie(movie *models.Movie) error {
	query := `UPDATE movies SET title=$1, description=$2, duration=$3 WHERE id=$4`
	_, err := r.DB.Exec(query, movie.Title, movie.Description, movie.Duration, movie.ID)
	return err
}

func (r *movieRepo) DeleteMovie(id int64) error {
	query := `DELETE FROM movies WHERE id=$1`
	_, err := r.DB.Exec(query, id)
	return err
}
