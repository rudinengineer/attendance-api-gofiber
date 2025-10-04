package handler

import (
	"absensi-api/dto"
	"absensi-api/internal/domain"
	"absensi-api/pkg/utils"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

type attendanceHandler struct {
	service domain.AttendanceService
}

func NewAttendance(app *fiber.App, service domain.AttendanceService) {
	handler := &attendanceHandler{
		service: service,
	}

	routes := app.Group("/attendance")
	routes.Post("/", handler.ClockIn)
	routes.Put("/", handler.ClockOut)
}

func (h *attendanceHandler) ClockIn(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*15)
	defer cancel()

	// Parse Request
	var req dto.AttendanceRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(utils.ResponseError("02", "cannot process request"))
	}

	// Validation Request
	errorMessages, err := utils.ValidationRequest(req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.ResponseErrorWithData("04", "validation error", errorMessages))
	}

	if err := h.service.ClockIn(c, req); err != nil {
		if err.Error() == "employee not found" {
			return ctx.Status(fiber.StatusNotFound).JSON(utils.ResponseError("06", "Employee ID not found"))
		}

		if err.Error() == "you have already clocked in" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(utils.ResponseError("10", "You have already clocked in"))
		}

		if err.Error() == "max time error" {
			return ctx.Status(fiber.StatusForbidden).JSON(utils.ResponseError("03", "Check-in time limit has been exceeded"))
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(utils.ResponseError("01", "Failed to save data"))
	}

	return ctx.Status(fiber.StatusOK).JSON(utils.ResponseSuccess("00", "success"))
}

func (h *attendanceHandler) ClockOut(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*15)
	defer cancel()

	// Parse Request
	var req dto.AttendanceRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(utils.ResponseError("02", "cannot process request"))
	}

	// Validation Request
	errorMessages, err := utils.ValidationRequest(req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.ResponseErrorWithData("04", "validation error", errorMessages))
	}

	if err := h.service.ClockOut(c, req); err != nil {
		if err.Error() == "employee not found" {
			return ctx.Status(fiber.StatusNotFound).JSON(utils.ResponseError("06", "Employee ID not found"))
		}

		if err.Error() == "you have not clocked in yet" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(utils.ResponseError("10", "You have not clocked in yet"))
		}

		if err.Error() == "you have already clocked out" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(utils.ResponseError("10", "You have already clocked out"))
		}

		if err.Error() == "max time error" {
			return ctx.Status(fiber.StatusForbidden).JSON(utils.ResponseError("10", "Check-out time limit has been exceeded"))
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(utils.ResponseError("01", "Failed to save data"))
	}

	return ctx.Status(fiber.StatusOK).JSON(utils.ResponseSuccess("00", "success"))
}
