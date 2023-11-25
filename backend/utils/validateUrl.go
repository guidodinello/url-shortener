package utils

import (
	"strings"
	"url-shortener/backend/models"
)

var domains = []string{"https://shorty.com/", "https://www.shorty.com/"}

func some(elements []string, condition func(string) bool) bool {
	for _, element := range elements {
		if condition(element) {
			return true
		}
	}
	return false
}

func ValidateUrl(url string) error {
	if url == "" {
		return models.ErrEmptyUrl
	}

	if !some(domains, func(domain string) bool {
		return strings.HasPrefix(url, domain)
	}) {
		return models.ErrInvalidUrl
	}

	return nil
}
