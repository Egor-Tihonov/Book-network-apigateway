package mail

import (
	"github.com/Egor-Tihonov/Book-network-apigateway/pkg/config"
	"github.com/Egor-Tihonov/Book-network-apigateway/pkg/mail/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.SendMailServiceClient
}

func InitMailClient(c *config.Config) pb.SendMailServiceClient {
	cc, err := grpc.Dial(c.MailService, grpc.WithInsecure())

	if err != nil {
		logrus.Errorf("Could not connect: %w", err)
	}

	return pb.NewSendMailServiceClient(cc)
}
