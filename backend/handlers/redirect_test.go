package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"url-shortener/backend/database"
	"url-shortener/backend/models"
)

var key = "abc123"
var expectedRedirectURL = "https://www.example.com"
var domain = "https://shorty.com/"

// Create a mock database
var db = database.NewInMemoryDB(map[string]string{
	key: expectedRedirectURL,
})

func TestRedirectURLHandler(t *testing.T) {
	requestData := models.RequestData{
		Url: fmt.Sprintf("%s%s", domain, key),
	}
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest("POST", "/redirect", bytes.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	RedirectURLHandler(rr, req, db)

	// Check the response status code
	if rr.Code != http.StatusSeeOther {
		t.Errorf("Expected status code %d, got %d", http.StatusSeeOther, rr.Code)
	}

	// Check if the response body contains the correct redirect URL
	if rr.Header().Get("Location") != expectedRedirectURL {
		t.Errorf("Expected redirect URL %s, got %s", expectedRedirectURL, rr.Header().Get("Location"))
	}

	// Check if the count of the shortened URL increased
	_, count, err := db.GetURL(key)
	if err != nil {
		t.Errorf("Error getting count: %v", err)
	}
	if count != 1 {
		t.Errorf("Expected count %d, got %d", 1, count)
	}
}

func TestRedirectURLHandler_InvalidURL(t *testing.T) {
	requestData := models.RequestData{
		Url: "invalid-url",
	}
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest("POST", "/redirect", bytes.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	RedirectURLHandler(rr, req, db)

	// Check the response status code
	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, rr.Code)
	}
}

func TestRedirectURLHandler_KeyNotFound(t *testing.T) {
	requestData := models.RequestData{
		Url: fmt.Sprintf("%s%s", domain, "not-real-key"),
	}
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest("POST", "/redirect", bytes.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	RedirectURLHandler(rr, req, db)

	// Check the response status code
	if rr.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, rr.Code)
	}
}
