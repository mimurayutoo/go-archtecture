package router

import (
	userHandler "practice-api/internal/user/handler"

	"github.com/gin-gonic/gin"
)

func SetUpUserRoutes(userHandler *userHandler.UserHandler) *gin.Engine {
	r := gin.Default()

	// User routes
	// r.GET("/users", userHandler.GetUsers)
	// r.GET("/users/:id", userHandler.GetUser)
	r.POST("/user", userHandler.CreateUser)
	// r.PUT("/users/:id", userHandler.UpdateUser)
	// r.DELETE("/users/:id", userHandler.DeleteUser)

	return r
}
