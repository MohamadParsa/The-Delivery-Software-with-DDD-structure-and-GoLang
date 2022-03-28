// Package entity holds all the entities that are shared across all subdomains
package entity

import (
	"github.com/google/uuid"
)

/*
Address is an entity that represents an address
*/
type Address struct {
	//ID is the identifier of the Entity, the ID is shared for all sub domains.
	ID uuid.UUID
	//Country is the country name.
	Country string
	//State is the state name.
	State string
	//City is the city name.
	City string
	//LocalAddress contain the steet and alley.
	LocalAddress string
	//ZipCode is the zip code of the address.
	ZipCode int
	//GeographicalCoordinates is the geographical coordinates of the address.
	GeographicalCoordinates [2]float64
}
