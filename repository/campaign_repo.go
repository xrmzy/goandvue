package repository

import (
	model "rmzstartup/model/entity"

	"gorm.io/gorm"
)

type CampaignRepo interface {
	FindAll() ([]model.Campaign, error)
	FindByUserID(userID string) ([]model.Campaign, error)
}

type campaignRepo struct {
	db *gorm.DB
}

func NewCampaignRepo(db *gorm.DB) *campaignRepo {
	return &campaignRepo{db: db}
}

func (r *campaignRepo) FindAll() ([]model.Campaign, error) {
	var campaigns []model.Campaign
	err := r.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (r *campaignRepo) FindByUserID(userID string) ([]model.Campaign, error) {
	var campaigns []model.Campaign
	err := r.db.Where("user_id = ?", userID).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}
