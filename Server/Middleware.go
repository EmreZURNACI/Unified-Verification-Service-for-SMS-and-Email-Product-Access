package server

import (
	"fmt"
	"net/http"

	c "Connection.go"
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

func IsDatabaseConnected(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := c.Connection().Ping()
		if err != nil {
			fmt.Fprintf(w, "%+v", helpers.Response(false, 400, "Database Connection Was Crashed", nil))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
