package domain

import (
	"context"
	"time"
)

type AttendanceHistory struct {
	ID             int       `gorm:"primaryKey;autoIncrement" json:"id"`
	EmployeeID     string    `gorm:"type:varchar(50)" json:"employee_id"`
	AttendanceID   string    `gorm:"type:varchar(100);index" json:"attendance_id"`
	DateAttendance time.Time `gorm:"type:timestamp" json:"date_attendance"`
	AttendanceType int       `gorm:"type:tinyint" json:"attendance_type"`
	Description    string    `gorm:"type:text" json:"description"`
	CreatedAt      time.Time `gorm:"createdAt" json:"created_at"`
	UpdatedAt      time.Time `gorm:"updatedAt" json:"updated_at"`

	// Relation
	Employee   Employee   `gorm:"foreignKey:EmployeeID;references:EmployeeID" json:"employee"`
	Attendance Attendance `gorm:"foreignKey:AttendanceID;references:AttendanceID" json:"attendance"`
}

type AttendanceHistoryRepository interface {
	Save(context.Context, AttendanceHistory) error
	Update(context.Context, AttendanceHistory) error
	FindByAttendanceID(ctx context.Context, attendance_id string) ([]AttendanceHistory, error)
	FindByType(context.Context, int) (AttendanceHistory, error)
	FindLastByType(ctx context.Context, attendance_id string, employee_id string) (AttendanceHistory, error)
}

type AttendanceHistoryResponse struct {
	ID             int        `json:"id"`
	Employee       Employee   `json:"employee"`
	Attendance     Attendance `json:"attendance"`
	DateAttendance time.Time  `json:"date_attendance"`
	AttendanceType int        `json:"attendance_type"`
	Description    string     `json:"description"`
}
