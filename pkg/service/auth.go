package service

import (
	"crypto/sha256"
	"fmt"

	"github.com/yarikTri/darty"
	"github.com/yarikTri/darty/pkg/repository"
)

const salt = "9/v89v#idv*&"

// AuthService ..
type AuthService struct {
	repo repository.Authorization
}

// NewAuthService ..
func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

// CreateUser ..
func (s *AuthService) CreateUser(user darty.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
