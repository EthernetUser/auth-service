package routes

import (
	controllers "auth-service/controllers"

	"github.com/gorilla/mux"
)

var userController = &controllers.UserController{}

func CreateRoutes(router *mux.Router) *mux.Router {
	userPrefix := "/user/"
	userController.Init()
	router.PathPrefix(userPrefix).Path("/search/").HandlerFunc(userController.Search).Methods("POST")
	router.PathPrefix(userPrefix).Path("/bulk/").HandlerFunc(userController.Create).Methods("POST")
	router.PathPrefix(userPrefix).Path("/bulk/").HandlerFunc(userController.Update).Methods("PUT")
	router.PathPrefix(userPrefix).Path("/bulk/").HandlerFunc(userController.Delete).Methods("DELETE")
	return router
}
