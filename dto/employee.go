package dto

type CreateEmployeeRequest struct {
	EmployeeID    string `json:"employee_id" validate:"required,numeric"`
	DepartementID int    `json:"departement_id" validate:"required,numeric"`
	Name          string `json:"name" validate:"required"`
	Address       string `json:"address" validate:"required"`
}

type UpdateEmployeeRequest struct {
	ID            int    `json:"id"`
	EmployeeID    string `json:"employee_id" validate:"required,numeric"`
	DepartementID int    `json:"departement_id" validate:"required,numeric"`
	Name          string `json:"name" validate:"required"`
	Address       string `json:"address" validate:"required"`
}
