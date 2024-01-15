package repository

import (
	model "rmzstartup/model/entity"

	"gorm.io/gorm"
)

type transactionRepo struct {
	db *gorm.DB
}

type TransactionRepo interface {
	GetByCampaignID(campaignID int) ([]model.Transaction, error)
	GetByUserID(userID string) ([]model.Transaction, error)
}

func NewTransactionRepo(db *gorm.DB) *transactionRepo {
	return &transactionRepo{db: db}
}

func (r *transactionRepo) GetByCampaignID(campaignID int) ([]model.Transaction, error) {
	var transaction []model.Transaction

	err := r.db.Preload("User").Where("campaign_id = ?", campaignID).Order("id desc").Find(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *transactionRepo) GetByUserID(userID string) ([]model.Transaction, error) {
	var transaction []model.Transaction
	err := r.db.Preload("Campaign.CampaignImages", "campaign_images.is_primary = 1").Where("user_id = ?", userID).Order("id desc").Find(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}
