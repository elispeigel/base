package light

import (
	"base/internal/devices"
)

// LightFactory is an interface that defines the method for creating a smart light
type LightFactory interface {
	CreateSmartLight() *SmartLight
}

// AbstractLightFactory is an interface that defines the method for creating a light factory
type AbstractLightFactory interface {
	CreateLightFactory() LightFactory
}

// ConcreteLightFactory is a struct that implements the LightFactory interface
type ConcreteLightFactory struct{}

// CreateSmartLight creates a new smart light with a unique ID, default name, and off status
func (clf *ConcreteLightFactory) CreateSmartLight() *SmartLight {
	return &SmartLight{
		ID:   devices.GenerateUniqueID(),
		Name: "Default Light",
		IsOn: false,
	}
}

// ConcreteAbstractLightFactory is a struct that implements the AbstractLightFactory interface
type ConcreteAbstractLightFactory struct{}

// CreateLightFactory creates a new ConcreteLightFactory instance
func (calf *ConcreteAbstractLightFactory) CreateLightFactory() LightFactory {
	return &ConcreteLightFactory{}
}