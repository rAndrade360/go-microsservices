package auth

import (
	"fmt"
	"io"
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
		rw.WriteHeader(500)
		fmt.Fprint(rw, `{"message": "Internal Server Error"}`)
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		rw.WriteHeader(500)
		fmt.Fprint(rw, `{"message": "Internal Server Error"}`)
	}

	defer res.Body.Close()

	rw.WriteHeader(res.StatusCode)
	fmt.Fprint(rw, string(b))
}
