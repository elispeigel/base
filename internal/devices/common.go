package devices

import (
	"github.com/google/uuid"
)

type SmartDevice interface {
	TurnOn() error
	TurnOff() error
	GetStatus() (Status, error)
	GetID() string
	GetName() string
}

type Status struct {
	DeviceID   string
	DeviceName string
	IsOn       bool
}

func GenerateUniqueID() string {
	u := uuid.New()
	return u.String()
}
