package handlers

import (
	"net/http"
	"testing"
	testingUtils "url-shortener/backend/testing"
)

func TestShortenURL(t *testing.T) {
	rr, _, err := testingUtils.TestResponse(testingUtils.LongUrl,
		testingUtils.Request{Method: "POST", Endpoint: "/shorten"},
		ShortenURLHandler)
	if err != nil {
		t.Fatal(err)
	}

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}
}

func TestShortenURL_MissingUrl(t *testing.T) {
	rr, _, err := testingUtils.TestResponse("",
		testingUtils.Request{Method: "POST", Endpoint: "/shorten"},
		ShortenURLHandler)
	if err != nil {
		t.Fatal(err)
	}

	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, rr.Code)
	}
}
