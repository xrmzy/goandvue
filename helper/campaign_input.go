package helper

import model "rmzstartup/model/entity"

type GetCampaignDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateCampaignInput struct {
	Name             string `json:"name" binding:"required"`
	ShortDescription string `json:"shortDescription" binding:"required"`
	Description      string `json:"description" binding:"required"`
	GoalAmount       int    `json:"goalAmount" binding:"required"`
	Perks            string `json:"perks" binding:"required"`
	User             model.User
}
