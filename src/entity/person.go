// Package entity holds all the entities that are shared across all subdomains.
package entity

import (
	"github.com/google/uuid"
)

/*
Person is an entity that represents a person in all domain.
*/
type Person struct {
	//ID is the identifier of the Entity, the ID is shared for all sub domains.
	ID uuid.UUID
	//FirsName is the first name of the person.
	FirstName string
	//LastName is the last name of the person.
	LastName string
	//PhoneNumber is the phone number of the person that contains the area code too.
	PhoneNumber string
	//Email is the email of the person.
	Email string
}
