package utils

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(password string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(result), err
}

func VerifyPassword(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
