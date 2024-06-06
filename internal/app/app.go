package app

import (
	"fmt"
	"log"

	"Avito-Project/internal/config"
	_ "Avito-Project/internal/config"
	"Avito-Project/internal/models"
	"Avito-Project/internal/server"
)

type App struct {
	Config          *config.Config
	DB              Storage
	ServerInterface server.ServerInterface
}

type Storage interface {
	GetBanner(int) (*models.Banner, error)
	GetUser(string) (*models.User, error)
	GetBannersByTagID(int) ([]models.Banner, error)
	GetBannersByFID(int) ([]models.Banner, error)
	Stop() error // todo вот тут ты не дописал возвращаемый тип по этому не работало
}

// порт указывается в конфигах.
func (a *App) Start() error {

	// todo прямо вот тут надо добавить добавление роутов и обработчиков
	// обработчики сделать методами App. Написать в отдельном файле handler.go
	userInfo, err := a.DB.GetUser("token1111") // todo в конфиг
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
