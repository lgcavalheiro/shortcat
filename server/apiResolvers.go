package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/lgcavalheiro/shortcat/model"
	"github.com/lgcavalheiro/shortcat/util"
	"github.com/spf13/viper"
)

type newShortcat struct {
	Url           string `json:"url"`
	IdSize        int    `json:"id_size"`
	UserDefinedId string `json:"user_defined_id"`
}

type authBody struct {
	User string `json:"user"`
	Pwd  string `json:"pwd"`
}

type resolution struct {
	Content interface{}
	Message string
	Status  int
}

func validateToken(token string) bool {
	admSecret := viper.GetString("ADM_SECRET")
	if len(token) == 0 || strings.Compare(token, admSecret) != 0 {
		return false
	}
	return true
}

func resolveRootGet(token string) resolution {
	if !validateToken(token) {
		return resolution{
			Content: "Invalid token",
			Message: "Failed",
			Status:  http.StatusUnauthorized,
		}
	}

	shortcats, err := model.GetAllShortCats()
	if err != nil {
		log.Println(err.Error())
		return resolution{
			Content: fmt.Sprintf("Bad request: %s", err.Error()),
			Message: "Failed",
			Status:  http.StatusBadRequest,
		}
	}

	return resolution{
		Content: shortcats,
		Message: "Success",
		Status:  http.StatusOK,
	}
}

func resolveRootPost(body io.ReadCloser) resolution {
	var newShortcat newShortcat
	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&newShortcat)
	if err != nil {
		log.Println(err.Error())
		return resolution{
			Content: fmt.Sprintf("Bad request: %s", err.Error()),
			Message: "Failed",
			Status:  http.StatusBadRequest,
		}
	}

	var shortenedUrl string
	if len(newShortcat.UserDefinedId) > 0 {
		shortenedUrl = newShortcat.UserDefinedId
	} else {
		shortenedUrl, err = util.GenerateShortUrl(newShortcat.IdSize)
		if err != nil {
			log.Println(err.Error())
			return resolution{
				Content: fmt.Sprintf("Server side error: %s", err.Error()),
				Message: "Failed",
				Status:  http.StatusInternalServerError,
			}
		}
	}

	shortcat := model.Shortcat{
		Url:      util.Urlify(newShortcat.Url),
		ShortUrl: shortenedUrl,
	}

	err = model.CreateShortCat(shortcat)
	if err != nil {
		log.Println(err.Error())
		return resolution{
			Content: fmt.Sprintf("Bad request: %s", err.Error()),
			Message: "Failed",
			Status:  http.StatusBadRequest,
		}
	}

	return resolution{
		Content: shortcat,
		Message: "Success",
		Status:  http.StatusOK,
	}
}

func resolveAuthPost(body io.ReadCloser) resolution {
	var authBody authBody
	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&authBody)
	if err != nil {
		log.Println(err.Error())
		return resolution{
			Content: fmt.Sprintf("Bad request: %s", err.Error()),
			Message: "Failed",
			Status:  http.StatusBadRequest,
		}
	}

	admUser := viper.GetString("ADM_USER")
	admPwd := viper.GetString("ADM_PWD")
	admSecret := viper.GetString("ADM_SECRET")

	if strings.Compare(authBody.User, admUser) != 0 || strings.Compare(authBody.Pwd, admPwd) != 0 {
		return resolution{
			Content: "Authentication failed",
			Message: "Failed",
			Status:  http.StatusUnauthorized,
		}
	}

	content := make(map[string]string)
	content["token"] = admSecret
	return resolution{
		Content: content,
		Message: "Success",
		Status:  http.StatusOK,
	}
}
