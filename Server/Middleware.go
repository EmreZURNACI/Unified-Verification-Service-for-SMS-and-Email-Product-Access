package server

import (
	"fmt"
	"net/http"

	helpers "Helpers.go"
)

func IsSetTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("Token")
		if err != nil || cookie == nil {
			fmt.Fprintf(w, "%+v", helpers.Response(false, 401, "Unauthorized", nil))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
