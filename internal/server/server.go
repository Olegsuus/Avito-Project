package server

import (
	"net/http"

	"Avito-Project/internal/app"
)

type ServerInterface interface {
	GetServer(app *app.App) http.Handler
}

type Server struct {
}
