// pkg/gateways/gateway_b.go

package gateways

import "fmt"

// GatewayB simulates a payment gateway using SOAP/XML over HTTP
type GatewayB struct{}

// NewGatewayB creates a new instance of GatewayB
func NewGatewayB() *GatewayB {
	return &GatewayB{}
}

func (g *GatewayB) Deposit(amount float64, accountID string) (string, error) {
	// Mock deposit process for GatewayB
	return fmt.Sprintf("GB-DEP-%s", accountID), nil
}

func (g *GatewayB) Withdraw(amount float64, accountID string) (string, error) {
	// Mock withdrawal process for GatewayB
	return fmt.Sprintf("GB-WITH-%s", accountID), nil
}

func (g *GatewayB) HandleCallback(data []byte) error {
	// Mock callback handling for GatewayB
	fmt.Println("GatewayB callback received:", string(data))
	return nil
}
