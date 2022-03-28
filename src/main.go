package main

import (
	"domain/customer"
	"log"

	"github.com/google/uuid"
)

func main() {
	customer, _ := customer.NewCustomer("Mohamad", "Parsa", "00", "W")
	id, _ := uuid.NewUUID()
	err := customer.SetID(id)
	log.Println(err)
	log.Println("customer:", customer.GetID(), uuid.Nil)
}
