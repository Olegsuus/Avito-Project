package main

import (
	"log"

	"Avito-Project/internal/app"
	"Avito-Project/internal/config"
	"Avito-Project/internal/db"
	"Avito-Project/internal/migration"
)

func main() {
	cfg := config.GetConfig()
	db := db.DataBase{}
	db.GetStorage(cfg)
	migration.Migrations(cfg)
	App := app.App{Config: cfg, DB: &db}
	//_ := App.Start()

	srv := &app.Server{}
	App.ServerInterface = srv

	if err := App.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
