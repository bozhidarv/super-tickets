package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	UserRepo        UserRepository
	MovieRepo       MovieRepository
	ProjectionRepo  ProjectionRepository
	ReservationRepo ReservationRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UserRepo:        NewUserRepository(db),
		MovieRepo:       NewMovieRepository(db),
		ProjectionRepo:  NewProjectionRepository(db),
		ReservationRepo: NewReservationRepository(db),
	}
}

func NewPostgresDB(dbURL string) (*sqlx.DB, error) {
	return sqlx.Connect("postgres", dbURL)
}

func ClosePostgresDB(db *sqlx.DB) {
	err := db.Close()
	if err != nil {
		return
	}
}
