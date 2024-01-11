package handler

import (
	"net/http"
	"rmzstartup/helper"
	model "rmzstartup/model/entity"
	"rmzstartup/service"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	service service.CampaignService
}

func NewCampaignHandler(service service.CampaignService) *campaignHandler {
	return &campaignHandler{service: service}
}

// api/v1/campaings

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID := c.Query("user_id")

	// FindALL()
	if userID == "" {
		campaigns, err := h.service.GetCampaigns("")
		if err != nil {
			response := helper.APIResponse("Error get campaigns", http.StatusBadRequest, "error", helper.FormatCampaigns(campaigns))
			c.JSON(http.StatusBadRequest, response)
			return
		}
		response := helper.APIResponse("List of campaign", http.StatusOK, "success", helper.FormatCampaigns(campaigns))
		c.JSON(http.StatusOK, response)
		return
	}

	// FindByID()
	campaigns, err := h.service.GetCampaigns(userID)
	if err != nil {
		response := helper.APIResponse("Error get campaigns", http.StatusBadRequest, "error", helper.FormatCampaigns(campaigns))
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("List of campaign", http.StatusOK, "success", helper.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	var input helper.GetCampaignDetailInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	campaignDetail, err := h.service.GetCampaignByID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Campaign detail", http.StatusOK, "success", helper.FormatCampaignDetail(campaignDetail))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	var input helper.CreateCampaignInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create Campaign", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	currentUser := c.MustGet("currentUser").(model.User)

	input.User = currentUser

	newCampaign, err := h.service.CreateCampaign(input)
	if err != nil {
		response := helper.APIResponse("Failed to create campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Successfully create campaign", http.StatusCreated, "success", helper.FormatCampaign(newCampaign))
	c.JSON(http.StatusCreated, response)
}
