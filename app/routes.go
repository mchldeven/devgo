package app

import (
	"github.com/gorilla/mux"
	"github.com/michaeldeven/devgo/app/controllers"
)

func (server *Server) initializeRoutes() {
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
}