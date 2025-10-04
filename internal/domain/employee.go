package domain

import (
	"absensi-api/dto"
	"context"
	"time"
)

type Employee struct {
	ID            int       `gorm:"primaryKey;autoIncrement" json:"id"`
	EmployeeID    string    `gorm:"type:varchar(50);uniqueIndex" json:"employee_id"`
	DepartementID int       `gorm:"index" json:"departement_id"`
	Name          string    `gorm:"type:varchar(255)" json:"name"`
	Address       string    `gorm:"type:text" json:"address"`
	CreatedAt     time.Time `gorm:"createdAt" json:"created_at"`
	UpdatedAt     time.Time `gorm:"updatedAt" json:"updated_at"`

	// Relation
	Departement Departement `gorm:"foreignKey:DepartementID;references:ID" json:"departement"`
}

type EmployeeRepository interface {
	FindAll(context.Context) ([]Employee, error)
	Find(ctx context.Context, id int) (Employee, error)
	FindByEmployeeID(ctx context.Context, employee_id string) (Employee, error)
	Save(context.Context, Employee) error
	Update(context.Context, Employee) error
	Delete(ctx context.Context, id int) error
}

type EmployeeService interface {
	GetAllEmployee(context.Context) ([]EmployeeResponse, error)
	DetailEmployee(ccx context.Context, id int) (EmployeeResponse, error)
	CreateEmployee(context.Context, dto.CreateEmployeeRequest) error
	UpdateEmployee(context.Context, dto.UpdateEmployeeRequest) error
	DeleteEmployee(ctx context.Context, id int) error
}

type EmployeeResponse struct {
	ID          int         `json:"id"`
	EmployeeID  string      `json:"employee_id"`
	Departement Departement `json:"departement"`
	Name        string      `json:"name"`
	Address     string      `json:"address"`
}
