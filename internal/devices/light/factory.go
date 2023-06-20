package light

import (
	"base/internal/app"
	"base/internal/devices"
	"image/color"
)

// LightFactory is an interface that defines the method for creating a smart light
type LightFactory interface {
	CreateSmartLight(mediator app.Mediator) *SmartLight
}

// AbstractLightFactory is an interface that defines the method for creating a light factory
type AbstractLightFactory interface {
	CreateLightFactory() LightFactory
}

// ConcreteLightFactory is a struct that implements the LightFactory interface
type ConcreteLightFactory struct{}

// CreateSmartLight creates a new smart light with a unique ID, default name, and off status
func (clf *ConcreteLightFactory) CreateSmartLight(mediator app.Mediator) *SmartLight {
	return &SmartLight{
		devices.DeviceBase{
			ID:   devices.GenerateUniqueID(),
			Name: "Default Light",
		},
		mediator,
		false,
		50,
		color.RGBA{uint8(255), uint8(255), uint8(255), uint8(1)},
	}
}

// ConcreteAbstractLightFactory is a struct that implements the AbstractLightFactory interface
type ConcreteAbstractLightFactory struct{}

// CreateLightFactory creates a new ConcreteLightFactory instance
func (calf *ConcreteAbstractLightFactory) CreateLightFactory() LightFactory {
	return &ConcreteLightFactory{}
}
