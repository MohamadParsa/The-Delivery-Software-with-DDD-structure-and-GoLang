package order

import (
	"domain/customer"
	"domain/product"
	"fmt"

	"github.com/google/uuid"
)

type OrderService struct {
	customer customer.CustomerRepository
	products product.ProductRepository
}

func (orderService *OrderService) AddOrder(customerID uuid.UUID, productsID uuid.UUID, count int) error {
	orderService, err := initialOrderService()
	if err != nil {
		return err
	}
	customer, err := orderService.customer.Get(customerID)
	if err != nil {
		return err
	}
	fmt.Println(customer)
	// selectedProducts := []product.Product{}
	// for _, productID := range productsIDs {
	// 	orderedProduct, err := orderService.products.Get(productID)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	selectedProducts = append(selectedProducts, *orderedProduct)

	// }
	return nil
}

func initialOrderService() (*OrderService, error) {
	customersLoader := loadCustomerMemoryRepository()
	productsLoader := loadProductMemoryRepository()
	return NewOrderService(customersLoader, productsLoader)
}
