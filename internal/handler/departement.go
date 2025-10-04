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

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type departementHandler struct {
	service domain.DepartementService
}

func NewDepartement(app *fiber.App, service domain.DepartementService) {
	handler := &departementHandler{
		service: service,
	}

	routes := app.Group("/departement", middleware.Auth())

	routes.Get("/", handler.Index)
	routes.Get("/:id", handler.Detail)
	routes.Post("/", handler.Create)
	routes.Put("/:id", handler.Update)
	routes.Delete("/:id", handler.Delete)
}

func (h *departementHandler) Index(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*15)
	defer cancel()

	result, err := h.service.GetAllDepartement(c)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(utils.ResponseError("01", "Failed to get data departement"))
	}

	return ctx.Status(fiber.StatusOK).JSON(utils.ResponseSuccessWithData("00", "success", result))
}

func (h *departementHandler) Detail(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*15)
	defer cancel()

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.ResponseError("05", "params id must be integer"))
	}

	result, err := h.service.DetailDepartement(c, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(utils.ResponseError("06", "Data departement not found"))
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(utils.ResponseError("01", "Failed to get detail departement"))
	}

	return ctx.Status(fiber.StatusOK).JSON(utils.ResponseSuccessWithData("00", "success", result))
}

func (h *departementHandler) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*15)
	defer cancel()

	// Parse Request
	var req dto.CreateDepartementRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(utils.ResponseError("02", "cannot process request"))
	}

	// Validation Request
	errorMessages, err := utils.ValidationRequest(req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.ResponseErrorWithData("04", "validation error", errorMessages))
	}

	if err := h.service.CreateDepartement(c, req); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(utils.ResponseError("01", "failed to save data departement"))
	}

	return ctx.Status(fiber.StatusCreated).JSON(utils.ResponseSuccess("00", "success"))
}

func (h *departementHandler) Update(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*15)
	defer cancel()

	// Parse Request
	var req dto.UpdateDepartementRequest
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

	if err := h.service.UpdateDepartement(c, req); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(utils.ResponseError("01", "failed to save data departement"))
	}

	return ctx.Status(fiber.StatusOK).JSON(utils.ResponseSuccess("00", "success"))
}

func (h *departementHandler) Delete(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*15)
	defer cancel()

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.ResponseError("05", "params id must be integer"))
	}

	if err := h.service.DeleteDepartement(c, id); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(utils.ResponseError("01", "failed to delete data departement"))
	}

	return ctx.Status(fiber.StatusOK).JSON(utils.ResponseSuccess("00", "success"))
}
