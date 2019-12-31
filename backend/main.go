package main

import (
	"fmt"
	"net/http"
	"os"

	"crud-and-websockets/backend/controllers"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	corsObj := handlers.AllowedOrigins([]string{"*"})
	corsObj = handlers.AllowedHeaders([]string{"Content-Type"})
	corsObj = handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "OPTIONS"})

	router.HandleFunc("/api/orders", controllers.CreateOrder).Methods("POST")
	// router.HandleFunc("/api/order", controllers.CreateOrder).Methods("POST")
	router.HandleFunc("/api/list", controllers.GetOrders).Methods("GET")
	router.HandleFunc("/api/delete/{id}", controllers.DeleteOrder).Methods("DELETE")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, handlers.CORS(corsObj)(router))
	if err != nil {
		fmt.Print(err)
	}
}
