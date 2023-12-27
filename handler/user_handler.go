package handler

import (
	"net/http"
	"rmzstartup/helper"
	"rmzstartup/service"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.UserService
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input service.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Invalid Input data", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	user, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := helper.FormatUser(user, "token")
	response := helper.APIResponse("Account has been registered", http.StatusCreated, "success", formatter)
	c.JSON(http.StatusCreated, response)

}

func NewUserHandler(userService service.UserService) *userHandler {
	return &userHandler{userService: userService}
}