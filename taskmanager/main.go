package main

import (
    "net/http"
    "taskmanager/db"
    "taskmanager/handlers"
    "github.com/gorilla/mux"
)

func main() {
    db.Init()

    r := mux.NewRouter()
    r.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
    r.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
    r.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")
    r.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")

    http.ListenAndServe(":8080", r)
}
