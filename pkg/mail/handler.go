package mail

import (
	"github.com/Egor-Tihonov/Book-network-apigateway/pkg/config"
	"github.com/Egor-Tihonov/Book-network-apigateway/pkg/mail/handlers"
	"github.com/labstack/echo/v4"
)

func RegisterHandlers(e *echo.Echo, conf *config.Config) *ServiceClient {
	svc := ServiceClient{
		Client: InitMailClient(conf),
	}

	e.POST("/email", svc.SendEmail)

	return &svc
}

func (s *ServiceClient) SendEmail(c echo.Context) error {
	return handlers.SendEmail(c, s.Client)
}
