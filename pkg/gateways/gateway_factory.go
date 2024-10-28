package gateways

import (
	"errors"
)

// GatewayFactory is responsible for managing and providing gateway instances
type GatewayFactory struct {
	gateways map[string]Gateway
}

// NewGatewayFactory initializes the factory with supported gateways
func NewGatewayFactory() *GatewayFactory {
	return &GatewayFactory{
		gateways: map[string]Gateway{
			"GatewayA": NewGatewayA(), // Initializes and stores an instance of GatewayA
			"GatewayB": NewGatewayB(), // Initializes and stores an instance of GatewayB
		},
	}
}

// GetGateway returns a gateway instance based on the provided name
func (f *GatewayFactory) GetGateway(name string) (Gateway, error) {
	gateway, exists := f.gateways[name]
	if !exists {
		return nil, errors.New("unsupported gateway")
	}
	return gateway, nil
}
