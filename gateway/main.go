package main

import (
	"net/http"

	healthhandler "github.com/rAndrade360/go-microsservices/gateway/handler/health"
	loginhandler "github.com/rAndrade360/go-microsservices/gateway/handler/login"
	validatejwt "github.com/rAndrade360/go-microsservices/gateway/middleware/validateJWT"
)

func main() {
	loginHandler := loginhandler.NewLoginHandler()
	healthHandler := healthhandler.NewHealthHandler()

	sm := http.NewServeMux()
	pm := http.NewServeMux()

	sm.Handle("/login", loginHandler)
	pm.Handle("/", healthHandler)
	sm.Handle("/health", validatejwt.NewValidateJWTMddleware().Wrap(pm))

	http.ListenAndServe(":8080", sm)
}
