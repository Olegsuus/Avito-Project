package migration

import (
	"Avito-Project/internal/config"
	"database/sql"
	"embed"
	"github.com/pressly/goose/v3"
)

//go:embedMigrations/*.sql
var embedMigrations embed.FS

func Migrations(cfg *config.Config) {
	var db *sql.DB
	//setup database

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}
}
