package thermostat

import (
	"base/internal/app"
	"base/internal/devices"
)

// ThermostatFactory is an interface that defines the method for creating a smart thermostat
type ThermostatFactory interface {
	CreateSmartThermostat(mediator app.Mediator) *SmartThermostat
}

// AbstractThermostatFactory is an interface that defines the method for creating a thermostat factory
type AbstractThermostatFactory interface {
	CreateThermostatFactory() ThermostatFactory
}

// ConcreteThermostatFactory is a struct that implements the ThermostatFactory interface
type ConcreteThermostatFactory struct{}

// CreateSmartThermostat creates a new smart thermostat with a unique ID, default name, off status, and volume
func (clf *ConcreteThermostatFactory) CreateSmartThermostat(mediator app.Mediator) *SmartThermostat {
	return &SmartThermostat{
		devices.DeviceBase{
			ID:   devices.GenerateUniqueID(),
			Name: "Default Thermostat",
		},
		mediator,
		false,
		50,
	}
}

// ConcreteAbstractThermostatFactory is a struct that implements the AbstractThermostatFactory interface
type ConcreteAbstractThermostatFactory struct{}

// CreateThermostatFactory creates a new ConcreteThermostatFactory instance
func (calf *ConcreteAbstractThermostatFactory) CreateThermostatFactory() ThermostatFactory {
	return &ConcreteThermostatFactory{}
}
