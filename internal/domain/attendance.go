package domain

import (
	"absensi-api/dto"
	"context"
	"time"
)

type Attendance struct {
	ID           int        `gorm:"primaryKey;autoIncrement" json:"id"`
	AttendanceID string     `gorm:"type:varchar(100);uniqueIndex" json:"attendance_id"`
	EmployeeID   string     `gorm:"type:varchar(50);index" json:"employee_id"`
	ClockIn      time.Time  `gorm:"type:timestamp" json:"clock_in"`
	ClockOut     *time.Time `gorm:"type:timestamp" json:"clock_out"`
	CreatedAt    time.Time  `gorm:"createdAt" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"updatedAt" json:"updated_at"`

	// Relation
	Employee Employee `gorm:"foreignKey:EmployeeID;references:EmployeeID" json:"employee"`
}

type AttendanceRepository interface {
	FindAll(context.Context) ([]Attendance, error)
	FindByFilter(ctx context.Context, departement_id *int, date *string) ([]Attendance, error)
	FindAttendanceToday(ctx context.Context, employee_id string) (Attendance, error)
	Save(context.Context, Attendance) error
	Update(context.Context, Attendance) error
}

type AttendanceService interface {
	GetAllAttendance(context.Context) ([]AttendanceResponse, error)
	FilterAttendance(ctx context.Context, departement_id *int, date *string) ([]AttendanceResponse, error)
	GetAttendanceHistory(ctx context.Context, attendance_id string) ([]AttendanceHistoryResponse, error)
	ClockIn(ctx context.Context, data dto.AttendanceRequest) error
	ClockOut(ctx context.Context, data dto.AttendanceRequest) error
}

type AttendanceResponse struct {
	ID           int              `json:"id"`
	AttendanceID string           `json:"attendance_id"`
	Employee     EmployeeResponse `json:"employee"`
	ClockIn      time.Time        `json:"clock_in"`
	ClockOut     *time.Time       `json:"clock_out"`
	CreatedAt    time.Time        `json:"created_at"`
}
