package models

import (
	"Avito-Project/internal/config"
	_ "Avito-Project/internal/config"
	"database/sql"
	"fmt"
	"net/http"
)

type App struct {
	Config          *config.Config
	DB              *sql.DB
	ServerInterface http.Handler
}

func (a *App) Start(port int) error {
	return http.ListenAndServe(fmt.Sprintf(":%d", port), a.ServerInterface)
}

func (a *App) Stop() {
	if a.DB != nil {
		a.DB.Close()
	}
}
