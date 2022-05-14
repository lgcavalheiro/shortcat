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
	content, message, status := resolveRootGet()
	sendJsonResponse(w, content, message, status)
}

func handleApiRootPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	content, message, status := resolveRootPost(r.Body)
	sendJsonResponse(w, content, message, status)
}

func sendJsonResponse(w http.ResponseWriter, content interface{}, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	resp := make(map[string]interface{})
	resp["message"] = message
	resp["content"] = content
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}
