package repository

// Authorization ..
type Authorization interface {
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
func NewRepository() *Repository {
	return &Repository{}
}
