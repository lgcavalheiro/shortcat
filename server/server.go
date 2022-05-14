package server

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Launch() {
	router := httprouter.New()

	ApiSetup(router)
	RoutesSetup(router, false)

	log.Println("[MAIN] - Server launched: http://localhost:8034")
	log.Fatal(http.ListenAndServe(":8034", router))
}
