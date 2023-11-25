package utils

import (
	"errors"
	"strings"
)

// Extracts the part of the URL after the last '/'.
func ExtractKeyFromURL(url string) (string, error) {
	lastSlashIndex := strings.LastIndex(url, "/")
	if lastSlashIndex == -1 {
		return "", errors.New("no '/' found in URL")
	}
	return url[lastSlashIndex+1:], nil
}
