package transactions

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

// Handler represents the HTTP handler
type Handler struct {
	TransactionService *TransactionService
}

// NewHandler initializes a new Handler with routes
func NewHandler(service *TransactionService) *Handler {
	return &Handler{TransactionService: service}
}

// Routes sets up the routes for the handler
func (h *Handler) Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/deposit", h.HandleDeposit)
	r.Post("/withdrawal", h.HandleWithdrawal)
	r.Post("/callback", h.HandleCallback)
	return r
}

// HandleDeposit handles a deposit request
func (h *Handler) HandleDeposit(w http.ResponseWriter, r *http.Request) {
	var req CreateTransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	// Process the deposit through the service layer
	txnID, err := h.TransactionService.ProcessDeposit(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the transaction ID
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"transaction_id": txnID})
}

// HandleWithdrawal handles a withdrawal request
func (h *Handler) HandleWithdrawal(w http.ResponseWriter, r *http.Request) {
	var req CreateTransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	// Process the withdrawal through the service layer
	txnID, err := h.TransactionService.ProcessWithdrawal(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the transaction ID
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"transaction_id": txnID})
}

// HandleCallback handles asynchronous callbacks for transaction updates
func (h *Handler) HandleCallback(w http.ResponseWriter, r *http.Request) {
	var callback CallbackRequest
	if err := json.NewDecoder(r.Body).Decode(&callback); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	// Validate the callback data
	if callback.TransactionID == "" || callback.Status == "" {
		http.Error(w, "invalid callback data", http.StatusBadRequest)
		return
	}

	// Process the callback, updating transaction status based on callback data
	err := h.TransactionService.Repository.UpdateTransactionStatus(callback.TransactionID, TransactionStatus(callback.Status))
	if err != nil {
		http.Error(w, "failed to update transaction status", http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusOK)
}

type CallbackRequest struct {
	TransactionID string `json:"transaction_id"`
	Status        string `json:"status"`
}
