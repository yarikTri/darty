package service

import "github.com/yarikTri/darty/pkg/repository"

// Authorization ..
type Authorization interface {
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
	return &Service{}
}
