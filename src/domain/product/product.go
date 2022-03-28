// Package product holds aggregates that combines many entities into a full object.
package product

import (
	entity "entity"
	"errors"

	"github.com/google/uuid"
)
//TODO: implement price actions.
//TODO: aggregate the quantity and unit. a product can have a different unit and quantity like 1 Box, 12 Pcs,...

//Product contains all data about a product.
type Product struct {
	commodity *entity.Commodity
	price     []*entity.Price
	quantity  float32
}

var (

	//ErrInvalidFirstName is returned when a commodity name is empty.
	ErrInvalidCommodityName        = errors.New("commodity name is required")
	//ErrInvalidCommodityModel is returned when a commodity model is empty.
	ErrInvalidCommodityModel       = errors.New("commodity model is required")
	//ErrInvalidCommodityCompanyName is returned when a commodity company name is empty.
	ErrInvalidCommodityCompanyName = errors.New("commodity company name is required")
	//ErrOverwriteCommodityID is returned when a valid id wants changes.
	ErrOverwriteCommodityID        = errors.New("commodity identifier cannot overwrite")
	//ErrIDIsNotValid is returned when a id is empty.
	ErrIDIsNotValid                = errors.New("commodity identifier is not valid")
	//ErrCountIsNotValid is returned when a count is less than zero.
	ErrCountIsNotValid             = errors.New("count is not valid")
)

//NewProduct is a constructs function to make a new product object. 
func NewProduct(name, model, companyName, description string) (*Product, error) {
	product := new(Product)
	newCommodity := &entity.Commodity{
		Name:        name,
		Model:       model,
		CompanyName: companyName,
		Description: description,
	}

	if err := checkCommodityRequirdField(newCommodity); err != nil {
		return product, err
	}

	product = &Product{
		commodity: newCommodity,
		price:     make([]*entity.Price, 0),
	}

	return product, nil
}

//GetID returns the product ID.
func (product *Product) GetID() uuid.UUID {
	return product.commodity.ID
}

//SetID sets the product ID.
func (product *Product) SetID(id uuid.UUID) error {
	product.commodityExists()
	if newIdIsNotValid(id) {
		return ErrIDIsNotValid
	}
	if commodityIdIsNotEmpty(product.commodity.ID) {
		return ErrOverwriteCommodityID
	}
	product.commodity.ID = id
	return nil
}

//GetModel returns the product model.
func (product *Product) GetModel() string {
	return product.commodity.Model
}

//SetModel sets the product model.
func (product *Product) SetModel(model string) error {
	product.commodityExists()
	if stringValueIsEmpty(model) {
		return ErrInvalidCommodityModel
	}

	product.commodity.Model = model
	return nil
}

//GetCompanyName returns the product company name.
func (product *Product) GetCompanyName() string {
	return product.commodity.CompanyName
}

//SetCompanyName sets the product company name.
func (product *Product) SetCompanyName(companyName string) error {
	product.commodityExists()
	if stringValueIsEmpty(companyName) {
		return ErrInvalidCommodityCompanyName
	}

	product.commodity.CompanyName = companyName
	return nil
}

//GetDescription returns the product description.
func (product *Product) GetDescription() string {
	return product.commodity.Description
}

//SetDescription sets the product description.
func (product *Product) SetDescription(description string) error {
	product.commodityExists()
	product.commodity.CompanyName = description
	return nil
}

//GetQuantity returns the product quantity.
func (product *Product) GetQuantity() float32 {
	return product.quantity
}

//SetQuantity sets the product quantity.
func (product *Product) SetQuantity(count float32) error {
	product.commodityExists()
	if countIsNotValid(count) {
		return ErrCountIsNotValid
	}

	product.quantity = count
	return nil
}
//AddQuantity adds to the product quantity.
func (product *Product) AddQuantity(count float32) error {
	product.commodityExists()
	if countIsNotValid(count) {
		return ErrCountIsNotValid
	}

	product.quantity += count
	return nil
}
//ReduceQuantity reduces the product quantity.
func (product *Product) ReduceQuantity(count float32) error {
	product.commodityExists()
	if countIsNotValid(count) || countIsNotValid(product.quantity-count) {
		return ErrCountIsNotValid
	}

	product.quantity -= count
	return nil
}
func countIsNotValid(count float32) bool {
	return count < 0
}
func checkCommodityRequirdField(commodity *entity.Commodity) error {
	if stringValueIsEmpty(commodity.Name) {
		return ErrInvalidCommodityName
	}
	if stringValueIsEmpty(commodity.Model) {
		return ErrInvalidCommodityModel
	}
	if stringValueIsEmpty(commodity.CompanyName) {
		return ErrInvalidCommodityCompanyName
	}
	return nil
}

func stringValueIsEmpty(value string) bool {
	return value == ""
}

func (product *Product) commodityExists() {
	if product.commodity == nil {
		product.commodity = new(entity.Commodity)
	}
}

func commodityIdIsNotEmpty(id uuid.UUID) bool {
	return id != uuid.Nil
}

func newIdIsNotValid(id uuid.UUID) bool {
	return id == uuid.Nil
}
