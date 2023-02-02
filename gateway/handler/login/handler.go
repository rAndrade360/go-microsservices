package auth

import (
	"fmt"
	"net/http"
	"os"
)

type loginHandler struct{}

func NewLoginHandler() http.Handler {
	return loginHandler{}
}

func (a loginHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		rw.WriteHeader(405)
		rw.Write([]byte("Can not " + r.Method + " /login"))
		return
	}

	res, err := http.Post(os.Getenv("LOGIN_URL"), "application/json", r.Body)
	if err != nil {
		fmt.Fprint(rw)
	}

	fmt.Fprint(rw, res.Body)
	rw.WriteHeader(res.StatusCode)
}
