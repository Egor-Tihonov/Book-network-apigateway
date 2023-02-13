package main

import (
	"github.com/Egor-Tihonov/Book-network-apigateway/pkg/auth"
	"github.com/Egor-Tihonov/Book-network-apigateway/pkg/config"
	"github.com/Egor-Tihonov/Book-network-apigateway/pkg/mail"
	"github.com/Egor-Tihonov/Book-network-apigateway/pkg/user"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		logrus.Fatalf("failing load configs, %w", err)
	}

	e := echo.New()

	auth.RegisterHandlers(e, &c)
	mail.RegisterHandlers(e, &c)
	user.RegisterHandlers(e, &c)

	if err := e.Start(c.Port); err != nil {
		logrus.Fatalf("error starting api gateway, %w", err)
	}
}
