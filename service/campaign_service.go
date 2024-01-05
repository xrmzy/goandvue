package service

import (
	"errors"
	model "rmzstartup/model/entity"
	"rmzstartup/repository"

	"github.com/google/uuid"
)

type CampaignService interface {
	GetCampaigns(userID string) ([]model.Campaign, error)
}

type campaignService struct {
	repository repository.CampaignRepo
}

func NewCampaignService(repository repository.CampaignRepo) *campaignService {
	return &campaignService{repository: repository}
}

func (s *campaignService) GetCampaigns(userID string) ([]model.Campaign, error) {
	// parseUserID, err := uuid.Parse(userID)
	// if err != nil {
	// 	return nil, err
	// }
	// if parseUserID != uuid.Nil {
	// 	// campaign, err := s.repository.FindByUserID(userID)
	// 	// if err != nil {
	// 	// 	return campaign, err
	// 	// }
	// 	// return campaign, nil
	// 	return s.repository.FindByUserID(userID)
	// }

	// // campaigns, err := s.repository.FindAll()
	// // if err != nil {
	// // 	return campaigns, err
	// // }
	// // return campaigns, nil
	// return s.repository.FindAll()

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
