package order

import (
	"domain/customer"
	customerMemory "domain/customer/memory"
	"domain/product"
	productMemory "domain/product/memory"
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
	customerRepository := customerMemory.New()
	return attachCustomerRepository(customerRepository)
}

func attachCustomerRepository(customerRepository customer.CustomerRepository) OrderConfiguration {
	return func(orderService *OrderService) error {
		orderService.customer = customerRepository
		return nil
	}
}

func loadProductMemoryRepository() OrderConfiguration {
	productRepository := productMemory.New()
	return attachProductRepository(productRepository)
}

func attachProductRepository(productRepository product.ProductRepository) OrderConfiguration {
	return func(orderService *OrderService) error {
		orderService.products = productRepository
		return nil
	}
}
