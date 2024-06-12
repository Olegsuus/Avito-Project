package app

import (
	"Avito-Project/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// HandleGetUser метод для получения юзера через
func (a *App) HandleGetUser(c echo.Context) error {
	token := c.QueryParam("token")
	user, err := a.DB.GetUser(token)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}

func (a *App) HandleGetUserById(c echo.Context) error {
	idStr := c.QueryParam("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	user, err := a.DB.GetUserByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, user)
}

// HandleGetBanner метод для получения баннера через
func (a *App) HandleGetBanner(c echo.Context) error {
	idStr := c.QueryParam("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	banner, err := a.DB.GetBanner(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, banner)
}

// HandleGetBannersByFID метод для получения баннера по фичи
func (a *App) HandleGetBannersByFID(c echo.Context) error {
	fIdStr := c.QueryParam("Fid")
	fId, err := strconv.Atoi(fIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	banner, err := a.DB.GetBannersByFID(fId)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, banner)

}

// HandleGetBannersByTagID метод для получения баннера по фичи
func (a *App) HandleGetBannersByTagID(c echo.Context) error {
	tagID, err := strconv.Atoi(c.Param("tag_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	banners, err := a.DB.GetBannersByTagID(tagID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, banners)
}

// HandleGetAllBanners метод для получения всех баннеров
func (a *App) HandleGetAllBanners(c echo.Context) error {
	banners, err := a.DB.GetAllBanners()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, banners)
}

// HandleGetAllUsers метод для получения всех баннеров
func (a *App) HandleGetAllUsers(c echo.Context) error {
	users, err := a.DB.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, users)
}

// HandleAddUser метод для добавления юзера через
func (a *App) HandleAddUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	err := a.DB.AddUser(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, user)
}

// HandleAddBanner метод для добавления баннера
func (a *App) HandleAddBanner(c echo.Context) error {
	var banner models.Banner
	if err := c.Bind(&banner); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	err := a.DB.AddBanner(&banner)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, banner)
}

// HandleAddAccessLevel метод для добавления уровня доступа
func (a *App) HandleAddAccessLevel(c echo.Context) error {
	var level models.AccessLevel

	if err := c.Bind(&level); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	err := a.DB.AddAccessLevel(&level)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, level)
}

// HandleDeleteUser метод для удаления юзера
func (a *App) HandleDeleteUser(c echo.Context) error {
	userId, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	err = a.DB.DeleteUser(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "success"})
}

// HandleDeleteBanner метод для удаления баннера
func (a *App) HandleDeleteBanner(c echo.Context) error {
	bannerId, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	err = a.DB.DeleteBanner(bannerId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "success"})
}
