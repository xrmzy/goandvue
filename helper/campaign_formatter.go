package helper

import (
	model "rmzstartup/model/entity"
	"strings"
)

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserID           string `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

func FormatCampaign(campaign model.Campaign) CampaignFormatter {
	campaignFormatter := CampaignFormatter{}
	campaignFormatter.ID = campaign.ID
	campaignFormatter.UserID = campaign.UserID
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ShortDescription = campaign.ShortDescription
	campaignFormatter.GoalAmount = campaign.GoalAmount
	campaignFormatter.CurrentAmount = campaign.CurrentAmount
	campaignFormatter.ImageURL = ""
	campaignFormatter.Slug = campaign.Slug

	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}
	return campaignFormatter
}

func FormatCampaigns(campaigns []model.Campaign) []CampaignFormatter {
	campaignsFormatter := []CampaignFormatter{}
	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}
	return campaignsFormatter
}

type CampaignDetailFormatter struct {
	ID               int                      `json:"id"`
	Name             string                   `json:"name"`
	ShortDescription string                   `json:"shortDescription"`
	ImageUrl         string                   `json:"imageUrl"`
	GoalAmount       int                      `json:"goalAmpount"`
	CurrentAmount    int                      `json:"currentAmount"`
	UserID           string                   `json:"userID"`
	Slug             string                   `json:"slug"`
	Description      string                   `json:"description"`
	User             CampaignUserFormatter    `json:"user"`
	Perks            []string                 `json:"perks"`
	Images           []CampaignImageFormatter `json:"images"`
}

type CampaignUserFormatter struct {
	Name     string `json:"name"`
	ImageUrl string `json:"imageURL"`
}

type CampaignImageFormatter struct {
	ImageUrl  string `json:"imageUrl"`
	IsPrimary bool   `json:"isPrimary"`
}

func FormatCampaignDetail(campaign model.Campaign) CampaignDetailFormatter {
	campaignDetailFormatter := CampaignDetailFormatter{}
	campaignDetailFormatter.ID = campaign.ID
	campaignDetailFormatter.Name = campaign.Name
	campaignDetailFormatter.ShortDescription = campaign.ShortDescription
	campaignDetailFormatter.Description = campaign.Description
	campaignDetailFormatter.ImageUrl = ""
	campaignDetailFormatter.GoalAmount = campaign.GoalAmount
	campaignDetailFormatter.CurrentAmount = campaign.CurrentAmount
	campaignDetailFormatter.UserID = campaign.UserID
	campaignDetailFormatter.Slug = campaign.Slug

	if len(campaign.CampaignImages) > 0 {
		campaignDetailFormatter.ImageUrl = campaign.CampaignImages[0].FileName
	}

	var perks []string
	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}

	user := campaign.User
	campaignUserFormatter := CampaignUserFormatter{}
	campaignUserFormatter.Name = user.Name
	campaignUserFormatter.ImageUrl = user.AvatarFileName

	campaignDetailFormatter.User = campaignUserFormatter
	campaignDetailFormatter.Perks = perks

	images := []CampaignImageFormatter{}

	for _, image := range campaign.CampaignImages {
		campaignImageFormatter := CampaignImageFormatter{}
		campaignImageFormatter.ImageUrl = image.FileName

		isPrimary := false
		if image.IsPrimary == 1 {
			isPrimary = true
		}
		campaignImageFormatter.IsPrimary = isPrimary

		images = append(images, campaignImageFormatter)
	}

	campaignDetailFormatter.Images = images

	return campaignDetailFormatter
}
