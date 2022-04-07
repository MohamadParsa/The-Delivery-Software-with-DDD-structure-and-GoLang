// Package entity holds all the entities that are shared across all subdomains
package entity

import (
	"time"

	"github.com/google/uuid"
)

//PackStatus is an enum that contains "unknown","not sent","sent","delivered"
type PackStatus string

//constant values for PackStatus enums.
const (
	UnknownStatus PackStatus = "unknown"
	NotSent       PackStatus = "not sent"
	Sent          PackStatus = "sent"
	Delivered     PackStatus = "delivered"
)

/*
Pack is an entity that represents a pack in all domain.
*/
type Delivery struct {
	//ID is the identifier of the Entity, the ID is shared for all sub domains.
	ID uuid.UUID
	//ChosenDeliveryDate is the chosen date and time to deliver the pack.
	ChosenDeliveryDate time.Time
	//DeliveryDate is the actual date and time to deliver the pack.
	DeliveryDate time.Time
	//Status is the pack sending status, contains "unknown","not sent","sent","delivered".
	Status PackStatus
	//Cost is the cost of pack sending.
	Cost float64
}
