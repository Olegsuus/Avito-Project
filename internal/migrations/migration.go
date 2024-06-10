package migration

import (
	"embed"
	"log"

	"Avito-Project/internal/config"
	"database/sql"
	"github.com/pressly/goose/v3"
)

//go:embedMigrations/*.sql
var embedMigrations embed.FS

func Migrations(cfg *config.Config) {
	var db *sql.DB
	//setup database

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("Failed to set goose dialect: %v", err)
	}

	if err := goose.Up(db, "."); err != nil {
		log.Fatalf("Failed to apply migrations: %v", err)
	}
}
