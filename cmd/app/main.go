package main

import (
	"Avito-Project/internal/app"
	"Avito-Project/internal/config"
	"Avito-Project/internal/db"
)

func main() {

	cfg := config.GetConfig()
	db := db.GetStorage(cfg) // Должен вернуть кастомный обьект (структура), с методами из интерфейса
	migrate.migrate() // добавить миграцию
	App := &app.App{DB: db, config: cfg}


}
