package main

import (
	"Avito-Project/internal/config"
	"Avito-Project/internal/db"
	"Avito-Project/internal/migration"
)

func main() {
	cfg := config.GetConfig()
	db := db.DataBase{}
	db.GetStorage(cfg)
	migration.Migrations(cfg)

}
