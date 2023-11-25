package handlers

import (
	"net/http"
	"testing"
	testingUtils "url-shortener/backend/testing"
)

func TestRedirectURLHandler(t *testing.T) {
	rr, data, err := testingUtils.TestResponse(testingUtils.Key,
		testingUtils.Request{Method: "POST", Endpoint: "/redirect"},
		RedirectURLHandler)
	if err != nil {
		t.Fatal(err)
	}

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}
	if data.Url != testingUtils.LongUrl {
		t.Errorf("Expected redirect URL %s, got %s", testingUtils.LongUrl, data.Url)
	}
	_, count, err := testingUtils.MockDB.GetURL(testingUtils.Key)
	if err != nil {
		t.Fatalf("Error getting count: %v", err)
	}
	if count != 1 {
		t.Errorf("Expected count %d, got %d", 1, count)
	}
}

func TestRedirectURLHandler_KeyNotFound(t *testing.T) {
	rr, _, err := testingUtils.TestResponse("not-real-key",
		testingUtils.Request{Method: "POST", Endpoint: "/redirect"},
		RedirectURLHandler)
	if err != nil {
		t.Fatal(err)
	}

	if rr.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, got %d", http.StatusNotFound, rr.Code)
	}
}
