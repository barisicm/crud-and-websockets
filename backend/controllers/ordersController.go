package controllers

import (
	"encoding/json"
	"net/http"
	"sartura-task/backend/models"
	u "sartura-task/backend/utils"
	"strconv"

	"github.com/gorilla/mux"
)

var CreateOrder = func(w http.ResponseWriter, r *http.Request) {
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
