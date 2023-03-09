package server

import (
	"net/http"
	"github.com/takumi616/golang_restAPI/router"
)

func CreateServer(port string) *http.Server {

	r := router.CreateRoutes()

	server := &http.Server {
		Addr: port,
		Handler: r,
	}

	return server
}