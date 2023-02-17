package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Egor-Tihonov/Book-network-apigateway/pkg/models"
	"github.com/Egor-Tihonov/Book-network-apigateway/pkg/user/pb"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context, uscl pb.UserServiceClient) error {
	userFromJwt := c.Get("user").(*jwt.Token) //why c.Get("user") to get auth header
	claims := userFromJwt.Claims.(jwt.MapClaims)
	idFromJwt := claims["id"].(string)

	res, err := uscl.GetUser(context.Background(), &pb.GetUserRequest{
		Id: idFromJwt,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}
	return c.JSON(200, &res)
}

func UpdateUser(c echo.Context, uscl pb.UserServiceClient) error {
	body := models.UserUpdate{}

	if err := json.NewDecoder(c.Request().Body).Decode(&body); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	userFromJwt := c.Get("user").(*jwt.Token) //why c.Get("user") to get auth header
	claims := userFromJwt.Claims.(jwt.MapClaims)
	idFromJwt := claims["id"].(string)

	res, err := uscl.UpdateUser(context.Background(), &pb.UpdateUserRequest{
		Id:          idFromJwt,
		Name:        &body.Name,
		Username:    &body.Username,
		Status:      &body.Status,
		Newpassword: &body.NewPassword,
		Oldpassword: &body.OldPassword,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}
	return c.JSON(http.StatusOK, &res)
}

func DeleteUser(c echo.Context, uscl pb.UserServiceClient) error {
	userFromJwt := c.Get("user").(*jwt.Token) //why c.Get("user") to get auth header
	claims := userFromJwt.Claims.(jwt.MapClaims)
	idFromJwt := claims["id"].(string)

	res, err := uscl.DeleteUser(context.Background(), &pb.DeleteUserRequest{
		Id: idFromJwt,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}

	return c.JSON(http.StatusOK, &res)
}

func GetOtherUser(c echo.Context, uscl pb.UserServiceClient) error {
	id := c.Param("id")
	res, err := uscl.GetUser(context.Background(), &pb.GetUserRequest{
		Id: id,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func GetNewUsers(c echo.Context, uscl pb.UserServiceClient) error {
	res, err := uscl.GetNewUsers(context.Background(), &pb.GetNewUsersRequest{})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}

	return c.JSON(http.StatusOK, &res)
}

func AddNewSubscription(c echo.Context, uscl pb.UserServiceClient) error {
	userFromJwt := c.Get("user").(*jwt.Token) //why c.Get("user") to get auth header
	claims := userFromJwt.Claims.(jwt.MapClaims)
	idFromJwt := claims["id"].(string)

	id := c.Param("id")
	res, err := uscl.AddNewSubscription(context.Background(), &pb.AddNewSubscriptionRequest{
		Id:     idFromJwt,
		Userid: id,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}

	return c.JSON(http.StatusOK, &res)
}

func DeleteSubscription(c echo.Context, uscl pb.UserServiceClient) error {
	userFromJwt := c.Get("user").(*jwt.Token) //why c.Get("user") to get auth header
	claims := userFromJwt.Claims.(jwt.MapClaims)
	idFromJwt := claims["id"].(string)

	id := c.Param("id")

	res, err := uscl.DeleteSubscription(context.Background(), &pb.DeleteSubscriptionRequest{
		Id:     idFromJwt,
		Userid: id,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}
	return c.JSON(http.StatusOK, &res)
}

func GetPost(c echo.Context, uscl pb.UserServiceClient) error {
	id := c.Param("id")
	res, err := uscl.GetPost(context.Background(), &pb.GetPostRequest{
		Id: id,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}
	return c.JSON(http.StatusOK, &res)
}

func GetLastPosts(c echo.Context, uscl pb.UserServiceClient) error {
	userFromJwt := c.Get("user").(*jwt.Token) //why c.Get("user") to get auth header
	claims := userFromJwt.Claims.(jwt.MapClaims)
	idFromJwt := claims["id"].(string)

	res, err := uscl.GetLastPosts(context.Background(), &pb.GetLastPostsRequest{
		Id: idFromJwt,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}

	return c.JSON(http.StatusOK, &res)
}

func CreatePost(c echo.Context, uscl pb.UserServiceClient) error {
	userFromJwt := c.Get("user").(*jwt.Token) //why c.Get("user") to get auth header
	claims := userFromJwt.Claims.(jwt.MapClaims)
	idFromJwt := claims["id"].(string)

	body := models.Post{}
	if err := json.NewDecoder(c.Request().Body).Decode(&body); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	post := pb.Post{
		Title:         body.Title,
		AuthorName:    body.Name,
		AuthorSurname: body.Surname,
		Content:       body.Content,
	}

	res, err := uscl.CreatePost(context.Background(), &pb.CreatePostRequest{
		Id:      idFromJwt,
		Newpost: &post,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}

	return c.JSON(http.StatusOK, &res)
}
