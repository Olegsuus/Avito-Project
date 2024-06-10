package app

import (
	"fmt"
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
	GetBannersByTagID(int) ([]models.Banner, error)
	GetBannersByFID(int) ([]models.Banner, error)
	GetAllUsers() ([]models.User, error)
	GetAllBanners() ([]models.Banner, error)
	Stop() error
}

func (a *App) Start() error {
	a.Echo = echo.New()
	a.ServerInterface.GetServer(a)

	addr := fmt.Sprintf(":%d", a.Config.Database.Port)
	log.Printf("Starting server on %s", addr)
	return a.Echo.Start(addr)
}

// Stop закрывает если есть ошибки
func (a *App) Stop() {
	a.DB.Stop()
}
