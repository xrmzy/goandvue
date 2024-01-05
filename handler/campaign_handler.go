package handler

import (
	"net/http"
	"rmzstartup/helper"
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
			response := helper.APIResponse("Error get campaigns", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		response := helper.APIResponse("List of campaign", http.StatusOK, "success", campaigns)
		c.JSON(http.StatusOK, response)
		return
	}

	// FindByID()
	campaigns, err := h.service.GetCampaigns(userID)
	if err != nil {
		response := helper.APIResponse("Error get campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("List of campaign", http.StatusOK, "success", campaigns)
	c.JSON(http.StatusOK, response)
}
