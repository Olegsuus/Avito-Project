package main

import (
	"Avito-Project/internal/config"
	"Avito-Project/internal/models"
	"Avito-Project/internal/storage"
)

func main() {
	App := &models.App{}
	cfg := config.GetConfig()
	db := storage.GetStorage(cfg)
	migration.Migration(cfg)
}
