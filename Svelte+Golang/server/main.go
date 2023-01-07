package main

import (
	"fmt"
	"log"
	"net/http"
	controller "server/controller"
	db "server/db"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", controller.Home).Methods("GET","OPTIONS")
	router.HandleFunc("/update/{id}", controller.UpdateTodo).Methods("PUT","OPTIONS")
	router.HandleFunc("/delete/{id}", controller.DeleteTodo).Methods("DELETE","OPTIONS")
	router.HandleFunc("/create", controller.CreateTodo).Methods("POST","OPTIONS")

	return router
}

func main() {
	fmt.Println("Welcome to backend of Todo App")
	router := Router()
	db.Connect()
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
