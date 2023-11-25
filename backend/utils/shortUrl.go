package utils

import (
	"fmt"
	"hash/fnv"
)

// Create a new FNV-1a hash object
var hasher = fnv.New64a()

func ShortenURL(url string, length int) string {
	// Reset the hash function to clear any previous state
	hasher.Reset()
	// Write the string to the hash object
	hasher.Write([]byte(url))
	// Get the hash as a byte slice
	hashBytes := hasher.Sum(nil)

	// Convert the hash to a hex string
	hashString := fmt.Sprintf("%x", hashBytes)
	// Take the first `length` characters of the hash string
	if length > len(hashString) {
		length = len(hashString)
	}
	return hashString[:length]
}
