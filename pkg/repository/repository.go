package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/yarikTri/darty"
)

// Authorization ..
type Authorization interface {
	CreateUser(user darty.User) (int, error)
	GetUser(username, password string) (darty.User, error)
}

// Event ..
type Event interface {
}

// Repository ..
type Repository struct {
	Authorization
	Event
}

// NewRepository inits empty Repository
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{Authorization: NewAuthPostgres(db)}
}
