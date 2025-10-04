package repository

import (
	"absensi-api/internal/domain"
	"context"
	"time"

	"gorm.io/gorm"
)

type attendanceHistoryRepository struct {
	db *gorm.DB
}

func NewAttendanceHistory(db *gorm.DB) domain.AttendanceHistoryRepository {
	return &attendanceHistoryRepository{
		db: db,
	}
}

func (r *attendanceHistoryRepository) FindByAttendanceID(ctx context.Context, attendance_id string) ([]domain.AttendanceHistory, error) {
	return gorm.G[domain.AttendanceHistory](r.db).
		Preload("Attendance", nil).
		Preload("Employee", nil).
		Where("attendance_id = ?", attendance_id).
		Find(ctx)
}

func (r *attendanceHistoryRepository) FindByType(ctx context.Context, attendance_type int) (domain.AttendanceHistory, error) {
	return gorm.G[domain.AttendanceHistory](r.db).Where("attendance_type = ?", attendance_type).First(ctx)
}

func (r *attendanceHistoryRepository) FindLastByType(ctx context.Context, attendance_id string, employee_id string) (domain.AttendanceHistory, error) {
	today := time.Now().Format("2006-01-02")

	return gorm.G[domain.AttendanceHistory](r.db).
		Where("DATE(date_attendance) = ?", today).
		Where("attendance_id = ?", attendance_id).
		Where("employee_id = ?", employee_id).
		Last(ctx)
}

func (r *attendanceHistoryRepository) Save(ctx context.Context, data domain.AttendanceHistory) error {
	return gorm.G[domain.AttendanceHistory](r.db).Create(ctx, &data)
}

func (r *attendanceHistoryRepository) Update(ctx context.Context, data domain.AttendanceHistory) error {
	_, err := gorm.G[domain.AttendanceHistory](r.db).Where("id = ?", data.ID).Updates(ctx, data)
	if err != nil {
		return err
	}
	return nil
}
