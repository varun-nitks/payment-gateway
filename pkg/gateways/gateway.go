package gateways

// Gateway interface defines methods that all payment gateways must implement
type Gateway interface {
	Deposit(amount float64, accountID string) (string, error)  // Processes deposit and returns transaction ID
	Withdraw(amount float64, accountID string) (string, error) // Processes withdrawal and returns transaction ID
	HandleCallback(data []byte) error                          // Processes asynchronous callbacks from gateway
}
