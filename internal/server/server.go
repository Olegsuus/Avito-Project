package server

import (
	"Avito-Project/internal/app"
	"net/http"
)

type ServerInterface interface {
	GetServer(app *app.App) http.Handler
}

type Server struct {
}

func (s *Server) GetServer(a *app.App) http.Handler {
	return nil
}
