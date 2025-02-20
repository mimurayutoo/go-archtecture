package userHandler

import (
	"net/http"
	"practice-api/internal/response"
	userModel "practice-api/internal/user/model"
	userService "practice-api/internal/user/service"
	"strconv"

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
	var newUser userModel.UserInput
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

func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "パラメータidは数値で入力してください"))
		return
	}
	user, err := h.userService.GetUserByID(uint(idUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "user found successfully", user))
}
