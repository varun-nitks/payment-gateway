package transactions

import (
	"time"
)

// TransactionType represents either deposit or withdrawal
type TransactionType string

const (
	Deposit    TransactionType = "DEPOSIT"
	Withdrawal TransactionType = "WITHDRAWAL"
)

// TransactionStatus represents the status of a transaction
type TransactionStatus string

const (
	Pending   TransactionStatus = "PENDING"
	Completed TransactionStatus = "COMPLETED"
	Failed    TransactionStatus = "FAILED"
)

// Transaction represents a single transaction
type Transaction struct {
	ID        string            `json:"id"`
	Type      TransactionType   `json:"type"`
	Amount    float64           `json:"amount"`
	Status    TransactionStatus `json:"status"`
	Gateway   string            `json:"gateway"`
	UserID    int               `json:"user_id"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}

// CreateTransactionRequest represents the incoming request for a new transaction
type CreateTransactionRequest struct {
	Type    TransactionType `json:"type"`
	Amount  float64         `json:"amount"`
	UserID  int             `json:"user_id"`
	Gateway string          `json:"gateway"`
}
