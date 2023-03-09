package main

import (
	"github.com/takumi616/golang_restAPI/server"
)

func main() {	
	s := server.CreateServer(":8080")
	
	s.ListenAndServe()
}