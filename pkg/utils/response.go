package utils

func ResponseSuccess(status_code string, message string) map[string]any {
	return map[string]any{
		"status":      true,
		"status_code": status_code,
		"message":     message,
	}
}

func ResponseSuccessWithData[T any](status_code string, message string, data T) map[string]any {
	return map[string]any{
		"status":      true,
		"status_code": status_code,
		"message":     message,
		"data":        data,
	}
}

func ResponseError(status_code string, message string) map[string]any {
	return map[string]any{
		"status":      false,
		"status_code": status_code,
		"message":     message,
	}
}

func ResponseErrorWithData[T any](status_code string, message string, data T) map[string]any {
	return map[string]any{
		"status":      false,
		"status_code": status_code,
		"message":     message,
		"data":        data,
	}
}
