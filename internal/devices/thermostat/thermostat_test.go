package thermostat

import (
	"base/internal/app"
	"base/internal/devices"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
)

func TestSmarThermostat(t *testing.T) {
	logger = zaptest.NewLogger(t)

	// Get BaseController as mediator
	mediator := app.GetBaseControllerInstance()

	// Factory for creating SmartThermostat
	factory := ConcreteThermostatFactory{}

	// Create SmartThermostat using the factory
	thermostat := factory.CreateSmartThermostat(mediator)

	// Test RegisterDevice
	err := mediator.RegisterDevice(thermostat)
	assert.NoError(t, err)

	// Test TurnOn
	err = mediator.SendCommand(thermostat, "turnOn", thermostat.GetID())
	assert.NoError(t, err)
	assert.True(t, thermostat.IsOn)

	// Test TurnOff
	err = mediator.SendCommand(thermostat, "turnOff", thermostat.GetID())
	assert.NoError(t, err)
	assert.False(t, thermostat.IsOn)

	// Test GetStatus
	status, err := thermostat.GetStatus()
	assert.NoError(t, err)
	expectedStatus := devices.Status{
		DeviceID:   thermostat.GetID(),
		DeviceName: "Default Thermostat",
		IsOn:       false,
	}
	assert.Equal(t, expectedStatus, status)

	// Test GetTemperature
	assert.Equal(t, 50, thermostat.GetTemperature())

	// Test SetTemperature
	err = thermostat.SetTemperature(80)
	assert.NoError(t, err)
	assert.Equal(t, 80, thermostat.GetTemperature())

	// Test SetTemperature with invalid value
	err = thermostat.SetTemperature(-1)
	assert.Error(t, err)

	// Test DeregisterDevice
	err = mediator.DeregisterDevice(thermostat)
	assert.NoError(t, err)
}
