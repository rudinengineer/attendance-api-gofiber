package repository

import (
	"absensi-api/internal/domain"
	"context"
	"time"

	"gorm.io/gorm"
)

type attendanceRepository struct {
	db *gorm.DB
}

func NewAttendance(db *gorm.DB) domain.AttendanceRepository {
	return &attendanceRepository{
		db: db,
	}
}

func (r *attendanceRepository) FindAll(ctx context.Context) ([]domain.Attendance, error) {
	return gorm.G[domain.Attendance](r.db).
		Preload("Employee", nil).
		Preload("AttendanceHistory", nil).
		Order("id DESC").
		Find(ctx)
}

func (r *attendanceRepository) FindByFilter(ctx context.Context, departement_id *int, date *string) ([]domain.Attendance, error) {
	query := r.db.Model(&domain.Attendance{}).
		Joins("JOIN employees ON employees.employee_id = attendances.employee_id").
		Preload("Employee").
		Preload("Employee.Departement").
		Order("id DESC")

	if departement_id != nil {
		query = query.Where("employees.departement_id = ?", *departement_id)
	}

	if date != nil {
		query = query.Where("DATE(attendances.created_at) = ?", *date)
	}

	var result []domain.Attendance
	if err := query.WithContext(ctx).Find(&result).Error; err != nil {
		return []domain.Attendance{}, err
	}

	return result, nil
}

func (r *attendanceRepository) FindAttendanceToday(ctx context.Context, employee_id string) (domain.Attendance, error) {
	today := time.Now().Format("2006-01-02")

	return gorm.G[domain.Attendance](r.db).
		Preload("Employee", nil).
		Where("employee_id = ?", employee_id).
		Where("DATE(created_at) = ?", today).
		Last(ctx)
}

func (r *attendanceRepository) Save(ctx context.Context, data domain.Attendance) error {
	return gorm.G[domain.Attendance](r.db).Create(ctx, &data)
}

func (r *attendanceRepository) Update(ctx context.Context, data domain.Attendance) error {
	_, err := gorm.G[domain.Attendance](r.db).Where("id = ?", data.ID).Updates(ctx, data)
	if err != nil {
		return err
	}
	return nil
}
