package models

// User represents a registered user.
type User struct {
	ID       int64  `db:"id"       json:"id"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"` // hashed password
	Role     string `db:"role"     json:"role"`
}
