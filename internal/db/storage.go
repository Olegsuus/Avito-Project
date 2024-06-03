package db

import (
	"Avito-Project/internal/config"
	"Avito-Project/internal/models"
	"database/sql"
	"fmt"
	"log"
)

// GetStorage функция для подключения к Базе Данных
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

// GetUser функция для получения данных юзера по токкену
func GetUser(db *sql.DB, token string) (*models.User, error) {
	var user models.User
	query := "SELECT id, name, access_level, created_at, updated_at, token FROM Users WHERE token = $1"
	row := db.QueryRow(query, token)

	err := row.Scan(&user.Id, &user.Name, &user.AccessLevels, &user.CreatedAt, &user.UpdatedAt, &user.Token)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Printf("Failed to scan row: %v", err)
		return nil, err
	}

	return &user, nil
}

// GetBanners функция для возврата данных баннера по id
func GetBanners(db *sql.DB, id int) (*models.Banner, error) {
	var banner models.Banner
	query := "SELECT id, title, text, url, created_at, updated_at, owner_id, f_id FROM Banners WHERE id = $1"
	row := db.QueryRow(query, id)

	err := row.Scan(&banner.Id, &banner.Title, &banner.Text, &banner.Url, &banner.CreatedAt, &banner.UpdatedAt, &banner.OwnerId, &banner.FId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		log.Printf("Failed to scan row: %v", err)
		return nil, err
	}
	return &banner, nil
}
