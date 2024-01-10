package routes

import (
	"github.com/gorilla/mux"
	controller "github.com/shantanubose01/learn_go/controller"
)

func Route() *mux.Router {
	R := mux.NewRouter()
	R.HandleFunc("/", controller.GetAllTodos).Methods("GET")

	// R.HandleFunc("/{id}", GetTodo).Methods("GET")
	// R.HandleFunc("/", AddTodo).Methods("POST")
	// R.HandleFunc("/{id}", UpdateTodo).Methods("PUT")
	// R.HandleFunc("/{id}", DeleteTodo).Methods("DELETE")

	return R
}
