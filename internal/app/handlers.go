package app

import (
	"Avito-Project/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// HandleGetUser метод на получения юзера через пакет http
//func (a *App) HandleGetUser(w http.ResponseWriter, r *http.Request) {
//	token := r.URL.Query().Get("token")
//
//	user, err := a.DB.GetUser(token)
//	if err != nil {
//		http.Error(w, "User not found", http.StatusNotFound)
//		return
//	}
//
//	json.NewEncoder(w).Encode(user)
//}

// HandleGetUser метод для получения юзера через фреймворк echo
func (a *App) HandleGetUser(c echo.Context) error {
	token := c.QueryParam("token")
	user, err := a.DB.GetUser(token)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"Error": "User not found"})
	}

	return c.JSON(http.StatusOK, user)
}

// HandleGetBanner метод на получения банера через пакет http
//func (a *App) HandleGetBanner(w http.ResponseWriter, r *http.Request) {
//	idStr := r.URL.Query().Get("id")
//
//	id, err := strconv.Atoi(idStr)
//	if err != nil {
//		http.Error(w, "Invalid banner ID", http.StatusBadRequest)
//		return
//	}
//
//	banner, err := a.DB.GetBanner(id)
//	if err != nil {
//		http.Error(w, "Banner not found", http.StatusNotFound)
//		return
//	}
//
//	json.NewEncoder(w).Encode(banner)
//}

// HandleGetBanner метод для получения баннера через фреймворк echo
func (a *App) HandleGetBanner(c echo.Context) error {
	idStr := c.QueryParam("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"Error": "Invalid banner ID"})
	}

	banner, err := a.DB.GetBanner(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"Error": "Banner not found"})
	}

	return c.JSON(http.StatusOK, banner)
}

//HandleGetBannersByFID метод для получения баннера по фичи через пакет http
//func (a *App) HandleGetBannersByFID(w http.ResponseWriter, r *http.Request) {
//	FIDStr := r.URL.Query().Get("FId")
//
//	fID, err := strconv.Atoi(FIDStr)
//	if err != nil {
//		http.Error(w, "Invalid feature ID", http.StatusBadRequest)
//		return
//	}
//
//	banner, err := a.DB.GetBannersByFID(fID)
//	if err != nil {
//		http.Error(w, "Banner not found", http.StatusNotFound)
//		return
//	}
//
//	json.NewEncoder(w).Encode(banner)
//}

// HandleGetBannersByFID метод для получения баннера по фичи через фреймворк echo
func (a *App) HandleGetBannersByFID(c echo.Context) error {
	fIdStr := c.QueryParam("Fid")
	fId, err := strconv.Atoi(fIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"Error": "Invalid F_ID"})
	}

	banner, err := a.DB.GetBannersByFID(fId)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"Error": "Banner not found"})
	}

	return c.JSON(http.StatusOK, banner)

}

// HandleGetBannersByTagID метод для получения баннера по фичи через фреймворк echo
func (a *App) HandleGetBannersByTagID(c echo.Context) error {
	tagID, err := strconv.Atoi(c.Param("tag_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"Error": "Invalid tag ID"})
	}

	banners, err := a.DB.GetBannersByTagID(tagID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"Error": "Banner not found"})
	}

	return c.JSON(http.StatusOK, banners)
}

// HandleGetAllBanners метод для получения всех баннеров через фреймворк echo
func (a *App) HandleGetAllBanners(c echo.Context) error {
	banners, err := a.DB.GetAllBanners()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"Error": "User not found"})
	}

	return c.JSON(http.StatusOK, banners)
}

// HandleGetAllUsers метод для получения всех баннеров через фреймворк echo
func (a *App) HandleGetAllUsers(c echo.Context) error {
	users, err := a.DB.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"Error": "User not found"})
	}

	return c.JSON(http.StatusOK, users)
}

// HandleAddUser метод для добавления юзера через фреймворк echo
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

// HandleAddBanner метод для добавления баннера через фреймворк echo
func (a *App) HandleAddUBanner(c echo.Context) error {
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
