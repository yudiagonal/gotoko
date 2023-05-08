package v0

import "gotoko/v0/controller"

func (server *Server) initializeRoutes() {
	server.Router.HandleFunc("/home", controller.Home).Methods("GET")
}
