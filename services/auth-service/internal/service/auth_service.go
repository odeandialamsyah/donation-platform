package service

import (
	"errors"
	"time"

	"donation-platform/auth-service/internal/auth"
	"donation-platform/auth-service/internal/entity"
	"donation-platform/auth-service/internal/repository"

	"github.com/google/uuid"
)

type AuthService struct {
	UserRepo *repository.UserRepository
	JWT      *auth.JWTService
}

func NewAuthService(userRepo *repository.UserRepository, jwt *auth.JWTService) *AuthService {
	return &AuthService{UserRepo: userRepo, JWT: jwt}
}

func (s *AuthService) Register(name, email, password string) (string, string, error) {

	existing, _ := s.UserRepo.FindByEmail(email)
	if existing.ID != "" {
		return "", "", errors.New("email already exists")
	}

	hashed, err := auth.HashPassword(password)
	if err != nil {
		return "", "", err
	}

	user := &entity.User{
		ID:        uuid.NewString(),
		Name:      name,
		Email:     email,
		Password:  hashed,
		Role:      "donor",
		CreatedAt: time.Now(),
	}

	err = s.UserRepo.Create(user)
	if err != nil {
		return "", "", err
	}

	accessToken, err := s.JWT.GenerateAccessToken(user.ID, user.Role)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := s.JWT.GenerateRefreshToken(user.ID)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
