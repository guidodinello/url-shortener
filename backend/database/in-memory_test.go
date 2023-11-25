package database

import (
	"testing"
	"url-shortener/backend/models"
)

func TestInMemoryDB(t *testing.T) {
	db := NewInMemoryDB()

	// Test SaveURL and GetURL
	key := "key"
	val := "val"

	if err := db.SaveURL(val, key); err != nil {
		t.Errorf("Error saving key: %v", err)
	}

	retVal, _, err := db.GetURL(key)
	if err != nil {
		t.Errorf("Error getting key: %v", err)
	}

	if retVal != val {
		t.Errorf("Expected val %s, got %s", retVal, val)
	}

	// Test non-existent key
	_, _, err = db.GetURL("nonexistent")
	if err != models.ErrNotFound {
		t.Errorf("Expected ErrNotFound, got %v", err)
	}
}
