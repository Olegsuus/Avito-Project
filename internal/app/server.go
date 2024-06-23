package app

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type ServerInterface interface {
	GetServer(*App)
}

type Server struct {
}

// GetServer метод для запуска роутера и обработчика запросов
func (s *Server) GetServer(app *App) {

	authGroup := app.Echo.Group("/auth")
	authGroup.Use(app.JWTMiddleware)

	authGroup.GET("/banner", app.HandleGetBanner)
	authGroup.GET("/bannersByFID", app.HandleGetBannersByFID)
	authGroup.GET("/bannersByTagId", app.HandleGetBannersByTagID)
	authGroup.GET("/banners", app.HandleGetAllBanners)
	authGroup.GET("/user", app.HandleGetUserByToken)
	authGroup.GET("/users", app.HandleGetAllUsers)
	authGroup.GET("/users/paginated", app.HandleGetUsersPaginated)
	authGroup.POST("/addUser", app.HandleAddUser)
	authGroup.POST("/addBanner", app.HandleAddBanner)
	authGroup.DELETE("/delete/user", app.HandleDeleteUser)
	authGroup.DELETE("/delete/banner", app.HandleDeleteBanner)
	authGroup.POST("/addAccessLevel", app.HandleAddAccessLevel)
	authGroup.PUT("/update/user", app.HandleUpdateUser)

	app.Echo.POST("/login", app.HandleLogin)

	// Маршруты без авторизации
	app.Echo.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
}
