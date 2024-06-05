package db

import (
	"Avito-Project/internal/config"
	"Avito-Project/internal/models"
	"database/sql"
	"fmt"
	"log"
)

type DataBase struct {
	DB *sql.DB
}

// GetStorage функция для подключения к Базе Данных
func (db *DataBase) GetStorage(cfg *config.Config) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.DBName)

	var err error
	db.DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.DB.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
}

// GetUser метод для получения данных юзера по токкену
func (db *DataBase) GetUser(token string) (*models.User, error) {
	var user models.User
	query := "SELECT id, name, access_level, created_at, updated_at, token FROM Users WHERE token = $1"
	row := db.DB.QueryRow(query, token)

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

// GetBanner метод для получения данных баннера по id
func (db *DataBase) GetBanner(id int) (*models.Banner, error) {
	var banner models.Banner
	query := "SELECT id, title, text, url, created_at, updated_at, owner_id, f_id FROM Banners WHERE id = $1"
	row := db.DB.QueryRow(query, id)

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

// GetUsersForID метод для получения юзера по id
func (db *DataBase) GetUserForID(id int) (*models.User, error) {
	var user models.User
	query := "SELECT id, name, access_level, created_at, updated_at, token FROM Users WHERE id = $1"
	row := db.DB.QueryRow(query, id)

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

// GetBannersForUserID метод для получения списка баннеров по id пользователя
func (db *DataBase) GetBannersForUserID(userID int) ([]models.Banner, error) {
	var banners []models.Banner
	query := "SELECT id, title, text, url, created_at, updated_at,owner_id, f_id FROM Banners WHERE owner_id = 1$"
	rows, err := db.DB.Query(query, userID)
	if err != nil {
		panic(err)
	}

	defer db.DB.Close()

	for rows.Next() {
		var banner models.Banner
		err := rows.Scan(&banner.Id, &banner.Title, &banner.Text, &banner.Url, &banner.CreatedAt, &banner.UpdatedAt, &banner.OwnerId, &banner.FId)
		if err != nil {
			log.Fatalf("Failed to Scan rows: %v", err)
			return nil, err
		}
		banners = append(banners, banner)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("Row iteration error: %v", err)
		return nil, err
	}

	return banners, nil
}

// GetBannersForFID метод для получение списка баннеров по f_id
func (db *DataBase) GetBannersForFID(fID int) ([]models.Banner, error) {
	var banners []models.Banner
	query := "SELECT id, title, text, url, created_at, updated_at,owner_id, f_id FROM Banners WHERE owner_id = 1$"
	rows, err := db.DB.Query(query, fID)
	if err != nil {
		log.Fatalf("Failed to execute query: %v", err)
		return nil, err
	}

	defer db.DB.Close()

	for rows.Next() {
		var banner models.Banner
		err := rows.Scan(&banner.Id, &banner.Title, &banner.Text, &banner.Url, &banner.CreatedAt, &banner.UpdatedAt, &banner.OwnerId, &banner.FId)
		if err != nil {
			panic(err)
		}
		banners = append(banners, banner)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("Failed iteration rows: %v ", err)
		return nil, err
	}

	return banners, nil
}
