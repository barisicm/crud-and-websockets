package models

import (
	"fmt"
	u "sartura-task/backend/utils"
)

// order contains the details of an order
type Order struct {
	ID     int     `json:"id"`
	Symbol string  `json:"symbol"`
	Type   string  `json:"type"`
	Side   string  `json:"side"`
	Price  float32 `json:"price"`
	Amount float32 `json:"amount"`
}

func (order *Order) Validate() (map[string]interface{}, bool) {

	if order.Amount <= 0 {
		return u.Message(false, "Order amount should be bigger than 0"), false
	}

	if order.Price <= 0 {
		return u.Message(false, "Order price must be bigger than 0"), false
	}

	fmt.Println(order.Side)
	if order.Side != "buy" {
		if order.Side != "sell" {
			return u.Message(false, "Order side not set correctly"), false
		}
	}

	order.Type = "market"

	return u.Message(true, "success"), true
}

func (order *Order) Create() map[string]interface{} {

	if resp, ok := order.Validate(); !ok {
		return resp
	}

	GetDB().Create(order)

	resp := u.Message(true, "success")
	resp["order"] = order
	return resp
}

func (order *Order) GetOrders() []*Order {

	orders := make([]*Order, 0)
	err := GetDB().Table("orders").Find(&orders).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return orders
}

func (order *Order) Delete(orderId int) {

	err := GetDB().Where("ID = ?", orderId).Delete(Order{})
	if err != nil {
		fmt.Println(err)
	}

}
