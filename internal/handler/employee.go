package handler

import (
	"absensi-api/dto"
	"absensi-api/internal/domain"
	"absensi-api/middleware"
	"absensi-api/pkg/utils"
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type employeeHandler struct {
	service domain.EmployeeService
}

func NewEmployee(app *fiber.App, service domain.EmployeeService) {
	handler := &employeeHandler{
		service: service,
	}

	routes := app.Group("/employee", middleware.Auth())

	routes.Get("/", handler.Index)
	routes.Get("/:id", handler.Detail)
	routes.Post("/", handler.Create)
	routes.Put("/:id", handler.Update)
	routes.Delete("/:id", handler.Delete)
}

func (h *employeeHandler) Index(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*15)
	defer cancel()

	result, err := h.service.GetAllEmployee(c)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(utils.ResponseError("01", "Failed to get data employee"))
	}

	return ctx.Status(fiber.StatusOK).JSON(utils.ResponseSuccessWithData("00", "success", result))
}

func (h *employeeHandler) Detail(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*15)
	defer cancel()

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.ResponseError("05", "params id must be integer"))
	}

	result, err := h.service.DetailEmployee(c, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(utils.ResponseError("06", "Data employee not found"))
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(utils.ResponseError("01", "Failed to get detail employee"))
	}

	return ctx.Status(fiber.StatusOK).JSON(utils.ResponseSuccessWithData("00", "success", result))
}

func (h *employeeHandler) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*15)
	defer cancel()

	// Parse Request
	var req dto.CreateEmployeeRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(utils.ResponseError("02", "cannot process request"))
	}

	// Validation Request
	errorMessages, err := utils.ValidationRequest(req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.ResponseErrorWithData("04", "validation error", errorMessages))
	}

	if err := h.service.CreateEmployee(c, req); err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1062: // duplicate key
				return ctx.Status(fiber.StatusBadRequest).JSON(utils.ResponseErrorWithData("04", "validation error", map[string]string{
					"employee_id": "Employee ID already taken",
				}))
			case 1452: // foreign key fail
				return ctx.Status(fiber.StatusBadRequest).JSON(utils.ResponseErrorWithData("04", "validation error", map[string]string{
					"departement_id": "Departement not found",
				}))
			}
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(utils.ResponseError("01", "failed to save data employee"))
	}

	return ctx.Status(fiber.StatusCreated).JSON(utils.ResponseSuccess("00", "success"))
}

func (h *employeeHandler) Update(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*15)
	defer cancel()

	// Parse Request
	var req dto.UpdateEmployeeRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(utils.ResponseError("02", "cannot process request"))
	}

	// Validation Request
	errorMessages, err := utils.ValidationRequest(req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.ResponseErrorWithData("04", "validation error", errorMessages))
	}

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.ResponseError("05", "params id must be integer"))
	}
	req.ID = id

	if err := h.service.UpdateEmployee(c, req); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(utils.ResponseError("01", "failed to save data employee"))
	}

	return ctx.Status(fiber.StatusCreated).JSON(utils.ResponseSuccess("00", "success"))
}

func (h *employeeHandler) Delete(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*15)
	defer cancel()

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.ResponseError("05", "params id must be integer"))
	}

	if err := h.service.DeleteEmployee(c, id); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(utils.ResponseError("01", "failed to delete data employee"))
	}

	return ctx.Status(fiber.StatusOK).JSON(utils.ResponseSuccess("00", "success"))
}
