package server

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func RoutesSetup(router *httprouter.Router, useDefaultFront bool) {
	router.GET("/go/:shorturl", handleShortcat)
	if useDefaultFront {
		log.Println("TBD frontend")
	}
}

func handleShortcat(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	shortUrl := params.ByName("shorturl")

	redirect := resolveShortCat(shortUrl)
	if len(redirect.Error) > 0 {
		res := []byte(redirect.Error)
		w.WriteHeader(redirect.Status)
		w.Write(res)
		return
	}

	http.Redirect(w, r, redirect.Redirect, redirect.Status)
}
