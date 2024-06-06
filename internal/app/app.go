package app

import (
	"Avito-Project/internal/config"
	_ "Avito-Project/internal/config"
	"Avito-Project/internal/models"
	"Avito-Project/internal/server"
	"fmt"
	"log"
)

type App struct {
	Config          *config.Config
	DB              Storage
	ServerInterface server.ServerInterface
}

type Storage interface {
	GetBanner(int) (models.Banner, error)
	GetUser(string) (models.User, error)
	GetBannersByTagID(int) ([]models.Banner, error)
	GetBannersByFID(int) ([]models.Banner, error)
	Stop()
}

func (a *App) Start(port int) error {
	userInfo, err := a.DB.GetUser("token1111")
	if err != nil {
		log.Printf("Failed to get informaition user: %v", err)
		return err
	}

	bannerInfo, err := a.DB.GetBanner(1)
	if err != nil {
		log.Printf("Failed to get informaition banner: %v", err)
	}

	fmt.Printf("User information: %v", userInfo)
	fmt.Printf("Banner information: %v", bannerInfo)
	return nil
}

// Stop закрывает если есть ошибки
func (a *App) Stop() {
	a.DB.Stop()
}
