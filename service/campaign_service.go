package service

import (
	model "rmzstartup/model/entity"
	"rmzstartup/repository"

	"github.com/google/uuid"
)

type CampaignService interface {
	FindCampaigns(userID string) ([]model.Campaign, error)
}

type campaignService struct {
	repository repository.CampaignRepo
}

func NewCampaignService(repository repository.CampaignRepo) *campaignService {
	return &campaignService{repository: repository}
}

func (s *campaignService) FindCampaigns(userID string) ([]model.Campaign, error) {
	if userID != uuid.Nil.String() {
		campaign, err := s.repository.FindByUserID(userID)
		if err != nil {
			return campaign, err
		}
		return campaign, nil
	}

	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}
