package routes

import (
	"absensi-api/internal/handler"
	"absensi-api/internal/repository"
	"absensi-api/internal/service"
	"absensi-api/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func New(app *fiber.App, db *gorm.DB) {
	// Auth
	authRepository := repository.NewAuth(db)
	authService := service.NewAuth(authRepository)
	handler.NewAuth(app, authService)

	// Departement
	departementRepository := repository.NewDepartement(db)
	departementService := service.NewDepartement(departementRepository)
	handler.NewDepartement(app, departementService)

	// Employee
	employeeRepository := repository.NewEmployee(db)
	employeeService := service.NewEmployee(employeeRepository)
	handler.NewEmployee(app, employeeService)

	// Attendance
	attendanceRepository := repository.NewAttendance(db)
	attendanceHistoryRepository := repository.NewAttendanceHistory(db)
	attendanceService := service.NewAttendance(
		attendanceRepository,
		attendanceHistoryRepository,
		employeeRepository,
	)
	handler.NewAttendance(app, attendanceService)

	// Report
	handler.NewAttendanceReport(app, attendanceService)

	// 404 Not Found
	app.Use(func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusNotFound).JSON(utils.ResponseError("12", "Route not found"))
	})
}
