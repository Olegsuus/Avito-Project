package models

import "time"

// Структура пользователя
type User struct {
	Id           uint      `json:"id"`
	Name         string    `json:"name"`
	AccessLevels int       `json:"access_levels"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Token        string    `json:"token"`
}

// Структура баннера
type Banner struct {
	Id        uint      `json:"id"`
	Title     string    `json:"title"`
	Text      string    `json:"text"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	OwnerId   uint      `json:"owner_id"`
	FId       int       `json:"f_id"`
}
