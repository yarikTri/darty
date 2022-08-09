package service

import (
	"crypto/sha256"
	"fmt"
	"time"

	// jwt
	"github.com/dgrijalva/jwt-go"
	"github.com/yarikTri/darty"
	"github.com/yarikTri/darty/pkg/repository"
)

const (
	salt       = "9/v89v#idvcsccd*&"
	signingKey = "mk!;asclv689*&ewkhfnc#$2"
	tokenTTL   = 12 * time.Hour
)

// tokenClaims ..
type tokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

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

// GenerateToken ..
func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
		},
		user.ID,
	})

	return token.SignedString([]byte(signingKey))
}

func generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
