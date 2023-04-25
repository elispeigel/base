package devices

import (
	"github.com/google/uuid"
)

// SmartDevice is an interface that defines the methods that a smart device must implement
type SmartDevice interface {
	TurnOn() error
	TurnOff() error
	GetStatus() (Status, error)
	GetID() string
	GetName() string
}

// Status represents the current status of a smart device
type Status struct {
	DeviceID   string
	DeviceName string
	IsOn       bool
}

// GenerateUniqueID generates a unique ID using the UUID package
func GenerateUniqueID() string {
	u := uuid.New()
	return u.String()
}
