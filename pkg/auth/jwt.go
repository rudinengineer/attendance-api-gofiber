package auth

import (
	"absensi-api/internal/config"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(payload jwt.MapClaims) (string, error) {
	configuration := config.Load()

	expiredTime, err := strconv.Atoi(configuration.JWT.ExpiredTime)
	if err != nil {
		return "", err
	}

	payload["exp"] = time.Now().Add(time.Minute * time.Duration(expiredTime)).Unix()

	generate := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	token, err := generate.SignedString([]byte(configuration.JWT.SecretKey))
	if err != nil {
		return "", err
	}

	return token, nil
}
