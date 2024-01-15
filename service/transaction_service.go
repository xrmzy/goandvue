package service

import (
	"errors"
	"rmzstartup/helper"
	model "rmzstartup/model/entity"
	"rmzstartup/repository"
)

type transactionService struct {
	repository   repository.TransactionRepo
	campaignRepo repository.CampaignRepo
}

type TransactionService interface {
	GetTransactionsByCampaignID(input helper.GetCampaignTransactionsInput) ([]model.Transaction, error)
	GetTransactionsByUserID(userID string) ([]model.Transaction, error)
}

func NewTransactionService(repository repository.TransactionRepo, campaignRepo repository.CampaignRepo) *transactionService {
	return &transactionService{
		repository:   repository,
		campaignRepo: campaignRepo,
	}
}

func (s *transactionService) GetTransactionsByCampaignID(input helper.GetCampaignTransactionsInput) ([]model.Transaction, error) {
	campaign, err := s.campaignRepo.FindByID(input.ID)
	if err != nil {
		return []model.Transaction{}, err
	}

	if campaign.UserID != input.User.Id.String() {
		return []model.Transaction{}, errors.New("not an owner of the campaign")
	}

	transaction, err := s.repository.GetByCampaignID(input.ID)
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (s *transactionService) GetTransactionsByUserID(userID string) ([]model.Transaction, error) {
	transactions, err := s.repository.GetByUserID(userID)
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}
