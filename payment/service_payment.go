package payment

import (
	"rmzstartup/helper"
	model "rmzstartup/model/entity"
	"rmzstartup/repository"
	"strconv"

	midtrans "github.com/veritrans/go-midtrans"
)

type paymentService struct {
	transactionRepo repository.TransactionRepo
	campaignRepo    repository.CampaignRepo
}

type PaymentService interface {
	GetPaymentURL(transaction Transaction, user model.User) (string, error)
	ProcessPayment(input helper.TransactionNotificationInput) error
}

func NewPaymentService(transactionRepo repository.TransactionRepo, campaignRepo repository.CampaignRepo) *paymentService {
	return &paymentService{
		transactionRepo: transactionRepo,
		campaignRepo:    campaignRepo,
	}
}

func (s *paymentService) GetPaymentURL(transaction Transaction, user model.User) (string, error) {
	midclient := midtrans.NewClient()
	midclient.ServerKey = "SB-Mid-server-xqlNy32UiDe9LoW44nSIfLj_"
	midclient.ClientKey = "SB-Mid-client-U9IpDlBuNXnUdqlC"
	midclient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}

	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName: user.Name,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transaction.ID),
			GrossAmt: int64(transaction.Amount),
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}
	return snapTokenResp.RedirectURL, nil
}

func (s *paymentService) ProcessPayment(input helper.TransactionNotificationInput) error {
	transaction_id, _ := strconv.Atoi(input.OrderID)
	transaction, err := s.transactionRepo.GetByTransactionID(transaction_id)
	if err != nil {
		return err
	}
	if input.PaymentType == "credit_card" && input.TransactionStatus == "capture" && input.FraudStatus == "accept" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "settlement" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {
		transaction.Status = "cancelled"
	}

	updatedTransaction, err := s.transactionRepo.Update(transaction)
	if err != nil {
		return err
	}

	campaign, err := s.campaignRepo.FindByID(updatedTransaction.CampaignID)
	if err != nil {
		return err
	}

	if updatedTransaction.Status == "paid" {
		campaign.BackerCount = campaign.BackerCount + 1
		campaign.CurrentAmount = campaign.CurrentAmount + updatedTransaction.Amount

		_, err := s.campaignRepo.Update(campaign)
		if err != nil {
			return err
		}
	}
	return nil
}
