package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/spf13/viper"
)

func Launch() {
	router := httprouter.New()

	ApiSetup(router)
	RoutesSetup(router, viper.GetBool("USE_DEFAULT_FRONTEND"))

	host := viper.GetString("APP_HOST")
	port := viper.GetString("APP_PORT")
	uri := fmt.Sprintf("%s:%s", host, port)

	log.Printf("[MAIN] - Server launched: %s\n", uri)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
