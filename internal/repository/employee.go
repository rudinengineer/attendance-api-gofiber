package repository

import (
	"absensi-api/internal/domain"
	"context"
	"errors"

	"gorm.io/gorm"
)

type employeeRepository struct {
	db *gorm.DB
}

func NewEmployee(db *gorm.DB) domain.EmployeeRepository {
	return &employeeRepository{
		db: db,
	}
}

func (r *employeeRepository) FindAll(ctx context.Context) ([]domain.Employee, error) {
	var data []domain.Employee

	// return gorm.G[domain.Employee](r.db).Preload("Departement", nil).Find(ctx)
	if err := r.db.WithContext(ctx).Preload("Departement").Order("id DESC").Find(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (r *employeeRepository) Find(ctx context.Context, id int) (domain.Employee, error) {
	return gorm.G[domain.Employee](r.db).Preload("Departement", nil).Where("id = ?", id).First(ctx)
}

func (r *employeeRepository) FindByEmployeeID(ctx context.Context, employee_id string) (domain.Employee, error) {
	return gorm.G[domain.Employee](r.db).Preload("Departement", nil).Where("employee_id = ?", employee_id).First(ctx)
}

func (r *employeeRepository) Save(ctx context.Context, data domain.Employee) error {
	return gorm.G[domain.Employee](r.db).Create(ctx, &data)
}

func (r *employeeRepository) Update(ctx context.Context, data domain.Employee) error {
	_, err := gorm.G[domain.Employee](r.db).Where("id = ?", data.ID).Updates(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func (r *employeeRepository) Delete(ctx context.Context, id int) error {
	result, err := gorm.G[domain.Employee](r.db).Where("id = ?", id).Delete(ctx)
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("failed to delete data")
	}
	return nil
}
