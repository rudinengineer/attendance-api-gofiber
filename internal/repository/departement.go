package repository

import (
	"absensi-api/internal/domain"
	"context"
	"errors"

	"gorm.io/gorm"
)

type departementRepository struct {
	db *gorm.DB
}

func NewDepartement(db *gorm.DB) domain.DepartementRepository {
	return &departementRepository{
		db: db,
	}
}

func (r *departementRepository) FindAll(ctx context.Context) ([]domain.Departement, error) {
	return gorm.G[domain.Departement](r.db).Order("id DESC").Find(ctx)
}

func (r *departementRepository) Find(ctx context.Context, id int) (domain.Departement, error) {
	return gorm.G[domain.Departement](r.db).Where("id = ?", id).First(ctx)
}

func (r *departementRepository) Save(ctx context.Context, data domain.Departement) error {
	return gorm.G[domain.Departement](r.db).Create(ctx, &data)
}

func (r *departementRepository) Update(ctx context.Context, data domain.Departement) error {
	_, err := gorm.G[domain.Departement](r.db).Where("id = ?", data.ID).Updates(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func (r *departementRepository) Delete(ctx context.Context, id int) error {
	result, err := gorm.G[domain.Departement](r.db).Where("id = ?", id).Delete(ctx)
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("failed to delete data")
	}
	return nil
}
