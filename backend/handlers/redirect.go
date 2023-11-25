package handlers

import (
	"encoding/json"
	"net/http"
	"url-shortener/backend/database"
	"url-shortener/backend/models"
	"url-shortener/backend/utils"
)

func RedirectURLHandler(w http.ResponseWriter, r *http.Request, db database.Database) {
	// get shortened url from request body
	var m models.RequestData
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := utils.ValidateUrl(m.Url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//key, err := utils.ExtractKeyFromURL(m.Url)
	key := m.Url

	// get original url
	original, _, err := db.GetURL(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// increase count of url
	err = db.IncreaseCount(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// store analytics data
	err = db.StoreAnalytics(key, r.RemoteAddr, r.UserAgent())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// // successfuly redirect to original url
	// http.Redirect(w, r, original, http.StatusSeeOther)

	// return the original url
	responseData := models.ResponseData{
		ShortUrl: key,
		Url:      original,
	}
	jsonResponse, err := json.Marshal(responseData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(jsonResponse); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
