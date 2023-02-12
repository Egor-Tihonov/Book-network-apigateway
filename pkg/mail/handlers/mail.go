package handlers

import (
	"context"
	"net/http"

	"github.com/Egor-Tihonov/Book-network-apigateway/pkg/mail/pb"
	"github.com/labstack/echo/v4"
)

func SendEmail(c echo.Context, mail pb.SendMailServiceClient) error {
	a := c.Param("authString")
	res, err := mail.SendMail(context.Background(), &pb.SendMailRequest{
		Id: a,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}

	return c.JSON(int(res.Status), nil)
}
