package order

import (
	"domain/customer"
	"domain/customer/customermemory"
)

type OrderConfiguration func(orderService *OrderService) error

func NewOrderService(configurators ...OrderConfiguration) (*OrderService, error) {
	orderService := &OrderService{}
	for _, configurator := range configurators {
		err := configurator(orderService)
		if err != nil {
			return nil, err
		}
	}
	return orderService, nil
}

func loadCustomerMemoryRepository() OrderConfiguration {
	customerRepository := customermemory.New()
	return attachCustomerRepository(customerRepository)
}

func attachCustomerRepository(customerRepository customer.CustomerRepository) OrderConfiguration {
	return func(orderService *OrderService) error {
		orderService.customer = customerRepository
		return nil
	}
}
