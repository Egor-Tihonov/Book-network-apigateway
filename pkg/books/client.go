package books

import (
	"github.com/Egor-Tihonov/Book-network-apigateway/pkg/books/pb"
	"github.com/Egor-Tihonov/Book-network-apigateway/pkg/config"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.BookServiceClient
}

func InitBooksClient(c *config.Config) pb.BookServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.BookService, grpc.WithInsecure())

	if err != nil {
		logrus.Errorf("Could not connect: %w", err)
	}

	return pb.NewBookServiceClient(cc)
}
