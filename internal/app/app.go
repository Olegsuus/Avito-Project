package app

import (
	"fmt"
	"log"
	"net/http"

	"Avito-Project/internal/config"
	_ "Avito-Project/internal/config"
	"Avito-Project/internal/models"
)

type App struct {
	Config          *config.Config
	DB              Storage
	ServerInterface ServerInterface
}

type Storage interface {
	GetBanner(int) (*models.Banner, error)
	GetUser(string) (*models.User, error)
	GetBannersByTagID(int) ([]models.Banner, error)
	GetBannersByFID(int) ([]models.Banner, error)
	Stop() error
}

func (a *App) Start() error {
	// todo прямо вот тут надо добавить добавление роутов и обработчиков
	// обработчики сделать методами App. Написать в отдельном файле handler.go
	mux := a.ServerInterface.GetServer(a)

	//todo: понять почему не передается порт с кофнига
	addr := fmt.Sprintf(":%d", 8082)
	log.Printf("Starting server on %s", addr)
	return http.ListenAndServe(addr, mux)
}

// Stop закрывает если есть ошибки
func (a *App) Stop() {
	a.DB.Stop()
}
