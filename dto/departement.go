package dto

type CreateDepartementRequest struct {
	DepartementName string `json:"departement_name" validate:"required"`
	MaxClockInTime  string `json:"max_clock_in_time" validate:"required"`
	MaxClockInOut   string `json:"max_clock_in_out" validate:"required"`
}

type UpdateDepartementRequest struct {
	ID              int    `json:"id"`
	DepartementName string `json:"departement_name" validate:"required"`
	MaxClockInTime  string `json:"max_clock_in_time" validate:"required"`
	MaxClockInOut   string `json:"max_clock_in_out" validate:"required"`
}
