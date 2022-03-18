package main

import (
	"GoMuxMongoDB/configs"
	"GoMuxMongoDB/routes"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// run database
	configs.ConnectDB()

	// run routes
	routes.UserRoute(router)

	log.Fatal(http.ListenAndServe(":6000", router))
}