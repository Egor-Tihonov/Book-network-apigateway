package auth

import (
	"github.com/Egor-Tihonov/Book-network-apigateway/pkg/auth/handlers"
	"github.com/Egor-Tihonov/Book-network-apigateway/pkg/config"
	"github.com/labstack/echo/v4"
)

func RegisterHandlers(e *echo.Echo, conf config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(&conf),
	}

	routes := e.Group("/auth")
	routes.POST("/sign-up", svc.Registration)
	routes.POST("/login", svc.Login)

	return svc
}

func (s *ServiceClient) Registration(c echo.Context) error {
	return handlers.Registration(c, s.Client)
}

func (s *ServiceClient) Login(c echo.Context) error {
	return handlers.Login(c, s.Client)
}
