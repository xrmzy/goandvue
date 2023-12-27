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
	var input helper.RegisterUserInput

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

func (h *userHandler) Login(c *gin.Context) {
	var input helper.LoginInputUser
	if err := c.ShouldBindJSON(&input); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Invalid input data", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := helper.FormatUser(loggedUser, "token")
	response := helper.APIResponse("Login Success", http.StatusOK, "OK", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmailAvalaible(c *gin.Context) {
	var input helper.CheckEmailInput
	if err := c.ShouldBindJSON(&input); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := h.userService.CheckEmailAvalaible(input)
	if err != nil {
		errorMessage := gin.H{"error": "Server Error"}
		response := helper.APIResponse("Email checking failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}
	metaMessage := "Email has been registered"
	if isEmailAvailable {
		metaMessage = "Email is Available"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)

}

func (s *userHandler) UploadAvatar(c *gin.Context) {

}

func NewUserHandler(userService service.UserService) *userHandler {
	return &userHandler{userService: userService}
}
