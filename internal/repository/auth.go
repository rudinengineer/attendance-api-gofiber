package repository

import (
	"absensi-api/internal/domain"
	"context"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuth(db *gorm.DB) domain.AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) FindUserByUsername(ctx context.Context, username string) (domain.User, error) {
	return gorm.G[domain.User](r.db).Where("username = ?", username).First(ctx)
}
