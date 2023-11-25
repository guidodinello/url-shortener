package utils

import "testing"

func TestShortenURL(t *testing.T) {
	var url string = "https://example.com"
	tests := []struct {
		name   string
		url    string
		length int
	}{
		{"Shorten 6 characters", url, 6},
		{"Shorten 10 characters", url, 10},
		{"Shorten 16 characters (full length)", url, 16},
		{"Shorten longer than hash", url, 16},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shortened := ShortenURL(tt.url, tt.length)
			if len(shortened) != tt.length {
				t.Errorf("shortened url length is %d, want %d", len(shortened), tt.length)
			}
		})
	}
}
