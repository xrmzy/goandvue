package service

import (
	"errors"
	"rmzstartup/helper"
	model "rmzstartup/model/entity"
	"rmzstartup/repository"

	"github.com/google/uuid"
)

type CampaignService interface {
	GetCampaigns(userID string) ([]model.Campaign, error)
	GetCampaignByID(input helper.GetCampaignDetailInput) (model.Campaign, error)
}

type campaignService struct {
	repository repository.CampaignRepo
}

func NewCampaignService(repository repository.CampaignRepo) *campaignService {
	return &campaignService{repository: repository}
}

func (s *campaignService) GetCampaigns(userID string) ([]model.Campaign, error) {
	if userID == "" {
		return s.repository.FindAll()
	}
	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}
	if parsedUserID == uuid.Nil {
		return nil, errors.New("Invalid user ID")
	}
	return s.repository.FindByUserID(userID)
}

func (s *campaignService) GetCampaignByID(input helper.GetCampaignDetailInput) (model.Campaign, error) {
	campaign, err := s.repository.FindByID(input.ID)
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}
