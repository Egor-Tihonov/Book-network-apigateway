package userservice

import (
	"github.com/Egor-Tihonov/Book-network-apigateway/pkg/auth-service"
	"github.com/Egor-Tihonov/Book-network-apigateway/pkg/config"
	"github.com/Egor-Tihonov/Book-network-apigateway/pkg/user-service/service"
	"github.com/labstack/echo/v4"
)

func RegisterHandlers(e *echo.Echo, conf *config.Config) *ServiceClient {
	svc := ServiceClient{
		Cliet: InitUserClient(conf),
	}
	routes := e.Group("/user")
	routes.Use(auth.IsLoggedIn)
	routes.POST("/post", svc.CreatePost)
	routes.DELETE("/post/:id", svc.DeletePost)
	routes.GET("/lastposts", svc.GetLastPosts)
	routes.GET("/:id", svc.GetOtherUser)
	routes.GET("/check/:id", svc.CheckMySubs)
	routes.GET("/", svc.GetUser)
	routes.GET("/feed", svc.GetMyFeed)
	routes.PUT("/", svc.UpdateUser)
	routes.DELETE("/", svc.DeleteUser)
	routes.GET("/subscriptions", svc.GetMySubs)
	routes.PUT("/subscriptions/add/:id", svc.AddNewSubscription)
	routes.PUT("/subscriptions/delete/:id", svc.DeleteSubscription)
	routes.GET("/new-users", svc.GetNewUsers)
	
	e.GET("/post/:id", svc.GetPost)

	return &svc
}

func (s *ServiceClient) GetMySubs(c echo.Context) error {
	return service.GetMySubs(c, s.Cliet)
}

func (s *ServiceClient) DeletePost(c echo.Context) error {
	return service.DeletePost(c, s.Cliet)
}

func (s *ServiceClient) CreatePost(c echo.Context) error {
	return service.CreatePost(c, s.Cliet)
}

func (s *ServiceClient) GetPost(c echo.Context) error {
	return service.GetPost(c, s.Cliet)
}

func (s *ServiceClient) GetUser(c echo.Context) error {
	return service.GetUser(c, s.Cliet)
}

func (s *ServiceClient) GetOtherUser(c echo.Context) error {
	return service.GetOtherUser(c, s.Cliet)
}

func (s *ServiceClient) UpdateUser(c echo.Context) error {
	return service.UpdateUser(c, s.Cliet)
}

func (s *ServiceClient) DeleteUser(c echo.Context) error {
	return service.DeleteUser(c, s.Cliet)
}

func (s *ServiceClient) GetNewUsers(c echo.Context) error {
	return service.GetNewUsers(c, s.Cliet)
}

func (s *ServiceClient) AddNewSubscription(c echo.Context) error {
	return service.AddNewSubscription(c, s.Cliet)
}

func (s *ServiceClient) DeleteSubscription(c echo.Context) error {
	return service.DeleteSubscription(c, s.Cliet)
}

func (s *ServiceClient) GetLastPosts(c echo.Context) error {
	return service.GetLastPosts(c, s.Cliet)
}

func (s *ServiceClient) GetMyFeed(c echo.Context) error {
	return service.GetMyFeed(c, s.Cliet)
}

func (s *ServiceClient) CheckMySubs(c echo.Context) error {
	return service.CheckSubs(c, s.Cliet)
}
