package dto

type AttendanceRequest struct {
	EmployeeID  string `json:"employee_id" validate:"required"`
	Description string `json:"description"`
}
