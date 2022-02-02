package main

import (
	"github/Shimaa-Ibrahim/todo/db"
	"github/Shimaa-Ibrahim/todo/decorators"
	"github/Shimaa-Ibrahim/todo/handlers"
	"log"
	"net/http"
)

func main() {
	db.OpenDBConnection()
	defer db.Close()
	handleRequests()
}

func handleRequests() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", decorators.IsAuthenticated(handlers.GetTodoHandler))
	http.HandleFunc("/addTask", decorators.IsAuthenticated(handlers.AddTask))
	http.HandleFunc("/completeTask/", decorators.IsAuthenticated(handlers.IsTaskCompleted))
	http.HandleFunc("/deleteTask/", decorators.IsAuthenticated(handlers.DeleteTask))
	http.HandleFunc("/auth", handlers.GetAuthHandler)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/signup", handlers.SignUp)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
