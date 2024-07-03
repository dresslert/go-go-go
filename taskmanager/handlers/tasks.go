package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "taskmanager/db"
    "taskmanager/models"
    "github.com/google/uuid"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
    rows, err := db.DB.Query("SELECT id, name, status FROM tasks")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var tasks []models.Task
    for rows.Next() {
        var task models.Task
        if err := rows.Scan(&task.ID, &task.Name, &task.Status); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        tasks = append(tasks, task)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
    var task models.Task
    json.NewDecoder(r.Body).Decode(&task)
    task.ID = uuid.New().String()

    _, err := db.DB.Exec("INSERT INTO tasks (id, name, status) VALUES ($1, $2, $3)", task.ID, task.Name, task.Status)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var task models.Task
    json.NewDecoder(r.Body).Decode(&task)
    task.ID = params["id"]

    _, err := db.DB.Exec("UPDATE tasks SET name=$1, status=$2 WHERE id=$3", task.Name, task.Status, task.ID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    _, err := db.DB.Exec("DELETE FROM tasks WHERE id=$1", params["id"])
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
