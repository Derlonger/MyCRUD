package routes

import (
	"github.com/gorilla/mux"
	"modulepath/pkg/controllers"
)

var RegisterUserRoutes = func(router *mux.Router) {
	router.HandleFunc("/users", controllers.GetUser).Methods("GET")
	router.HandleFunc("/users", controllers.AddUser).Methods("POST")
	router.HandleFunc("/users/{userId}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/users/{userId}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{userId}", controllers.DeleteUser).Methods("DELETE")
}
