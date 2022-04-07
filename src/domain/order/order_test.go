package order

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewOrder(t *testing.T) {

	type TestCase struct {
		test        string
		customerID  uuid.UUID
		expectedErr error
	}
	customerID, err := uuid.NewUUID()
	checkError(err, t)
	testCases := []TestCase{
		{
			test:        "empty customer id",
			expectedErr: ErrEmptyCustomerID,
		},
		{
			test:        "new order with valid data",
			customerID:  customerID,
			expectedErr: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.test, func(t *testing.T) {
			_, err := NewOrder(testCase.customerID)
			if err != testCase.expectedErr {
				t.Errorf("expected error :%v, got %v", testCase.expectedErr, err)
			}
		})
	}
}

func TestSetOrderID(t *testing.T) {

	type TestCase struct {
		test        string
		order       *Order
		orderID     uuid.UUID
		expectedErr error
	}
	customerID, err := uuid.NewUUID()
	checkError(err, t)
	orderID, err := uuid.NewUUID()
	checkError(err, t)
	orderWithOrderID, err := NewOrder(customerID)
	checkError(err, t)
	err = orderWithOrderID.SetID(orderID)
	checkError(err, t)
	orderWithOutOrderID, err := NewOrder(customerID)
	checkError(err, t)

	checkError(err, t)
	testCases := []TestCase{
		{
			test:        "empty order id",
			orderID:     uuid.Nil,
			order:       orderWithOutOrderID,
			expectedErr: ErrEmptyOrderID,
		},
		{
			test:        "overwite order id",
			orderID:     orderID,
			order:       orderWithOrderID,
			expectedErr: ErrOverwriteOrderID,
		},
		{
			test:        "set valid order id",
			orderID:     orderID,
			order:       orderWithOutOrderID,
			expectedErr: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.test, func(t *testing.T) {
			err := testCase.order.SetID(testCase.orderID)
			if err != testCase.expectedErr {
				t.Errorf("expected error :%v, got %v", testCase.expectedErr, err)
			}
		})
	}
}
func TestGetOrderID(t *testing.T) {

	type TestCase struct {
		test    string
		order   *Order
		orderID uuid.UUID
	}
	customerID, err := uuid.NewUUID()
	checkError(err, t)
	orderID, err := uuid.NewUUID()
	checkError(err, t)
	orderWithOrderID, err := NewOrder(customerID)
	checkError(err, t)
	err = orderWithOrderID.SetID(orderID)
	checkError(err, t)
	orderWithOutOrderID, err := NewOrder(customerID)
	checkError(err, t)
	testCases := []TestCase{
		{
			test:    "empty order id",
			orderID: uuid.Nil,
			order:   orderWithOutOrderID,
		},
		{
			test:    "get valid order id",
			orderID: orderID,
			order:   orderWithOrderID,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.test, func(t *testing.T) {
			id := testCase.order.GetID()
			if id != testCase.orderID {
				t.Errorf("expected order id :%v, got %v", orderID, id)
			}
		})
	}
}
func checkError(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}
}
