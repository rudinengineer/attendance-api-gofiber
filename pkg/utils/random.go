package utils

import (
	"crypto/rand"
	"math/big"
)

func RandomNumber(length int) (string, error) {
	characters := "0123456789"

	result := make([]byte, length)
	for i := range result {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(characters))))
		if err != nil {
			return "", err
		}
		result[i] = characters[num.Int64()]
	}
	return string(result), nil
}

func RandomCharacters(length int) (string, error) {
	characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	result := make([]byte, length)
	for i := range result {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(characters))))
		if err != nil {
			return "", err
		}
		result[i] = characters[num.Int64()]
	}
	return string(result), nil
}

func RandomCharactersWithNumbers(length int) (string, error) {
	characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	result := make([]byte, length)
	for i := range result {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(characters))))
		if err != nil {
			return "", err
		}
		result[i] = characters[num.Int64()]
	}
	return string(result), nil
}
