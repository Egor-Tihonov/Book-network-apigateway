package user

import (
	"github.com/Egor-Tihonov/Book-network-apigateway/pkg/auth"
	"github.com/Egor-Tihonov/Book-network-apigateway/pkg/config"
	"github.com/Egor-Tihonov/Book-network-apigateway/pkg/user/handlers"
	"github.com/labstack/echo/v4"
)

func RegisterHandlers(e *echo.Echo, conf *config.Config) *ServiceClient {
	svc := ServiceClient{
		Cliet: InitUserClient(conf),
	}
	routes := e.Group("/user")
	routes.Use(auth.IsLoggedIn)
	routes.POST("/post", svc.CreatePost)
	//routes.DELETE("/post", svc.DeletePost)
	routes.GET("/:id", svc.GetOtherUser)
	routes.GET("/", svc.GetUser)
	routes.PUT("/", svc.UpdateUser)
	routes.DELETE("/", svc.DeleteUser)
	routes.POST("/subscriptions/add/:id", svc.AddNewSubscription)
	routes.DELETE("/subscriptions/delete/:id", svc.DeleteSubscription)

	e.GET("/new-users", svc.GetNewUsers)
	e.GET("/post/:id", svc.GetPost)

	return &svc
}

func (s *ServiceClient) CreatePost(c echo.Context) error {
	return handlers.CreatePost(c, s.Cliet)
}

func (s *ServiceClient) GetPost(c echo.Context) error {
	return handlers.GetPost(c, s.Cliet)
}

func (s *ServiceClient) GetUser(c echo.Context) error {
	return handlers.GetUser(c, s.Cliet)
}

func (s *ServiceClient) GetOtherUser(c echo.Context) error {
	return handlers.GetOtherUser(c, s.Cliet)
}

func (s *ServiceClient) UpdateUser(c echo.Context) error {
	return handlers.UpdateUser(c, s.Cliet)
}

func (s *ServiceClient) DeleteUser(c echo.Context) error {
	return handlers.DeleteUser(c, s.Cliet)
}

func (s *ServiceClient) GetNewUsers(c echo.Context) error {
	return handlers.GetNewUsers(c, s.Cliet)
}

func (s *ServiceClient) AddNewSubscription(c echo.Context) error {
	return handlers.AddNewSubscription(c, s.Cliet)
}

func (s *ServiceClient) DeleteSubscription(c echo.Context) error {
	return handlers.DeleteSubscription(c, s.Cliet)
}

func (s *ServiceClient) GetLastPosts(c echo.Context) error {
	return handlers.GetLastPosts(c, s.Cliet)
}
