package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sartura-task/backend/models"
	u "sartura-task/backend/utils"
	"strconv"

	"github.com/gorilla/mux"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

var CheckRequest = func(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

var CreateOrder = func(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	order := &models.Order{}
	fmt.Println(r.Body)
	fmt.Println(r.FormValue("amount"))
	fmt.Println(r.FormValue("price"))
	fmt.Println(r.FormValue("side"))
	fmt.Println(r.FormValue("type"))
	fmt.Println(r.FormValue("symbol"))
	fmt.Println(r.Body)

	if s, err := strconv.ParseFloat(r.FormValue("amount"), 64); err == nil {
		order.Amount = float32(s)
	}
	if s, err := strconv.ParseFloat(r.FormValue("price"), 64); err == nil {
		order.Price = float32(s)
	}
	order.Side = r.FormValue("side")
	order.Type = r.FormValue("type")
	order.Symbol = r.FormValue("symbol")

	resp := order.Create()
	u.Respond(w, resp)
}

var GetOrders = func(w http.ResponseWriter, r *http.Request) {
	order := &models.Order{}
	fmt.Println("GetOrdersPozvan")
	data, err := json.Marshal(order.GetOrders())
	if err != nil {
		u.Respond(w, u.Message(false, "Error while fetching all orders from database"))
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

var DeleteOrder = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteOrderPozvan")
	fmt.Println(r.Method)
	order := &models.Order{}
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil && id < 0 {
		u.Respond(w, u.Message(false, "Error while converting id to int"))
		u.Respond(w, u.Message(false, "Error, someone is inputing negative numbers"))
	} else {
		order.Delete(id)
	}

}
