package main

import (
	"domain/customer"
	"domain/product"

	"fmt"
	"services/order"

	"github.com/google/uuid"
)

func main() {

	newCudtomer := addNewCustomer()
	customerID := newCudtomer.GetID()
	newProduct := addNewProduct()
	productIDs := []uuid.UUID{newProduct.GetID()}
	orders := order.OrderService{}
	err := orders.AddOrder(customerID, productIDs)
	fmt.Println(orders, err)
}
func addNewCustomer() *customer.Customer {
	validCustomer, _ := customer.NewCustomer("Mohamad", "Parsa", "00989122212221", "Parsa@gmail.com")
	id, _ := uuid.NewUUID()
	validCustomer.SetID(id)
	return validCustomer
}
func addNewProduct() *product.Product {
	product, _ := product.NewProduct("Apple Watch", "SE 2021", "Apple", "")
	id, _ := uuid.NewUUID()
	product.SetID(id)
	return product
}
