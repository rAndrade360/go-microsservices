package health

import "net/http"

type healthHandler struct{}

func NewHealthHandler() http.Handler {
	return healthHandler{}
}

func (healthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
