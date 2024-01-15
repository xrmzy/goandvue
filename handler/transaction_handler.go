package handler

import (
	"net/http"
	"rmzstartup/helper"
	model "rmzstartup/model/entity"
	"rmzstartup/service"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	service service.TransactionService
}

func NewTransactionHandler(service service.TransactionService) *transactionHandler {
	return &transactionHandler{service: service}
}

func (h *transactionHandler) GetCampaignTransactions(c *gin.Context) {
	var input helper.GetCampaignTransactionsInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("failed to get transaction campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(model.User)
	input.User = currentUser

	transactions, err := h.service.GetTransactionsByCampaignID(input)
	if err != nil {
		response := helper.APIResponse("failed to get transaction campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("User's transactions", http.StatusOK, "success", helper.FormatCampaignTransactions(transactions))
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) GetUserTransactions(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(model.User)
	userID := currentUser.Id.String()

	transactions, err := h.service.GetTransactionsByUserID(userID)
	if err != nil {
		response := helper.APIResponse("failed to get user's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("campaign's transactions", http.StatusOK, "success", helper.FormatUserTransactions(transactions))
	c.JSON(http.StatusOK, response)
}
