package grpcClient

import (
	"github.com/Shahboz4131/to-do-service/config"
)

// IGrpcClient ...
type IGrpcClient interface{}

// GrpcClient ...
type GrpcClient struct {
	cfg         config.Config
	connections map[string]interface{}
}

// New ...
func New(cfg config.Config) (*GrpcClient, error) {
	return &GrpcClient{
		cfg:         cfg,
		connections: map[string]interface{}{},
	}, nil
}
