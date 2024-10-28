// main.go

package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"payment-gateway/config"
	"payment-gateway/pkg/database"
	"payment-gateway/pkg/gateways"
	"payment-gateway/pkg/transactions"
)

func main() {
	// Load environment variables from .env file
	_ = godotenv.Load()

	// Load configuration using config package
	cfg := config.LoadConfig()

	// Initialize the PostgreSQL database connection with config
	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Set up the repository and service
	gatewayFactory := gateways.NewGatewayFactory()
	txnRepo := transactions.NewRepository(db)
	txnService := transactions.NewTransactionService(txnRepo, gatewayFactory)

	// Initialize the HTTP handler
	h := transactions.NewHandler(txnService)

	// Set up the HTTP server
	http.Handle("/", h.Routes())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8082" // Default port
	}

	log.Printf("Starting server on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
