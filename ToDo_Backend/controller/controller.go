package controller

import (
	"encoding/json"
	"net/http"

	handler "github.com/shantanubose01/learn_go/handler"
)

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	data := handler.GetAllTodosHandler()
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(data); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

}
func GetTodo(w http.Response, r *http.Request) {

}

func AddTodo(w http.Response, r *http.Request) {

}
func UpdateTodo(w http.Response, r *http.Request) {

}
func DeleteTodo(w http.Response, r *http.Request) {

}
