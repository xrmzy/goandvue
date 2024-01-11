package service

import (
	"errors"
	"fmt"
	"rmzstartup/helper"
	model "rmzstartup/model/entity"
	"rmzstartup/repository"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
)

type CampaignService interface {
	GetCampaigns(userID string) ([]model.Campaign, error)
	GetCampaignByID(input helper.GetCampaignDetailInput) (model.Campaign, error)
	CreateCampaign(input helper.CreateCampaignInput) (model.Campaign, error)
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

func (s *campaignService) CreateCampaign(input helper.CreateCampaignInput) (model.Campaign, error) {
	campaign := model.Campaign{}
	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.GoalAmount = input.GoalAmount
	campaign.Perks = input.Perks
	campaign.UserID = input.User.Id.String()

	slugCandidate := fmt.Sprintf("%s %s", input.Name, input.User.Id.String())
	campaign.Slug = slug.Make(slugCandidate)

	newCampaign, err := s.repository.Save(campaign)
	if err != nil {
		return newCampaign, err
	}
	return newCampaign, nil
}
