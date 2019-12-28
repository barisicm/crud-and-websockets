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

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

var CreateOrder = func(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	if (r).Method == "OPTIONS" {
		fmt.Println("JEBO TE JA")
		return
	}

	order := &models.Order{}

	err := json.NewDecoder(r.Body).Decode(order)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := order.Create()
	u.Respond(w, resp)
}

var GetOrders = func(w http.ResponseWriter, r *http.Request) {
	order := &models.Order{}

	data, err := json.Marshal(order.GetOrders())
	if err != nil {
		u.Respond(w, u.Message(false, "Error while fetching all orders from database"))
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

var DeleteOrder = func(w http.ResponseWriter, r *http.Request) {

	order := &models.Order{}
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil && id < 0 {
		u.Respond(w, u.Message(false, "Error while converting id to int"))
		u.Respond(w, u.Message(false, "Error, someone is inputing negative numbers"))
	} else {
		order.Delete(id)
	}

}
