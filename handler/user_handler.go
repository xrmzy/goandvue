package handler

import (
	"fmt"
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

func (h *userHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helper.APIResponse("failed to upload avatar image", http.StatusUnprocessableEntity, "error", data)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	//harusnya dapet dari jwt nanti
	userID := "2dd7e5e5-01c8-4d6d-8cb7-9c9bdde8e061"
	path := fmt.Sprintf("images/%s-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helper.APIResponse("failed to upload avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helper.APIResponse("failed to save avatar image", http.StatusInternalServerError, "success", data)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	data := gin.H{
		"is_uploaded": true,
	}
	response := helper.APIResponse("successfully uploaded avatar image", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func NewUserHandler(userService service.UserService) *userHandler {
	return &userHandler{userService: userService}
}
