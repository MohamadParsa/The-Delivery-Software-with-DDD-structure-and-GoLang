// Package customer holds aggregates that combines many entities into a full object.
package customer

//TODO: implement set and get addresses
//TODO: implement logical delete

import (
	"errors"

	"entity"

	"github.com/google/uuid"
)

//Customer contains all personal data about a customer.
type Customer struct {
	person    *entity.Person
	addresses []*entity.Address
}

var (
	//ErrInvalidFirstName is returned when a first name is empty.
	ErrInvalidFirstName = errors.New("firstname is required")
	//ErrInvalidLastName is returned when a last name is empty.
	ErrInvalidLastName = errors.New("lastName is required")
	//ErrInvalidPhoneNumber is returned when a phone number is empty.
	ErrInvalidPhoneNumber = errors.New("phoneNumber is required")
	//ErrInvalidEmail is returned when a email is empty.
	ErrInvalidEmail = errors.New("email is required")
	//ErrOverwritePersonID is returned when a valid id wants changes.
	ErrOverwritePersonID = errors.New("customer identifier cannot overwrite")
	//ErrIDIsNotValid is returned when id is empty.
	ErrIDIsNotValid = errors.New("customer identifier is not valid")
)

//NewCustomer is a constructs function to make a new customer object.
func NewCustomer(firsName, lastName, phoneNumber, email string) (*Customer, error) {
	customer := new(Customer)
	newPerson := &entity.Person{FirstName: firsName,
		LastName:    lastName,
		PhoneNumber: phoneNumber,
		Email:       email,
	}
	if err := checkCustomerRequirdField(newPerson); err != nil {
		return customer, err
	}

	customer = &Customer{person: newPerson,
		addresses: make([]*entity.Address, 0),
	}
	return customer, nil
}

//GetID returns the customer ID.
func (customer *Customer) GetID() uuid.UUID {
	return customer.person.ID
}

//SetID sets the customer ID.
func (customer *Customer) SetID(id uuid.UUID) error {
	customer.createPersonEntityIfNotExists()

	if personIdIsNotEmpty(customer.person.ID) {
		return ErrOverwritePersonID
	}
	if newIdIsNotValid(id) {
		return ErrIDIsNotValid
	}
	customer.person.ID = id
	return nil
}

//GetFirstName returns the customer first name.
func (customer *Customer) GetFirstName() string {
	return customer.person.FirstName
}

//SetFirstName sets the customer first name.
func (customer *Customer) SetFirstName(firstName string) error {
	customer.createPersonEntityIfNotExists()
	if stringValueIsEmpty(firstName) {
		return ErrInvalidFirstName
	}
	customer.person.FirstName = firstName
	return nil
}

//GetLasttName returns the customer last name.
func (customer *Customer) GetLasttName() string {
	return customer.person.LastName
}

//SetLastName sets the customer last name.
func (customer *Customer) SetLastName(lastName string) error {
	customer.createPersonEntityIfNotExists()
	if stringValueIsEmpty(lastName) {
		return ErrInvalidLastName
	}
	customer.person.LastName = lastName
	return nil
}

//GetPhoneNumber returns the customer phone number.
func (customer *Customer) GetPhoneNumber() string {
	return customer.person.PhoneNumber
}

//SetPhoneNumber sets the customer phone number.
func (customer *Customer) SetPhoneNumber(phoneNumber string) error {
	customer.createPersonEntityIfNotExists()
	if stringValueIsEmpty(phoneNumber) {
		return ErrInvalidPhoneNumber
	}
	customer.person.PhoneNumber = phoneNumber
	return nil
}

//GetEmail returns the customer email.
func (customer *Customer) GetEmail() string {
	return customer.person.Email
}

//SetEmail sets the customer email.
func (customer *Customer) SetEmail(email string) error {
	customer.createPersonEntityIfNotExists()
	if stringValueIsEmpty(email) {
		return ErrInvalidEmail
	}
	customer.person.Email = email
	return nil
}

func personIdIsNotEmpty(id uuid.UUID) bool {
	return id != uuid.Nil
}
func newIdIsNotValid(id uuid.UUID) bool {
	return id == uuid.Nil
}
func checkCustomerRequirdField(newPerson *entity.Person) error {
	if stringValueIsEmpty(newPerson.FirstName) {
		return ErrInvalidFirstName
	}
	if stringValueIsEmpty(newPerson.LastName) {
		return ErrInvalidLastName
	}
	if stringValueIsEmpty(newPerson.PhoneNumber) {
		return ErrInvalidPhoneNumber
	}
	if stringValueIsEmpty(newPerson.Email) {
		return ErrInvalidEmail
	}
	return nil
}
func (customer *Customer) createPersonEntityIfNotExists() {
	if customer.person == nil {
		customer.person = new(entity.Person)
	}
}
func stringValueIsEmpty(value string) bool {
	return value == ""
}
