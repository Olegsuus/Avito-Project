package app

type ServerInterface interface {
	GetServer(*App)
}

type Server struct {
}

// GetServer метод для запуска роутера и обработчика запросов
func (s *Server) GetServer(app *App) {

	//admin := app.Echo.Group("auth/")
	//admin.Use(auth.Basic)
	app.Echo.GET("/banner", app.HandleGetBanner)
	app.Echo.GET("/banners/paginated", app.HandleGetBannersPaginated)
	app.Echo.GET("/bannersByFID", app.HandleGetBannersByFID)
	app.Echo.GET("/bannersByTagId", app.HandleGetBannersByTagID)
	app.Echo.GET("/banners", app.HandleGetAllBanners)
	app.Echo.GET("/user", app.HandleGetUserByToken)
	app.Echo.GET("/users/paginated", app.HandleGetUsersPaginated)
	app.Echo.GET("/userId", app.HandleGetUserById)
	app.Echo.GET("/users", app.HandleGetAllUsers)
	app.Echo.POST("/addUser", app.HandleAddUser)
	app.Echo.POST("/addBanner", app.HandleAddBanner)
	app.Echo.POST("/addALevel", app.HandleAddAccessLevel)
	app.Echo.DELETE("/deleteUser", app.HandleDeleteUser)
	app.Echo.DELETE("/deleteBanner", app.HandleDeleteBanner)
}
