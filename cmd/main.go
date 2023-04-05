package main

import (
	"fmt"
	"net/http"
	"runtime"
	"strings"

	auth "github.com/Egor-Tihonov/Book-network-apigateway/pkg/auth-service"
	book "github.com/Egor-Tihonov/Book-network-apigateway/pkg/book-service"
	"github.com/Egor-Tihonov/Book-network-apigateway/pkg/config"
	mail "github.com/Egor-Tihonov/Book-network-apigateway/pkg/mail-service"
	user "github.com/Egor-Tihonov/Book-network-apigateway/pkg/user-service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func main() {
	InitLog()

	c, err := config.LoadConfig()
	if err != nil {
		logrus.Fatalf("failing load configs, %w", err)
	}

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}))

	auth.RegisterHandlers(e, &c)
	user.RegisterHandlers(e, &c)
	book.RegisterHandlers(e, &c)
	mail.RegisterHandlers(e, &c)

	if err := e.Start(c.Port); err != nil {
		logrus.Fatalf("error starting api gateway, %w", err)
	}
}

func InitLog() {
	logrus.SetReportCaller(true)

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		ForceColors:     true,
		DisableColors:   false,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return "", fmt.Sprintf(" %s:%d", formatFilePath(f.File), f.Line)
		},
	})

}

func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}
