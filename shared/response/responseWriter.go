package response

import "github.com/gin-gonic/gin"

func SuccessResponse(statusCode int, message string, data any) *gin.H {
	return &gin.H{
		"status":  statusCode,
		"message": message,
		"data":    data,
	}
}

func ErrorResponse(statusCode int, message string) *gin.H {
	return &gin.H{
		"status":  statusCode,
		"message": message,
	}
}
