//Package entity holds all the entities that are shared across all subdomains
package entity

import (
	"github.com/google/uuid"
)

//Currencies is a currencies enum that contains "unknown","EUR","USD","GBP"
type Currencies string

//constant values for Currencies enums.
const (
	UnknownCurrency Currencies = "unknown"
	EUR             Currencies = "EUR"
	USD             Currencies = "USD"
	GBP             Currencies = "GBP"
)

//Price represents a Price for all sub domains
type Price struct {
	//ID is the identifier of the Entity, the ID is shared for all sub domains.
	ID uuid.UUID
	//Amount is the amount of the price.
	Amount float64
	//Discount is the discount of the price.
	Discount float32
	//Currency is the currency of the price.
	Currency Currencies
}
