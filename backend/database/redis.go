package database

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// RedisDB implements the Database interface for Redis.
type RedisDB struct {
	client *redis.Client
}

// NewRedisDB creates a new instance of RedisDB.
func NewRedisDB(addr, password string, db int) *RedisDB {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &RedisDB{client: client}
}

// SaveURL saves the URL with the given key in the database.
func (r *RedisDB) SaveURL(url, key string) error {
	return r.client.Set(context.Background(), key, url, 0).Err()
}

func (r *RedisDB) GetURL(key string) (string, error) {
	return r.client.Get(context.Background(), key).Result()
}

func (r *RedisDB) IncreaseCount(key string) error {
	return r.client.Incr(context.Background(), key).Err()
}

func (r *RedisDB) StoreAnalytics(key, ip, ua string) error {
	return r.client.LPush(context.Background(), key, ip, ua).Err()
}
