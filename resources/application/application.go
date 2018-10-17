package application

func ApplicationErrorResponse(code interface{}) map[string]interface{} {
	return map[string]interface{}{"error": code}
}
