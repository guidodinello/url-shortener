package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"url-shortener/backend/models"
)

func TestShortenURL(t *testing.T) {
	payload := models.RequestData{Url: "https://www.example.com"}
	requestBody, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/shorten", bytes.NewReader(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ShortenURLHandler(w, r, db)
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response models.ResponseData
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	if shortenedURL := response.ShortUrl; shortenedURL == "" {
		t.Errorf("handler returned an empty shortened URL")
	}

}
