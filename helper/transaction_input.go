package helper

import model "rmzstartup/model/entity"

type GetCampaignTransactionsInput struct {
	ID   int `uri:"id" binding:"required"`
	User model.User
}

type CreateTransactionInput struct {
	Amount     int `json:"amount" binding:"required"`
	CampaignID int `json:"campaignId" binding:"required"`
	User       model.User
}

type TransactionNotificationInput struct {
	TransactionStatus string `json:"transaction_status"`
	OrderID           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
}
