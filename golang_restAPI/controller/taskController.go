package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/takumi616/golang_restAPI/model"
)


func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	taskList := model.GetAllTasks()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Method", r.Method)
	err := json.NewEncoder(w).Encode(taskList)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetSingleTask(w http.ResponseWriter, r *http.Request) {
	paramMap := mux.Vars(r)
	param, _ := strconv.Atoi(paramMap["id"])

	task := model.GetSingleTask(param)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Method", r.Method)
	err := json.NewEncoder(w).Encode(task)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	task := model.Task{}
	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := model.CreateTask(task)
	w.Header().Set("Method", r.Method)
	w.WriteHeader(http.StatusCreated)
	responseMessage := fmt.Sprintf("Created record successfully and its id is %d\n", id)
	_, err = w.Write([]byte(responseMessage))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	task := model.Task{}
	paramMap := mux.Vars(r)
	param, _ := strconv.Atoi(paramMap["id"])
	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := model.UpdateTask(param, task)

	w.Header().Set("Method", r.Method)
	responseMessage := fmt.Sprintf("Updated record successfully and its id is %d\n", id)
	_, err = w.Write([]byte(responseMessage))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	paramMap := mux.Vars(r)
	param, _ := strconv.Atoi(paramMap["id"])
	id := model.DeleteTask(param)

	w.Header().Set("Method", r.Method)
	resultMessage := fmt.Sprintf("Deleted record successfully and its id is %d\n", id)
	w.Write([]byte(resultMessage))
}