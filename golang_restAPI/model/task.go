package model

import (
	"log"
    "github.com/takumi616/golang_restAPI/db"
)

type Task struct {
	Id int         `json:"id"`	
	Task string    `json:"task"`
	Status string  `json:"status"`
}

func GetAllTasks() ([]Task) {
	Db := db.Init()
    rows, err := Db.Query("SELECT * FROM task")
	
	if err != nil {
		log.Fatalf("failed to get records: %v", err)
	}

	var taskList []Task 
	for rows.Next() {
		task := Task{}
		err = rows.Scan(&task.Id, &task.Task, &task.Status)

		if err != nil {
			log.Fatalf("failed to copy a record into task: %v", err)
		}

		taskList = append(taskList, task)
	}

	return taskList
}

func GetSingleTask(id int) Task {
	Db := db.Init()
	row := Db.QueryRow("SELECT * FROM task WHERE id = $1", id)
	task := Task{}
	err := row.Scan(&task.Id, &task.Task, &task.Status)

	if err != nil {
		log.Fatalf("failed to copy a record into task: %v", err)
	}

	return task
}

func CreateTask(task Task) int {
	Db := db.Init()
	var createdTaskId int
	err := Db.QueryRow("INSERT INTO task(task, status) VALUES($1, $2) RETURNING id", task.Task, task.Status).Scan(&createdTaskId)

	if err != nil {
		log.Fatalf("failed to copy returnedId into createdTaskId: %v", err)
	}

	return createdTaskId
}

func UpdateTask(id int, task Task) int {
	Db := db.Init()
	var updatedTaskId int
	err := Db.QueryRow("UPDATE task SET task = $2, status = $3 WHERE id = $1 RETURNING id", id, task.Task, task.Status).Scan(&updatedTaskId)

	if err != nil {
		log.Fatalf("failed to copy returnedId into updatedTaskId: %v", err)
	}

	return updatedTaskId
}

func DeleteTask(id int) int {
	Db := db.Init()
	var deletedTaskId int
	err := Db.QueryRow("DELETE FROM task WHERE id = $1 RETURNING id", id).Scan(&deletedTaskId)

	if err != nil {
		log.Fatalf("failed to copy returnedId into deletedTaskId: %v", err)
	}

	return deletedTaskId
}
