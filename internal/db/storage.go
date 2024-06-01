package db

import (
	"Avito-Project/internal/config"
	"database/sql"
	"fmt"
	"log"
)

func GetStorage(cfg *config.Config) *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.DBName)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	return db
}

//func GetUser(db *sql.DB, token string) *Users {
//	//TODO написать функцию которая возвращает данные юзера по токену
//}

//func GetBannary(db *sql.DB, id int) *Bannary{
//	//TODO написать функцию которая возвращает данные по банеру по id
//}
