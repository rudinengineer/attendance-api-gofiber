package domain

import (
	"absensi-api/dto"
	"context"
)

type AuthRepository interface {
	FindUserByUsername(ctx context.Context, username string) (User, error)
}

type AuthService interface {
	Login(context.Context, dto.AuthRequest) (string, error)
}
