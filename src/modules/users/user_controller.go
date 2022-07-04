package users

import (
	"go-api/src/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService UserService
}

func NewUserController(userService UserService) *userController {
	return &userController{userService}
}

func (handler userController) FindAll(c *gin.Context) {
	user, err := handler.userService.FindAll()

	if err != nil {
		c.JSON(http.StatusNotFound, helper.ApiResponse(http.StatusNotFound, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, helper.ApiResponse(http.StatusOK, "success", user))
}
