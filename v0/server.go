package v0

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) initialize() {
	fmt.Println("Wellcome to gotoko")

	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Printf("Listening to port %s", addr)

	log.Fatal(http.ListenAndServe(addr, server.Router))

}

func Run() {
	server := Server{}

	server.initialize()
	server.Run(":9000")
}
