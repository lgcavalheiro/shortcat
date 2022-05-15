package server

import (
	"fmt"
	"log"
	"net/http"
	"strings"

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

func resolveLogin(user, pwd string) redirection {
	if len(user) == 0 || len(pwd) == 0 {
		return redirection{
			Error:  fmt.Sprintf("Error: %s", "Please provide valid credentials and try again."),
			Status: http.StatusBadRequest,
		}
	}

	if strings.Compare(user, admUser) != 0 || strings.Compare(pwd, admPwd) != 0 {
		return redirection{
			Error:  fmt.Sprintf("Error: %s", "Invalid credentials."),
			Status: http.StatusUnauthorized,
		}
	}

	return redirection{
		Status:   http.StatusTemporaryRedirect,
		Redirect: "/admin",
	}
}

func resolveAdminGet() map[string]interface{} {
	data := make(map[string]interface{})
	shortcats, err := model.GetAllShortCats()
	if err != nil {
		log.Println(err.Error())
		return data
	}
	data["Shortcats"] = shortcats
	return data
}
