package handler

import (
	"absensi-api/internal/domain"
	"absensi-api/middleware"
	"absensi-api/pkg/utils"
	"context"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type attendanceReportHandler struct {
	service domain.AttendanceService
}

func NewAttendanceReport(app *fiber.App, service domain.AttendanceService) {
	handler := &attendanceReportHandler{
		service: service,
	}

	routes := app.Group("/report", middleware.Auth())
	routes.Get("/", handler.Index)
	routes.Get("/:id", handler.History)
}

func (h *attendanceReportHandler) Index(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*15)
	defer cancel()

	var departementID *int
	params := ctx.Query("departement_id")
	if params != "" {
		id, err := strconv.Atoi(params)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(utils.ResponseError("05", "params departement_id must be integer"))
		}
		departementID = &id
	}

	var date *string
	params = ctx.Query("date")
	if params != "" {
		date = &params
	}

	result, err := h.service.FilterAttendance(c, departementID, date)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(utils.ResponseError("01", "Failed to get report"))
	}

	return ctx.Status(fiber.StatusOK).JSON(utils.ResponseSuccessWithData("00", "success", result))
}

func (h *attendanceReportHandler) History(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*15)
	defer cancel()

	result, err := h.service.GetAttendanceHistory(c, ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(utils.ResponseError("01", "failed to get attendance history"))
	}

	return ctx.Status(fiber.StatusOK).JSON(utils.ResponseSuccessWithData("00", "success", result))
}
