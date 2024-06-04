package app

import (
	"Avito-Project/internal/config"
	_ "Avito-Project/internal/config"
	"Avito-Project/internal/db"
	"Avito-Project/internal/server"
	"database/sql"
	"fmt"
	"log"
)

type App struct {
	Config          *config.Config
	DB              Storage
}

func (a *App) Start(port int) error {
	userInfo, err := db.GetUser(a.DB, "token1111")
	if err != nil {
		log.Printf("Failed to get informaition user: %v", err)
		return err
	}

	bannerInfo, err := db.GetBanners(a.DB, 1)
	if err != nil {
		log.Printf("Failed to get informaition banner: %v", err)
	}

	fmt.Printf("User information: %v", userInfo)
	fmt.Printf("Banner information: %v", bannerInfo)
	return nil
}

// Stop закрывает если есть ошибки
func (a *App) Stop() {
	if a.DB != nil {
		a.DB.Close()
	}
}

type Storage interface{
	GetBanner(id int) (model.Banner, error)
	GetBanners(fId int) ([]model.Banner, error)
	GetUser(token string) (model.User, error)
}

1.