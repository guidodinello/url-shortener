package database

// Database is an interface for interacting with the database.
type Database interface {
	SaveURL(url, key string) error
	GetURL(key string) (string, int, error)
	IncreaseCount(key string) error
	StoreAnalytics(key, ip, ua string) error
}
