package bookservice

import (
	"github.com/Egor-Tihonov/Book-network-apigateway/pkg/book-service/service"
	"github.com/Egor-Tihonov/Book-network-apigateway/pkg/config"
	"github.com/labstack/echo/v4"
)

func RegisterHandlers(e *echo.Echo, conf *config.Config) *ServiceClient {
	svc := ServiceClient{
		Client: InitClient(conf),
	}

	e.GET("/books", svc.GetBooks)
	e.GET("/authors", svc.GetAuthors)

	return &svc
}

func (s *ServiceClient) GetBooks(c echo.Context) error {
	return service.GetBooks(c, s.Client)
}

func (s *ServiceClient) GetAuthors(c echo.Context) error {
	return service.GetAuthors(c, s.Client)
}
