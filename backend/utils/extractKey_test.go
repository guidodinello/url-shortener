package utils

import "testing"

func TestExtractKeyFromURL(t *testing.T) {
	var prefix string = "https://example.com/"
	tests := []struct {
		name   string
		prefix string
		suffix string
	}{
		{"numbers only", prefix, "86169"},
		{"letters only", prefix, "abcde"},
		{"letters and numbers", prefix, "A98B6cde4"},
		{"letters, numbers, and symbols", prefix, "A98B6cde4!@#$%^&*()_+-=abc"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := ExtractKeyFromURL(tt.prefix + tt.suffix)
			if err != nil {
				t.Errorf("error: %v", err)
			}
			if key != tt.suffix {
				t.Errorf("key is %s, want %s", key, tt.suffix)
			}
		})
	}
}
