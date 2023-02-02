package validatejwt

import (
	"fmt"
	"net/http"
	"os"
)

type validateJWTMiddleware struct{}

type Middleware interface {
	Wrap(h http.Handler) http.Handler
}

func NewValidateJWTMddleware() Middleware {
	return validateJWTMiddleware{}
}

func (a validateJWTMiddleware) Wrap(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		req, err := http.NewRequest(http.MethodGet, os.Getenv("VALIDATE_TOKEN_URL"), nil)
		if err != nil {
			rw.WriteHeader(401)
			fmt.Fprint(rw)
			return
		}

		req.Header.Set("Authorization", r.Header.Get("Authorization"))

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			rw.WriteHeader(401)
			fmt.Fprint(rw)
			return
		}

		if res.StatusCode == 204 {
			h.ServeHTTP(rw, r)
		} else {
			rw.WriteHeader(401)
			fmt.Fprint(rw)
			return
		}
	})
}
