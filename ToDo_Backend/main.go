package main

import (
	"fmt"
	"log"
	"net/http"

	router "github.com/shantanubose01/learn_go/routes"
)

func main() {
	fmt.Println("Building a Todo backend server")
	r := router.Route()
	// http.HandleFunc("/", r)
	log.Fatal(http.ListenAndServe(":3001", r))
	fmt.Println("server listening to 3001..")
}
