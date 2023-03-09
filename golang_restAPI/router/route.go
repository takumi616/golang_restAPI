package router

import (
	
	"github.com/takumi616/golang_restAPI/controller"
	"github.com/gorilla/mux"
)

func CreateRoutes() *mux.Router {
	
	r := mux.NewRouter()

	r.HandleFunc("/tasks", controller.GetAllTasks).Methods("GET")
	r.HandleFunc("/tasks/{id}", controller.GetSingleTask).Methods("GET")
	r.HandleFunc("/tasks", controller.CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", controller.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", controller.DeleteTask).Methods("DELETE")
	
	return r
}

