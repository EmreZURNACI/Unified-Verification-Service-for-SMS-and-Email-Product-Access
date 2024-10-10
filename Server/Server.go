package Server

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var wg sync.WaitGroup

func Server() {
	_mux := mux.NewRouter()
	wg.Add(2)
	go AuthServer(_mux, &wg)
	go ProductServer(_mux, &wg)
	//http.ListenAndServeTLS(":443", "server.crt", "server.key", _mux)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})
	//to connect with react client
	handler := c.Handler(_mux)
	http.ListenAndServe(":8080", handler)
	wg.Wait()
}
