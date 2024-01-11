package repository

import (
	model "rmzstartup/model/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CampaignRepo interface {
	FindAll() ([]model.Campaign, error)
	FindByUserID(userID string) ([]model.Campaign, error)
	FindByID(ID int) (model.Campaign, error)
	Save(campaign model.Campaign) (model.Campaign, error)
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

	_, err := uuid.Parse(userID)
	if err != nil {
		return campaigns, err
	}

	err = r.db.Where("user_id = ?", userID).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (r *campaignRepo) FindByID(ID int) (model.Campaign, error) {
	var campaign model.Campaign
	err := r.db.Preload("User").Preload("CampaignImages").Where("id = ?", ID).Find(&campaign).Error

	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (r *campaignRepo) Save(campaign model.Campaign) (model.Campaign, error) {
	err := r.db.Create(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}
