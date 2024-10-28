package transactions

import (
	"errors"
	"payment-gateway/pkg/gateways"
	"time"

	"github.com/google/uuid"
)

// TransactionService provides business logic for handling transactions
type TransactionService struct {
	Repository     *Repository
	GatewayFactory *gateways.GatewayFactory
}

// NewTransactionService creates a new instance of TransactionService
func NewTransactionService(repo *Repository, gatewayFactory *gateways.GatewayFactory) *TransactionService {
	return &TransactionService{
		Repository:     repo,
		GatewayFactory: gatewayFactory,
	}
}

// ProcessDeposit handles a deposit transaction request
func (s *TransactionService) ProcessDeposit(req CreateTransactionRequest) (string, error) {
	if req.Amount <= 0 {
		return "", errors.New("invalid deposit amount")
	}

	txn := &Transaction{
		ID:        uuid.New().String(),
		Type:      Deposit,
		Amount:    req.Amount,
		Status:    Pending,
		Gateway:   req.Gateway,
		UserID:    req.UserID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Save transaction in database
	if err := s.Repository.CreateTransaction(txn); err != nil {
		return "", err
	}

	// Select gateway and process deposit
	gateway, err := s.GatewayFactory.GetGateway(req.Gateway)
	if err != nil {
		return "", err
	}

	go func() {
		if _, err := gateway.Deposit(req.Amount, txn.ID); err != nil {
			txn.Status = Failed
		} else {
			txn.Status = Completed
		}
		txn.UpdatedAt = time.Now()
		_ = s.Repository.UpdateTransactionStatus(txn.ID, txn.Status)
	}()

	return txn.ID, nil
}

// ProcessWithdrawal handles a withdrawal transaction request
func (s *TransactionService) ProcessWithdrawal(req CreateTransactionRequest) (string, error) {
	if req.Amount <= 0 {
		return "", errors.New("invalid withdrawal amount")
	}

	txn := &Transaction{
		ID:        uuid.New().String(),
		Type:      Withdrawal,
		Amount:    req.Amount,
		Status:    Pending,
		Gateway:   req.Gateway,
		UserID:    req.UserID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Save transaction in database
	if err := s.Repository.CreateTransaction(txn); err != nil {
		return "", err
	}

	// Select gateway and process withdrawal
	gateway, err := s.GatewayFactory.GetGateway(req.Gateway)
	if err != nil {
		return "", err
	}

	go func() {
		if _, err := gateway.Withdraw(req.Amount, txn.ID); err != nil {
			txn.Status = Failed
		} else {
			txn.Status = Completed
		}
		txn.UpdatedAt = time.Now()
		_ = s.Repository.UpdateTransactionStatus(txn.ID, txn.Status)
	}()

	return txn.ID, nil
}
