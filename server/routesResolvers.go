package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/lgcavalheiro/shortcat/model"
	"github.com/spf13/viper"
)

type redirection struct {
	Error    string
	Redirect string
	Status   int
}

type resolveAdminGetResponse struct {
	Content []model.Shortcat
	Message string
	Status  int
}

type resolveLoginResponse struct {
	Content interface{}
	Message string
	Status  int
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

	host := viper.GetString("APP_HOST")
	port := viper.GetString("APP_PORT")
	jsonData := fmt.Sprintf("{ \"user\": \"%s\", \"pwd\": \"%s\" }", user, pwd)

	resp, err := http.Post(fmt.Sprintf("%s:%s/api/auth", host, port), "application/json", bytes.NewBuffer([]byte(jsonData)))
	if err != nil {
		log.Println(err.Error())
		return redirection{
			Error:  fmt.Sprintf("Internal server error"),
			Status: http.StatusInternalServerError,
		}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return redirection{
			Error:  fmt.Sprintf("Internal server error"),
			Status: http.StatusInternalServerError,
		}
	}

	var parsedResp resolveLoginResponse
	err = json.Unmarshal([]byte(body), &parsedResp)
	if err != nil {
		log.Println(err.Error())
		return redirection{
			Error:  fmt.Sprintf("Internal server error: %s", err.Error()),
			Status: http.StatusInternalServerError,
		}
	}
	if parsedResp.Message == "Failed" {
		return redirection{
			Error:  fmt.Sprintf("Error: %s", parsedResp.Content),
			Status: http.StatusBadRequest,
		}
	}

	return redirection{
		Status:   http.StatusTemporaryRedirect,
		Redirect: "/admin",
	}
}

func resolveAdminGet(token string) map[string]interface{} {
	data := make(map[string]interface{})
	host := viper.GetString("APP_HOST")
	port := viper.GetString("APP_PORT")

	resp, err := http.Get(fmt.Sprintf("%s:%s/api?t=%s", host, port, token))
	if err != nil {
		log.Println(err.Error())
		return data
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return data
	}

	var parsedResp resolveAdminGetResponse
	err = json.Unmarshal([]byte(body), &parsedResp)
	if err != nil {
		log.Println(err.Error())
		return data
	}

	data["Shortcats"] = parsedResp.Content
	return data
}
