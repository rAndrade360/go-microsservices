package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	healthhandler "github.com/rAndrade360/go-microsservices/gateway/handler/health"
	loginhandler "github.com/rAndrade360/go-microsservices/gateway/handler/login"
)

func init() {
	godotenv.Load()
}

func main() {
	loginHandler := loginhandler.NewLoginHandler()
	healthHandler := healthhandler.NewHealthHandler()

	sm := http.NewServeMux()

	sm.Handle("/health", healthHandler)
	sm.Handle("/login", loginHandler)

	log.Fatal(http.ListenAndServe(":8080", sm))
}
