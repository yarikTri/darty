package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/yarikTri/darty"
)

// AuthPostgres ..
type AuthPostgres struct {
	db *sqlx.DB
}

// NewAuthPostgres ..
func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db}
}

// CreateUser ..
func (r *AuthPostgres) CreateUser(user darty.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
