#***Payment Gateway Microservice***#                                                                                         
This project implements a payment gateway microservice in Go, designed to integrate with multiple payment gateways while providing essential functionalities such as deposits, withdrawals, and asynchronous callbacks. The service is built with a focus on maintainability, extensibility, and robustness.

**Table of Contents**
1. Architecture
2. Project Structure
3. API Endpoints
4. Local Setup
    1. Install PostgreSQL
    2. Create Database
    3. Create Tables
5. Testing the Service

**Architecture**
The architecture of the payment gateway microservice follows a layered approach:

**API Layer:** Handles incoming requests and routes them to the appropriate service.
**Service Layer:** Contains business logic for processing payments.
**Repository Layer:** Interacts with the database to perform CRUD operations.
**Integration Layer:** Manages communication with external payment gateways.
**Configuration Management:** Loads environment-specific configurations for seamless switching between local and cloud setups.

**Project Structure**
payment-gateway/
├── cmd/
│   └── main.go                # Entry point of the application
├── config/
│   └── config.go              # Configuration loading
├── internal/
│   └── retry.go                # Retry logic
├── pkg/
│   ├── database/
│   │   └── database.go         # Database connection and utilities
│   ├── gateways/
│   │   ├── gateway.go          # General gateway logic
│   │   ├── gatewayA.go         # Specific logic for Gateway A
│   │   └── gatewayB.go         # Specific logic for Gateway B
│   ├── handlers/
│   │   └── payment_handlers.go  # HTTP handlers for payment operations
│   ├── middleware/
│   │   ├── auth_middleware.go  # Middleware for authentication
│   │   ├── errorhandling.go    # Middleware for error handling
│   │   └── logging.go          # Middleware for logging
│   └── transactions/
│       ├── handler.go          # HTTP handlers for transactions
│       ├── repository.go       # Database interaction logic for transactions
│       ├── service.go          # Business logic for transactions
│       └── transactions.go     # Transaction management
├── test/
│   └── unit_test.go        # Tests for database package
├── .env                        # Environment variables
├── Dockerfile                  # Docker configuration
└── go.mod                      # Go module file

**API Endpoints**
1. Deposit
Endpoint: POST /api/deposit
Description: Initiates a deposit transaction.
    Request Body:
    {
        "amount": 100.00,
        "paymentMethod": "credit_card"
    }
    Response:
    {
        "transactionId": "123456",
        "status": "success"
    }
2. Withdraw
Endpoint: POST /api/withdraw
Description: Initiates a withdrawal transaction.
    Request Body:
    {
        "amount": 50.00,
        "paymentMethod": "bank_transfer"
    }
    Response:
    {
        "transactionId": "654321",
        "status": "success"
    }
3. Callback
Endpoint: POST /api/callback
Description: Receives asynchronous callbacks from payment gateways.
    Request Body:
    {
        "transactionId": "123456",
        "status": "completed"
    }
    Response:
    {
        "message": "Callback received successfully."
    }

**Local Setup
Install PostgreSQL**
1.Download PostgreSQL: Visit the PostgreSQL official site and download the installer for your operating system.
2.Install PostgreSQL: Follow the installation instructions specific to your OS.
3.Start PostgreSQL: Use the command line or GUI tools like pgAdmin to start your PostgreSQL server.
**Create Database**
1.Open your terminal or command prompt.
2.Connect to PostgreSQL:
    psql -U postgres
3.Create a new database:
    CREATE DATABASE payment_gateway;
**Create Tables**
1.Connect to your newly created database:
    \c payment_gateway;
2.Create tables for transactions (example structure):
   CREATE TABLE transactions (
      id VARCHAR(20) PRIMARY KEY,
      user_id INT NOT NULL,
      amount NUMERIC NOT NULL,
      type VARCHAR(20) NOT NULL,
      status VARCHAR(20) NOT NULL,
      gateway VARCHAR(20) NOT NULL,
      created_at TIMESTAMPTZ DEFAULT NOW(),
      updated_at TIMESTAMPTZ DEFAULT NOW()
  );
    
**Testing the Service**
1.Clone the Repository:
    git clone <repository-url>
    cd payment-gateway
2.Install Dependencies:
    go mod tidy
    go mod vendor
3.Set Environment Variables: Create a .env file in the root directory with the following content:
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=postgres
    DB_PASSWORD=yourpassword
    DB_NAME=payment_gateway
4.Run the Service:
    go run main/main.go
5.Test API Endpoints: Use tools like Postman or curl to test the API endpoints.

Example of Testing with curl
Deposit:
    curl -X POST http://localhost:8080/api/deposit -H "Content-Type: application/json" -d '{"amount": 100.00, "type": "gatewayA", "userId": 12345}'
Withdraw:
    curl -X POST http://localhost:8080/api/withdraw -H "Content-Type: application/json" -d '{"amount": 50.00, "type": "gatewayA", "userId": 12345}'

**Future Scopes**
Run the application on cloud server
Dockerisation part done, but due to some error we will be hosting our service on cloud later



