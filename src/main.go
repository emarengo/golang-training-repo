package main

import (
	"fmt"
	"time"
)

type Order struct {
	id           int
	creationDate time.Time
	orderLine    *OrderLine
	totalPrice   int
	createdBy    string
}

type OrderLine struct {
	id           int
	creationDate time.Time
	item         string
	qty          int
	unitPrice    int
}

func NewOrder(id int, creationDate time.Time, orderLine *OrderLine, totalPrice int, createdBy string) *Order {
	return &Order{
		creationDate: creationDate,
		id:           id,
		orderLine:    orderLine,
		totalPrice:   totalPrice,
		createdBy:    createdBy,
	}
}

func NewOrderLine(id int, creationDate time.Time, item string, qty int, unitPrice int) *OrderLine {
	return &OrderLine{
		id:           id,
		creationDate: creationDate,
		item:         item,
		qty:          qty,
		unitPrice:    unitPrice,
	}
}

type RegisterOrder map[string]*Order

func main() {

	//create order
	order1 := NewOrder(1, time.Now(), nil, 0, "Ernesto")

	//create orderLine
	orderLine := NewOrderLine(2, time.Now(), "nutrients", 1, 100)

	//append orderLine
	order1.orderLine = orderLine
	order1.totalPrice = orderLine.unitPrice

	//update the orderLine
	order1.orderLine.unitPrice = 150
	order1.totalPrice = orderLine.unitPrice

	//remove the orderLine
	order1.orderLine = nil

	fmt.Println(order1.totalPrice)

}
