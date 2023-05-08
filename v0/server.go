package v0

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func (server *Server) initialize(appConfig AppConfig) {
	fmt.Println("Wellcome to" + appConfig.AppName)

	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Printf("Listening to port %s", addr)

	log.Fatal(http.ListenAndServe(addr, server.Router))

}

func Run() {
	server := Server{}
	appConfig := AppConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error .env file")
	}

	appConfig.AppName = getEnv("APP_NAME", "GotokoWeb")
	appConfig.AppEnv = getEnv("APP_ENV", "development")
	appConfig.AppPort = getEnv("APP_PORT", "9000")

	server.initialize(appConfig)
	server.Run(":" + appConfig.AppPort)
}
