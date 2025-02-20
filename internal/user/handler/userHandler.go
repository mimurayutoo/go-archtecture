package userHandler

import (
	"fmt"
	"net/http"
	"practice-api/internal/response"
	userModel "practice-api/internal/user/model"
	userService "practice-api/internal/user/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService userService.IUserService
}

// NewUserHandler ハンドラの初期化
func NewUserHandler(service userService.IUserService) *UserHandler {
	return &UserHandler{userService: service}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	fmt.Println("user作成処理スタート")
	var newUser userModel.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "入力が必要な情報が不足しています"))
		return
	}

	err := h.userService.CreateUser(&newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse(http.StatusCreated, "user created successfully", newUser))
}
