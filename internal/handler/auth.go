package handler

import (
	"absensi-api/dto"
	"absensi-api/internal/domain"
	"absensi-api/pkg/utils"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	service domain.AuthService
}

func NewAuth(app *fiber.App, service domain.AuthService) {
	handler := &authHandler{
		service: service,
	}

	routes := app.Group("/auth")
	routes.Post("/login", handler.Login)
}

func (h *authHandler) Login(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*15)
	defer cancel()

	// Parse Request
	var req dto.AuthRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(utils.ResponseError("02", "cannot proccess request"))
	}

	// Validation Request
	errorMessages, err := utils.ValidationRequest(req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.ResponseErrorWithData("04", "validation error", errorMessages))
	}

	token, err := h.service.Login(c, req)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(utils.ResponseError("11", "bad credentials"))
	}

	return ctx.Status(fiber.StatusOK).JSON(utils.ResponseSuccessWithData("00", "success", map[string]string{
		"access_token": token,
	}))
}
