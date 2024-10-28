package gateways

import "fmt"

// GatewayA simulates a payment gateway using JSON over HTTP
type GatewayA struct{}

// NewGatewayA creates a new instance of GatewayA
func NewGatewayA() *GatewayA {
	return &GatewayA{}
}

func (g *GatewayA) Deposit(amount float64, accountID string) (string, error) {
	// Mock deposit process for GatewayA
	return fmt.Sprintf("GA-DEP-%s", accountID), nil
}

func (g *GatewayA) Withdraw(amount float64, accountID string) (string, error) {
	// Mock withdrawal process for GatewayA
	return fmt.Sprintf("GA-WITH-%s", accountID), nil
}

func (g *GatewayA) HandleCallback(data []byte) error {
	// Mock callback handling for GatewayA
	fmt.Println("GatewayA callback received:", string(data))
	return nil
}
