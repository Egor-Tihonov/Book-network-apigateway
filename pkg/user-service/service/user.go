package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Egor-Tihonov/Book-network-apigateway/pkg/models"
	"github.com/Egor-Tihonov/Book-network-apigateway/pkg/user-service/pb"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func GetMySubs(c echo.Context, uscl pb.UserServiceClient) error {
	userFromJwt := c.Get("user").(*jwt.Token)
	claims := userFromJwt.Claims.(jwt.MapClaims)
	idFromJwt := claims["id"].(string)

	res, err := uscl.GetMySubscriptions(context.Background(), &pb.GetMySubscriptionsRequest{
		Id: idFromJwt,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func DeletePost(c echo.Context, uscl pb.UserServiceClient) error {
	userFromJwt := c.Get("user").(*jwt.Token)
	claims := userFromJwt.Claims.(jwt.MapClaims)
	idFromJwt := claims["id"].(string)

	postid := c.Param("id")

	res, err := uscl.DeletePost(context.Background(), &pb.DeletePostRequest{
		Userid: idFromJwt,
		Postid: postid,
	})
	if err != nil {
		return echo.NewHTTPError(int(res.Response.Status), err.Error())
	}

	return c.JSON(http.StatusOK, nil)
}

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

func CheckSubs(c echo.Context, uscl pb.UserServiceClient) error {
	userFromJwt := c.Get("user").(*jwt.Token) //why c.Get("user") to get auth header
	claims := userFromJwt.Claims.(jwt.MapClaims)
	idFromJwt := claims["id"].(string)

	id := c.Param("id")

	check, err := uscl.CheckMySubs(context.Background(), &pb.CheckMySubsRequest{
		Myid:   idFromJwt,
		Userid: id,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}

	return c.JSON(http.StatusOK, check)
}

func GetNewUsers(c echo.Context, uscl pb.UserServiceClient) error {
	userFromJwt := c.Get("user").(*jwt.Token) //why c.Get("user") to get auth header
	claims := userFromJwt.Claims.(jwt.MapClaims)
	idFromJwt := claims["id"].(string)

	res, err := uscl.GetNewUsers(context.Background(), &pb.GetNewUsersRequest{
		Id: idFromJwt,
	})
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
		logrus.Errorf("auth service error: %w", err.Error())
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}

	return c.JSON(http.StatusOK, &res)
}

func DeleteSubscription(c echo.Context, uscl pb.UserServiceClient) error {
	fmt.Println("delete subs")
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
	postid := c.Param("id")
	res, err := uscl.GetPost(context.Background(), &pb.GetPostRequest{
		Postid: postid,
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
		logrus.Errorf(err.Error())
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}

	return c.JSON(http.StatusOK, &res)
}

func GetMyFeed(c echo.Context, uscl pb.UserServiceClient) error {
	userFromJwt := c.Get("user").(*jwt.Token) //why c.Get("user") to get auth header
	claims := userFromJwt.Claims.(jwt.MapClaims)
	idFromJwt := claims["id"].(string)

	feed, err := uscl.GetMyFeed(context.Background(), &pb.GetMyFeedRequest{
		Id: idFromJwt,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}

	return c.JSON(http.StatusOK, feed)
}
