package middleware

import (
	"absensi-api/internal/config"
	"absensi-api/pkg/utils"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func Auth() fiber.Handler {
	configuration := config.Load()

	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: []byte(configuration.JWT.SecretKey),
		},
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "missing or malformed JWT" {
		return c.Status(fiber.StatusUnauthorized).
			JSON(utils.ResponseError("07", "Unauthorized"))
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(utils.ResponseError("07", "Invalid or expired token"))
}
