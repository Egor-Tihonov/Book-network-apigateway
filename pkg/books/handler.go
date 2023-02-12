package books

import (
	mid "github.com/Egor-Tihonov/Book-network-apigateway/pkg/auth/middleware"
	"github.com/Egor-Tihonov/Book-network-apigateway/pkg/books/handlers"
	"github.com/Egor-Tihonov/Book-network-apigateway/pkg/config"
	"github.com/labstack/echo/v4"
)

func RegisterHandlers(e *echo.Echo, conf config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitBooksClient(&conf),
	}

	e.GET("/book/:name", svc.GetBook, mid.IsLoggedIn)
	e.POST("/new-book", svc.CreateBook, mid.IsLoggedIn)
	e.GET("/author", svc.GetAuthor, mid.IsLoggedIn)

	return svc
}

func (s *ServiceClient) GetBook(c echo.Context) error {
	return handlers.GetBook(c, s.Client)
}

func (s *ServiceClient) CreateBook(c echo.Context) error {
	return handlers.CreateBook(c, s.Client)
}

func (s *ServiceClient) GetAuthor(c echo.Context) error {
	return handlers.GetAuthor(c, s.Client)
}
