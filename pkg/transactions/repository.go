package transactions

import (
	"database/sql"
	"fmt"
	"time"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

// InsertTransaction CreateTransaction inserts a new transaction into the database
func (r *Repository) CreateTransaction(txn *Transaction) error {
	query := `INSERT INTO transactions (id, user_id, amount, type, status, gateway, created_at, updated_at)
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := r.DB.Exec(query, txn.ID, txn.UserID, txn.Amount, txn.Type, txn.Status, txn.Gateway, txn.CreatedAt, txn.UpdatedAt)
	//log error
	if err != nil {
		fmt.Printf("error is %s", err.Error())
	}
	return err
}

// UpdateTransactionStatus updates the status of a transaction in the database
func (r *Repository) UpdateTransactionStatus(transactionID string, status TransactionStatus) error {
	query := `UPDATE transactions SET status = $1, updated_at = $2 WHERE id = $3`
	_, err := r.DB.Exec(query, status, time.Now(), transactionID)
	return err
}
