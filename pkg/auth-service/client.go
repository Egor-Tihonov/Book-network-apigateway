package auth

import (
	"github.com/Egor-Tihonov/Book-network-apigateway/pkg/auth-service/pb"
	"github.com/Egor-Tihonov/Book-network-apigateway/pkg/config"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.AuthService, grpc.WithInsecure())

	if err != nil {
		logrus.Errorf("Could not connect: %w", err)
	}

	return pb.NewAuthServiceClient(cc)
}
