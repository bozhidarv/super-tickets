package repository

import "github.com/jmoiron/sqlx"

// Repository aggregates all sub-repositories.
type Repository struct {
	UserRepo        UserRepository
	MovieRepo       MovieRepository
	ProjectionRepo  ProjectionRepository
	ReservationRepo ReservationRepository
}

// NewRepository creates a new Repository given a sqlx.DB connection.
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UserRepo:        NewUserRepository(db),
		MovieRepo:       NewMovieRepository(db),
		ProjectionRepo:  NewProjectionRepository(db),
		ReservationRepo: NewReservationRepository(db),
	}
}

// NewPostgresDB returns a new database connection.
func NewPostgresDB(dbURL string) (*sqlx.DB, error) {
	return sqlx.Connect("postgres", dbURL)
}
