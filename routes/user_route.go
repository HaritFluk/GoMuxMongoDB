package routes

import (
	"GoMuxMongoDB/controllers"

	"github.com/gorilla/mux"
)

func UserRoute(router *mux.Router) {
	// All routes user comes here
	router.HandleFunc("/user", controllers.CreateUser()).Methods("POST") // Create New User

	router.HandleFunc("/user/{userId}", controllers.GetAUser()).Methods("GET") // Get A User

	router.HandleFunc("/user/{userId}", controllers.EditAUser()).Methods("PUT") // Edit A User

	router.HandleFunc("/user/{userId}", controllers.DeleteAUser()).Methods("DELETE") // Delete A User

	router.HandleFunc("/users", controllers.GetAllUser()).Methods("GET") // Get All Users

}