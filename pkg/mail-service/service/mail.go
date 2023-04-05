package service

import (
	"context"
	"net/http"

	"github.com/Egor-Tihonov/Book-network-apigateway/pkg/mail-service/pb"
	"github.com/labstack/echo/v4"
)

func SendEmail(c echo.Context, mail pb.SendMailServiceClient) error {
	a := c.Param("email")
	res, err := mail.SendMail(context.Background(), &pb.SendMailRequest{
		Email: a,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}

	return c.JSON(int(res.Status), nil)
}
