package userservice

import (
	"github.com/Egor-Tihonov/Book-network-apigateway/pkg/config"
	"github.com/Egor-Tihonov/Book-network-apigateway/pkg/user-service/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Cliet pb.UserServiceClient
}

func InitUserClient(conf *config.Config) pb.UserServiceClient {
	cc, err := grpc.Dial(conf.UserService, grpc.WithInsecure())
	if err != nil {
		logrus.Errorf("Could not connect: %w", err)
	}

	return pb.NewUserServiceClient(cc)
}
