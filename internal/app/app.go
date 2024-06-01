package app

import (
	"Avito-Project/internal/config"
	_ "Avito-Project/internal/config"
	"Avito-Project/internal/server"
	"database/sql"
)

type App struct {
	Config          *config.Config
	DB              *sql.DB
	ServerInterface server.ServerInterface
}

func (a *App) Start(port int) error {
	//TODO db.GetUser("token")
	//TODO db.GetBanner(id)
	return nil
}

func (a *App) Stop() {
	if a.DB != nil {
		a.DB.Close()
	}
}
