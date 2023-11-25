package handlers

import (
	"encoding/json"
	"net/http"
	"url-shortener/backend/database"
	"url-shortener/backend/models"
	"url-shortener/backend/utils"
)

func ShortenURLHandler(w http.ResponseWriter, r *http.Request, db database.Database) {
	// Decode the JSON request body into a Message struct
	var requestMessage models.RequestData
	if err := json.NewDecoder(r.Body).Decode(&requestMessage); err != nil {
		http.Error(w, models.ApiJsonError(err.Error()), http.StatusBadRequest)
		return
	}

	// Validate the URL
	if requestMessage.Url == "" {
		http.Error(w, models.ApiJsonError(models.ErrEmptyUrl.Error()), http.StatusBadRequest)
		return
	}

	// Shorten the URL (you need to implement this function)
	shortened := utils.ShortenURL(requestMessage.Url, 5)

	// Save the URL
	if err := db.SaveURL(requestMessage.Url, shortened); err != nil {
		http.Error(w, models.ApiJsonError(err.Error()), http.StatusInternalServerError)
		return
	}

	// Create a response map
	responseData := models.ResponseData{Key: shortened}

	// Encode the response data as JSON
	jsonResponse, err := json.Marshal(responseData)
	if err != nil {
		http.Error(w, models.ApiJsonError(err.Error()), http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response to the response writer
	if _, err := w.Write(jsonResponse); err != nil {
		http.Error(w, models.ApiJsonError(err.Error()), http.StatusInternalServerError)
		return
	}
}
