package service

import (
	"github.com/yarikTri/darty"
	"github.com/yarikTri/darty/pkg/repository"
)

// Authorization ..
type Authorization interface {
	CreateUser(user darty.User) (int, error)
}

// Event ..
type Event interface {
}

// Service ..
type Service struct {
	Authorization
	Event
}

// NewService inits empty Service
func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
