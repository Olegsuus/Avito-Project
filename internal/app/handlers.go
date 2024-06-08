package app

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (a *App) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")

	user, err := a.DB.GetUser(token)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (a *App) HandleGetBanner(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid banner ID", http.StatusBadRequest)
		return
	}

	banner, err := a.DB.GetBanner(id)
	if err != nil {
		http.Error(w, "Banner not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(banner)
}

func (a *App) HandleGetBannersByFID(w http.ResponseWriter, r *http.Request) {
	FIDStr := r.URL.Query().Get("FId")

	fID, err := strconv.Atoi(FIDStr)
	if err != nil {
		http.Error(w, "Invalid feature ID", http.StatusBadRequest)
		return
	}

	banner, err := a.DB.GetBannersByFID(fID)
	if err != nil {
		http.Error(w, "Banner not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(banner)
}
