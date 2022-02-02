package handlers

import (
	"fmt"
	"github/Shimaa-Ibrahim/todo/db"
	"github/Shimaa-Ibrahim/todo/decorators"
	"github/Shimaa-Ibrahim/todo/models"
	"html/template"
	"log"
	"net/http"
	"strings"
)

// var UserID uint

type Todo struct {
	Error string
	Tasks []models.Task
}

var todo Todo

func GetTodoHandler(w http.ResponseWriter, r *http.Request) {
	db.DB.Where("user_id = ?", decorators.UserID).Find(&todo.Tasks)
	template, _ := template.ParseFiles("templates/todo.html")
	err := template.Execute(w, todo)
	if err != nil {
		log.Println("Error executing template :", err)
	}
}

func AddTask(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	todo.Error = ""
	newtask := r.FormValue("task")
	if newtask == "" {
		todo.Error = "You must add task!"
		http.Redirect(w, r, "/", http.StatusFound)
		return

	}
	task := &models.Task{Text: newtask, UserID: decorators.UserID}
	result := db.DB.Create(&task)
	if result.Error != nil {
		todo.Error = "Error Occurred, Please Try Again"
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func IsTaskCompleted(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	task_id := strings.Split(r.URL.Path, "/")[2]
	db.DB.First(&task, task_id)
	completed := !task.IsCompleted
	task.IsCompleted = completed
	result := db.DB.Save(task)
	if result.Error != nil {
		todo.Error = "Error Occurred, Please Try Again"
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	task_id := strings.Split(r.URL.Path, "/")[2]
	fmt.Println(task_id)
	result := db.DB.Delete(&models.Task{}, task_id)
	if result.Error != nil {
		todo.Error = "Error Occurred, Please Try Again"
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
