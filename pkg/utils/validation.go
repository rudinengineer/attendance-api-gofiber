package utils

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidationRequest[T any](req T) (map[string]string, error) {
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	errorMessages := map[string]string{}

	if err := validate.Struct(req); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errorMessages[e.Field()] = fmt.Sprintf("Field %s error on %s", e.Field(), e.Tag())
		}

		return errorMessages, err
	}

	return errorMessages, nil
}
