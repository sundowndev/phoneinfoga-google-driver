package api

import (
	"strings"
)

func successResponse(msg ...string) JsonResponse {
	var message string

	if len(msg) == 0 {
		message = ""
	} else {
		message = strings.Join(msg, "")
	}

	return JsonResponse{
		Success: true,
		Error:   message,
	}
}

func errorResponse(msg ...string) JsonResponse {
	var message string

	if len(msg) == 0 {
		message = "An error occurred"
	} else {
		message = strings.Join(msg, "")
	}

	return JsonResponse{
		Success: false,
		Error:   message,
	}
}
