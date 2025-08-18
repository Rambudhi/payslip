package service

import (
	"errors"
	"os"
	"time"

	"github.com/Rambudhi/payslip/internal/repository"
	"github.com/Rambudhi/payslip/internal/request"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(req request.LoginRequest) (string, error)
}

type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &authService{repo}
}

func (s *authService) Login(req request.LoginRequest) (string, error) {
	user, err := s.repo.GetUserByUsername(req.Username)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", errors.New("invalid username or password")
	}

	claims := jwt.MapClaims{
		"user_id":   user.ID,
		"user_name": user.Username,
		"role":      user.Role,
		"exp":       time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}
