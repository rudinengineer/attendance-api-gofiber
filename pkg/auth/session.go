package auth

import (
	"absensi-api/dto"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GetSession(ctx *fiber.Ctx) *dto.Session {
	session := ctx.Locals("user")
	if session == nil {
		return &dto.Session{}
	}

	user, ok := session.(*jwt.Token)
	if !ok {
		return &dto.Session{}
	}

	claims, ok := user.Claims.(jwt.MapClaims)
	if !ok {
		return &dto.Session{}
	}

	return &dto.Session{
		EmployeeID:    claims["employee_id"].(int),
		DepartementID: claims["departement_id"].(int),
	}
}
