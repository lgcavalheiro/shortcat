package server

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func ApiSetup(router *httprouter.Router) {
	router.GET("/api/", handleApiRootGet)
	router.POST("/api/", handleApiRootPost)
}

func handleApiRootGet(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	resolution := resolveRootGet()
	sendJsonResponse(w, resolution)
}

func handleApiRootPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	resolution := resolveRootPost(r.Body)
	sendJsonResponse(w, resolution)
}

func sendJsonResponse(w http.ResponseWriter, res resolution) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Status)
	resp := make(map[string]interface{})
	resp["message"] = res.Message
	resp["content"] = res.Content
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}
