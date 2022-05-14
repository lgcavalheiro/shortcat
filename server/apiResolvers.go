package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/lgcavalheiro/shortcat/model"
	"github.com/lgcavalheiro/shortcat/util"
)

type postBody struct {
	Url           string `json:"url"`
	IdSize        int    `json:"id_size"`
	UserDefinedId string `json:"user_defined_id"`
}

func resolveRootGet() (interface{}, string, int) {
	shortcats, err := model.GetAllShortCats()
	if err != nil {
		log.Println(err.Error())
		return fmt.Sprintf("Bad request: %s", err.Error()), "Failed", http.StatusBadRequest
	}

	return shortcats, "Success", http.StatusOK
}

func resolveRootPost(body io.ReadCloser) (interface{}, string, int) {
	var postBody postBody
	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&postBody)
	if err != nil {
		log.Println(err.Error())
		return fmt.Sprintf("Bad request: %s", err.Error()), "Failed", http.StatusBadRequest
	}

	var shortenedUrl string
	if len(postBody.UserDefinedId) > 0 {
		shortenedUrl = postBody.UserDefinedId
	} else {
		shortenedUrl, err = util.GenerateShortUrl(postBody.IdSize)
		if err != nil {
			log.Println(err.Error())
			return fmt.Sprintf("Server side error: %s", err.Error()), "Failed", http.StatusInternalServerError
		}
	}

	shortcat := model.Shortcat{
		Url:      util.Urlify(postBody.Url),
		ShortUrl: shortenedUrl,
	}

	err = model.CreateShortCat(shortcat)
	if err != nil {
		log.Println(err.Error())
		return fmt.Sprintf("Bad request: %s", err.Error()), "Failed", http.StatusBadRequest
	}

	return shortcat, "Success", http.StatusOK
}
