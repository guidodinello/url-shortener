package database

import (
	"sync"
	"url-shortener/backend/models"
)

// InMemoryDB is a simple in-memory database for testing purposes.
type InMemoryDB struct {
	mu   sync.Mutex
	urls map[string]models.DataModel
}

// NewInMemoryDB creates a new instance of InMemoryDB.
func NewInMemoryDB(initialData ...map[string]string) *InMemoryDB {
	db := &InMemoryDB{
		urls: make(map[string]models.DataModel),
	}

	// If initial data is provided, initialize the database with it.
	if len(initialData) > 0 {
		for key, value := range initialData[0] {
			db.urls[key] = models.DataModel{Url: value}
		}
	}

	return db
}

// SaveURL saves the URL with the given key in the in-memory database.
func (i *InMemoryDB) SaveURL(value, key string) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.urls[key] = models.DataModel{
		Url:   value,
		Count: 0,
	}

	return nil
}

// GetURL retrieves the URL with the given key from the in-memory database.
func (i *InMemoryDB) GetURL(key string) (string, int, error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	data, ok := i.urls[key]
	if !ok {
		return "", 0, models.ErrNotFound
	}

	return data.Url, data.Count, nil
}

// IncreaseCount increases the count for the given key in the in-memory database.
func (i *InMemoryDB) IncreaseCount(key string) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	data, ok := i.urls[key]
	if !ok {
		return models.ErrNotFound
	}

	data.Count++
	i.urls[key] = data

	return nil
}

func (i *InMemoryDB) StoreAnalytics(key, ip, ua string) error {
	return nil
}

func (i *InMemoryDB) Dump() {
	for k, v := range i.urls {
		print(" key: ", k, " url: ", v.Url, " count: ", v.Count, "\n")
	}
}
