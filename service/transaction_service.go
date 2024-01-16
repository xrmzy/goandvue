package service

import (
	"errors"
	"rmzstartup/helper"
	model "rmzstartup/model/entity"
	"rmzstartup/payment"
	"rmzstartup/repository"
)

type transactionService struct {
	repository     repository.TransactionRepo
	campaignRepo   repository.CampaignRepo
	paymentService payment.PaymentService
}

type TransactionService interface {
	GetTransactionsByCampaignID(input helper.GetCampaignTransactionsInput) ([]model.Transaction, error)
	GetTransactionsByUserID(userID string) ([]model.Transaction, error)
	CreateTransaction(input helper.CreateTransactionInput) (model.Transaction, error)
}

func NewTransactionService(repository repository.TransactionRepo, campaignRepo repository.CampaignRepo, paymentService payment.PaymentService) *transactionService {
	return &transactionService{
		repository:     repository,
		campaignRepo:   campaignRepo,
		paymentService: paymentService,
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

func (s *transactionService) CreateTransaction(input helper.CreateTransactionInput) (model.Transaction, error) {
	transaction := model.Transaction{}
	transaction.CampaignID = input.CampaignID
	transaction.Amount = input.Amount
	transaction.UserID = input.User.Id.String()
	transaction.Status = "pending"
	transaction.Code = ""

	newTransaction, err := s.repository.Save(transaction)
	if err != nil {
		return newTransaction, err
	}

	paymentTransaction := payment.Transaction{
		ID:     newTransaction.ID,
		Amount: newTransaction.Amount,
	}

	paymentURL, err := s.paymentService.GetPaymentURL(paymentTransaction, input.User)
	if err != nil {
		return newTransaction, err
	}

	newTransaction.PaymentURL = paymentURL

	newTransaction, err = s.repository.Update(newTransaction)
	if err != nil {
		return newTransaction, err
	}
	return newTransaction, nil
}
