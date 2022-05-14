package util

import (
	"strings"

	gonanoid "github.com/matoous/go-nanoid"
)

var DEFAULT_SIZE = 6

func GenerateShortUrl(size int) (string, error) {
	if size <= 0 || size > 16 {
		size = DEFAULT_SIZE
	}
	id, err := gonanoid.Nanoid(size)
	return id, err
}

func Urlify(url string) string {
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		return url
	}
	return "https://" + url
}
