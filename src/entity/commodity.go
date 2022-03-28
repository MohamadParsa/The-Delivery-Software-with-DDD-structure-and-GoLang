//Package entity holds all the entities that are shared across all subdomains
package entity

import (
	"github.com/google/uuid"
)

//Commodity represents a commodity for all sub domains
type Commodity struct {
	//ID is the identifier of the Entity, the ID is shared for all sub domains.
	ID uuid.UUID
	//Name is the name of the commodity.
	Name string
	//Model is the model of the commodity.
	Model string
	//CompanyName is the manufacturer name of the commodity.
	CompanyName string
	//Description is a description for the commodity.
	Description string
}
