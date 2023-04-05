package service

import (
	"context"
	"net/http"

	"github.com/Egor-Tihonov/Book-network-apigateway/pkg/book-service/pb"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func GetBooks(c echo.Context, book pb.BookServiceClient) error {
	books, err := book.GetBooks(context.Background(), &pb.GetBooksRequest{})
	if err != nil {
		logrus.Errorf("api gateway: %w", err.Error())
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}

	return c.JSON(http.StatusOK, books)
}

func GetAuthors(c echo.Context, book pb.BookServiceClient) error {
	books, err := book.GetAuthors(context.Background(), &pb.GetAuthorsRequest{})
	if err != nil {
		logrus.Errorf("api gateway: %w", err.Error())
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}

	return c.JSON(http.StatusOK, books)
}
