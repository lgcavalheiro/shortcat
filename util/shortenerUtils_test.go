package util

import (
	"strings"
	"testing"
)

func TestGenerateShortUrl(t *testing.T) {
	size := 8
	res, _ := GenerateShortUrl(size)

	if len(res) != size {
		t.Errorf("FAILED, expected %d - got %d", size, len(res))
	}
}

func TestGenerateShortUrlNoSize(t *testing.T) {
	size := 0
	expectedSize := 6
	res, _ := GenerateShortUrl(size)

	if len(res) != expectedSize {
		t.Errorf("FAILED, expected %d - got %d", expectedSize, len(res))
	}
}

func TestGenerateShortUrlLargeSize(t *testing.T) {
	size := 32
	expectedSize := 6
	res, _ := GenerateShortUrl(size)

	if len(res) != expectedSize {
		t.Errorf("FAILED, expected %d - got %d", expectedSize, len(res))
	}
}

func TestUrlify(t *testing.T) {
	url := "https://searx.space"
	res := Urlify(url)

	if strings.Compare(url, res) != 0 {
		t.Errorf("FAILED, expected %s - got %s", url, res)
	}
}

func TestUrlifyNoPrefix(t *testing.T) {
	res := Urlify("searx.space")

	if !strings.HasPrefix(res, "https://") {
		t.Errorf("FAILED, url is broken, received: %s", res)
	}
}
