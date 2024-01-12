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
	UpdateCampaign(inputID helper.GetCampaignDetailInput, inputData helper.CreateCampaignInput) (model.Campaign, error)
	SaveCampaignImage(input helper.CreateCampaignImageInput, fileLocation string) (model.CampaignImage, error)
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

func (s *campaignService) UpdateCampaign(inputID helper.GetCampaignDetailInput, inputData helper.CreateCampaignInput) (model.Campaign, error) {
	campaign, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return campaign, err
	}

	if campaign.UserID != inputData.User.Id.String() {
		return campaign, errors.New("not an owner of the campaign")
	}

	campaign.Name = inputData.Name
	campaign.ShortDescription = inputData.ShortDescription
	campaign.Description = inputData.Description
	campaign.Perks = inputData.Perks
	campaign.GoalAmount = inputData.GoalAmount

	updatedCampaign, err := s.repository.Update(campaign)
	if err != nil {
		return updatedCampaign, err
	}
	return updatedCampaign, nil
}

func (s *campaignService) SaveCampaignImage(input helper.CreateCampaignImageInput, fileLocation string) (model.CampaignImage, error) {
	campaign, err := s.repository.FindByID(input.CampaignID)
	if err != nil {
		return model.CampaignImage{}, err
	}

	if campaign.UserID != input.User.Id.String() {
		return model.CampaignImage{}, errors.New("not an owner of the campaign")
	}

	isPrimary := 0
	if input.IsPrimary {
		isPrimary = 1
		_, err := s.repository.MarkAllImagesAsNonPrimary(input.CampaignID)
		if err != nil {
			return model.CampaignImage{}, err
		}
	}
	campaignImage := model.CampaignImage{}
	campaignImage.CampaignID = input.CampaignID
	campaignImage.IsPrimary = isPrimary
	campaignImage.FileName = fileLocation

	newCampaignImage, err := s.repository.CreateImage(campaignImage)
	if err != nil {
		return newCampaignImage, err
	}
	return newCampaignImage, nil
}
