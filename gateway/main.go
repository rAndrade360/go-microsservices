package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	healthhandler "github.com/rAndrade360/go-microsservices/gateway/handler/health"
	loginhandler "github.com/rAndrade360/go-microsservices/gateway/handler/login"
	validatejwt "github.com/rAndrade360/go-microsservices/gateway/middleware/validateJWT"
)

func init() {
	godotenv.Load()
}

func main() {
	loginHandler := loginhandler.NewLoginHandler()
	healthHandler := healthhandler.NewHealthHandler()

	sm := http.NewServeMux()
	pm := http.NewServeMux()

	sm.Handle("/login", loginHandler)
	pm.Handle("/", validatejwt.NewValidateJWTMddleware().Wrap(healthHandler))
	sm.Handle("/health", healthHandler)

	log.Fatal(http.ListenAndServe(":8080", sm))
}
