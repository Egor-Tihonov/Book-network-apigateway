package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

var IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey:              []byte("MY-SUPER-KEY"),
	TokenLookup:             "cookie:token",
	ErrorHandlerWithContext: JWTErrorChecker,
})

func JWTErrorChecker(err error, c echo.Context) error {
	logrus.Errorf("Unuatharized, %w", err.Error())
	c.SetCookie(&http.Cookie{
		Name:   "token",
		Path:   "/",
		Value:  "",
		MaxAge: -1,
	})
	return echo.NewHTTPError(401)
}
