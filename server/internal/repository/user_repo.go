package repository

import (
	"supertickets/internal/models"
	"time"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByUsername(username string) (*models.User, error)
	GetAllUsers() ([]models.User, error)
	GetUserById(id int64) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id int64) error
}

type userRepo struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepo{DB: db}
}

func (r *userRepo) CreateUser(user *models.User) error {
	query := `INSERT INTO users (username, password, role, created_at)
	          VALUES ($1, $2, $3, $4) RETURNING id`
	return r.DB.QueryRow(query, user.Username, user.Password, user.Role, time.Now()).Scan(&user.ID)
}

func (r *userRepo) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	query := `SELECT id, username, password, role FROM users WHERE username = $1`
	err := r.DB.Get(&user, query, username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) GetAllUsers() ([]models.User, error) {
	var users []models.User
	query := `SELECT id, username, role FROM users`
	err := r.DB.Select(&users, query)
	return users, err
}

func (r *userRepo) GetUserById(id int64) (*models.User, error) {
	var user models.User
	query := `SELECT id, username, password, role FROM users WHERE id = $1`
	err := r.DB.Get(&user, query, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) UpdateUser(user *models.User) error {
	query := `UPDATE users SET username = $1, password = $2, role = $3 WHERE id = $4`
	_, err := r.DB.Exec(query, user.Username, user.Password, user.Role, user.ID)
	return err
}

func (r *userRepo) DeleteUser(id int64) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}
