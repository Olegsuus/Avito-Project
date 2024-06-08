package app

import "net/http"

type ServerInterface interface {
	GetServer(*App) http.Handler
}

type Server struct {
}

// GetServer метод для запуска роутера и обработчика запросов
func (s *Server) GetServer(app *App) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/user", app.HandleGetUser)
	mux.HandleFunc("/banner", app.HandleGetBanner)
	mux.HandleFunc("/bannersByFID", app.HandleGetBannersByFID)

	return mux
}
