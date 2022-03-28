package order

import (
	"domain/customer"
	"fmt"

	"github.com/google/uuid"
)

type OrderService struct {
	customer customer.CustomerRepository
}

func AddOrder(customerID uuid.UUID, productsID []uuid.UUID) error {
	orderService, err := initialOrderService()
	if err != nil {
		return err
	}
	customer, err := orderService.customer.Get(customerID)
	if err != nil {
		return err
	}
	fmt.Println(customer)
	return nil
}

func initialOrderService() (*OrderService, error) {
	customersLoader := loadCustomerMemoryRepository()

	return NewOrderService(customersLoader)
}
