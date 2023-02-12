package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Egor-Tihonov/Book-network-apigateway/pkg/auth/pb"
	"github.com/Egor-Tihonov/Book-network-apigateway/pkg/models"
	"github.com/labstack/echo/v4"
)

func Registration(c echo.Context, auth pb.AuthServiceClient) error {
	body := models.RegistrationRequest{}

	if err := json.NewDecoder(c.Request().Body).Decode(&body); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	res, err := auth.Registration(context.Background(), &pb.RegistrationRequest{
		Name:     body.Name,
		Username: body.Username,
		Email:    body.Email,
		Password: body.Password,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}
	return c.JSON(int(res.Status), &res)
}

func Login(c echo.Context, auth pb.AuthServiceClient) error {
	body := models.LoginRequest{}

	if err := json.NewDecoder(c.Request().Body).Decode(&body); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	res, err := auth.Login(context.Background(), &pb.LoginRequest{
		Authstring: body.AuthString,
		Password:   body.Password,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}
	return c.JSON(http.StatusOK, &res)
}
