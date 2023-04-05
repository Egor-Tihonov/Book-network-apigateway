package auth

import (
	"github.com/Egor-Tihonov/Book-network-apigateway/pkg/auth-service/service"
	"github.com/Egor-Tihonov/Book-network-apigateway/pkg/config"
	"github.com/labstack/echo/v4"
)

func RegisterHandlers(e *echo.Echo, conf *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(conf),
	}

	e.POST("/sign-up", svc.Registration)
	e.POST("/login", svc.Login)
	e.POST("/logout", svc.Logout)

	return svc
}

func (s *ServiceClient) Registration(c echo.Context) error {
	return service.Registration(c, s.Client)
}

func (s *ServiceClient) Logout(c echo.Context) error {
	return service.Logout(c, s.Client)
}

func (s *ServiceClient) Login(c echo.Context) error {
	return service.Login(c, s.Client)
}
