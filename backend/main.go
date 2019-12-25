package main

import (
	"fmt"
	"net/http"
	"os"
	"sartura-task/backend/controllers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/order", controllers.CreateOrder).Methods("POST")
	router.HandleFunc("/api/list", controllers.GetOrders).Methods("GET")
	router.HandleFunc("/api/delete/{id}", controllers.DeleteOrder).Methods("DELETE")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
