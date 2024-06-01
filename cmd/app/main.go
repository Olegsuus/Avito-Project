package main

import (
	"Avito-Project/internal/app"
	"Avito-Project/internal/config"
	"Avito-Project/internal/db"
)

func main() {
	App := &app.App{}
	cfg := config.GetConfig()
	db := db.GetStorage(cfg)
}
