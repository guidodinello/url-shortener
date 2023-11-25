package main

import (
	"log"
	"net/http"
	"url-shortener/backend/database"
	"url-shortener/backend/handlers"
)

func main() {
	// Initialize a RedisDB instance
	//db := database.NewRedisDB("localhost:6379", "", 0)
	// Initialize an InMemoryDB instance
	db := database.NewInMemoryDB()

	// health check
	// curl -v http://localhost:8080/healthcheck
	http.HandleFunc("/healthcheck", handlers.HealthCheckHandler)

	// endpoint to shorten url
	// curl -v -H "Content-Type: application/json" -d '{"Url": "https://www.example.com"}' http://localhost:8080/shorten
	http.HandleFunc("/shorten", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handlers.ShortenURLHandler(w, r, db)
	})

	// endpoint to redirect to original url
	// curl -v -H "Content-Type: application/json" -d '{"shortUrl": "https://shorty.com/86169"}' http://localhost:8080/redirect
	http.HandleFunc("/redirect", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handlers.RedirectURLHandler(w, r, db)
	})

	http.ListenAndServe(":8080", nil)
}
