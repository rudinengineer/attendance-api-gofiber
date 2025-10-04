package service

import (
	"absensi-api/dto"
	"absensi-api/internal/domain"
	"absensi-api/pkg/auth"
	"absensi-api/pkg/utils"
	"context"

	"github.com/golang-jwt/jwt/v5"
)

type authService struct {
	repository domain.AuthRepository
}

func NewAuth(repository domain.AuthRepository) domain.AuthService {
	return &authService{
		repository: repository,
	}
}

func (s *authService) Login(ctx context.Context, req dto.AuthRequest) (string, error) {
	// Find username
	result, err := s.repository.FindUserByUsername(ctx, req.Username)
	if err != nil {
		return "", err
	}

	// Verify Password
	if err := utils.VerifyPassword(result.Password, req.Password); err != nil {
		return "", err
	}

	payload := jwt.MapClaims{
		"id":       result.ID,
		"name":     result.Name,
		"username": result.Username,
	}

	// Generate JWT Token
	token, err := auth.GenerateToken(payload)
	if err != nil {
		return "", err
	}

	return token, nil
}
