package app

type ServerInterface interface {
	GetServer(*App)
}

type Server struct {
}

// GetServer метод для запуска роутера и обработчика запросов
func (s *Server) GetServer(app *App) {
	app.Echo.GET("/banner", app.HandleGetBanner)
	app.Echo.GET("/bannersByFID", app.HandleGetBannersByFID)
	app.Echo.GET("/bannersByTagId", app.HandleGetBannersByTagID)
	app.Echo.GET("/banners", app.HandleGetAllBanners)
	app.Echo.GET("/user", app.HandleGetUser)
	app.Echo.GET("/users", app.HandleGetAllUsers)
	app.Echo.POST("/addUser", app.HandleAddUser)
	app.Echo.POST("/addBanner", app.HandleAddUBanner)
	app.Echo.POST("/addALevel", app.HandleAddAccessLevel)
}
