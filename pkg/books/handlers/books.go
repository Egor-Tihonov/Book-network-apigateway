package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Egor-Tihonov/Book-network-apigateway/pkg/books/pb"
	"github.com/Egor-Tihonov/Book-network-apigateway/pkg/models"
	"github.com/labstack/echo/v4"
)

func CreateBook(c echo.Context, book pb.BookServiceClient) error {
	body := models.Book{}
	if err := json.NewDecoder(c.Request().Body).Decode(&body); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	res, err := book.CreateBook(context.Background(), &pb.CreateBookRequest{
		Title:   body.Title,
		Name:    body.Name,
		Surname: body.Surname,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}

	return c.JSON(int(res.Response.Status), &res)
}

func GetBook(c echo.Context, book pb.BookServiceClient) error {
	name := c.Param("name")
	res, err := book.GetBook(context.Background(), &pb.GetBookRequest{
		Name: name,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}

	return c.JSON(http.StatusOK, &res)
}

func GetAuthor(c echo.Context, book pb.BookServiceClient) error {
	body := models.Author{}
	if err := json.NewDecoder(c.Request().Body).Decode(&body); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	res, err := book.GetAuthor(context.Background(), &pb.GetAuthorRequest{
		Name:    body.Name,
		Surname: body.Surname,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}

	return c.JSON(http.StatusOK, &res)
}
