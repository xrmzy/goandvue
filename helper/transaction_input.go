package helper

import model "rmzstartup/model/entity"

type GetCampaignTransactionsInput struct {
	ID   int `uri:"id" binding:"required"`
	User model.User
}
