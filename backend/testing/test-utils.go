package testingUtils

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"url-shortener/backend/database"
	"url-shortener/backend/models"
)

var Key = "abc123"
var LongUrl = "https://www.example.com"
var Domain = "https://shorty.com/"

// Create a mock database
var MockDB = database.NewInMemoryDB(map[string]string{
	Key: LongUrl,
})

type EndpointHandler func(w http.ResponseWriter, r *http.Request, db database.Database)

type Request struct {
	Method   string
	Endpoint string
}

func TestResponse(s string, ri Request, handler EndpointHandler) (*httptest.ResponseRecorder, *models.ResponseData, error) {
	requestData := models.RequestData{Url: s}
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return nil, nil, err
	}
	req := httptest.NewRequest(ri.Method, ri.Endpoint, bytes.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler(rr, req, MockDB)

	var responseData models.ResponseData
	if err := json.NewDecoder(rr.Body).Decode(&responseData); err != nil {
		return nil, nil, err
	}

	return rr, &responseData, nil
}
