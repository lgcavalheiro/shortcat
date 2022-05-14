package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/lgcavalheiro/shortcat/model"
)

func RoutesSetup(router *httprouter.Router, useDefaultFront bool) {
	router.GET("/go/:shorturl", handleShortcat)
	if useDefaultFront {
		log.Println("TBD frontend")
	}
}

func handleShortcat(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	shortUrl := params.ByName("shorturl")

	shortCat, err := model.GetShortCatByShortUrl(shortUrl)
	if err != nil {
		res := []byte(fmt.Sprintf("<h1>Error: %s</h1>", err.Error()))
		w.Write(res)
		return
	}

	http.Redirect(w, r, shortCat.Url, http.StatusTemporaryRedirect)
}
