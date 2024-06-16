package app

import (
	"fmt"
	"github.com/labstack/echo/v4/middleware"
	"log"

	"Avito-Project/internal/config"
	_ "Avito-Project/internal/config"
	"Avito-Project/internal/models"
	"github.com/labstack/echo/v4"
)

type App struct {
	Config          *config.Config
	DB              Storage
	ServerInterface ServerInterface
	Echo            *echo.Echo
}

type Storage interface {
	GetBanner(int) (*models.Banner, error)
	GetUser(string) (*models.User, error)
	GetUserByID(int) (*models.User, error)
	GetBannerByTagID(int) ([]models.Banner, error)
	GetBannerByFID(int) ([]models.Banner, error)
	GetAllUsers() ([]models.User, error)
	GetAllBanners() ([]models.Banner, error)
	Stop() error
	AddUser(*models.User) error
	DeleteUser(int) error
	AddBanner(*models.Banner) error
	DeleteBanner(int) error
	AddAccessLevel(*models.AccessLevel) error
	GetUsersPaginated(int, int) ([]models.User, error)
	GetBannersPaginated(int, int) ([]models.Banner, error)
}

func (a *App) Start() error {
	a.Echo = echo.New()
	a.ServerInterface.GetServer(a)
	a.Echo.Use(middleware.Logger())
	a.Echo.Use(middleware.Recover())

	addr := fmt.Sprintf(":%d", a.Config.Server.Port)
	log.Printf("Starting server on %s", addr)
	return a.Echo.Start(addr)
}

// Stop закрывает если есть ошибки
func (a *App) Stop() {
	if err := a.DB.Stop(); err != nil {
		log.Fatalf("Failed to close database: %v", err)
	}
}
