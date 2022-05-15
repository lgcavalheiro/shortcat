package server

import (
	"fmt"
	"net/http"

	"github.com/lgcavalheiro/shortcat/model"
)

type redirection struct {
	Error    string
	Redirect string
	Status   int
}

func resolveShortCat(shortUrl string) redirection {
	if len(shortUrl) == 0 {
		return redirection{
			Error:  fmt.Sprintf("<h1>Error: %s</h1>", "Invalid short url"),
			Status: http.StatusNotFound,
		}
	}

	shortCat, err := model.GetShortCatByShortUrl(shortUrl)
	if err != nil {
		return redirection{
			Error:  fmt.Sprintf("<h1>Error: %s</h1>", err.Error()),
			Status: http.StatusInternalServerError,
		}
	}

	shortCat.Clicks++
	model.UpdateShortCat(shortCat)
	return redirection{
		Redirect: shortCat.Url,
		Status:   http.StatusTemporaryRedirect,
	}
}
