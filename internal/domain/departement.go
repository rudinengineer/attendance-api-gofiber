package domain

import (
	"absensi-api/dto"
	"context"
	"time"
)

type Departement struct {
	ID              int       `gorm:"primaryKey;autoIncrement" json:"id"`
	DepartementName string    `gorm:"type:varchar(255)" json:"departement_name"`
	MaxClockInTime  time.Time `gorm:"type:time" json:"max_clock_in_time"`
	MaxClockInOut   time.Time `gorm:"type:time" json:"max_clock_in_out"`
}

type DepartementRepository interface {
	FindAll(context.Context) ([]Departement, error)
	Find(ctx context.Context, id int) (Departement, error)
	Save(context.Context, Departement) error
	Update(context.Context, Departement) error
	Delete(ctx context.Context, id int) error
}

type DepartementService interface {
	GetAllDepartement(context.Context) ([]Departement, error)
	DetailDepartement(ctx context.Context, id int) (Departement, error)
	CreateDepartement(context.Context, dto.CreateDepartementRequest) error
	UpdateDepartement(context.Context, dto.UpdateDepartementRequest) error
	DeleteDepartement(ctx context.Context, id int) error
}
