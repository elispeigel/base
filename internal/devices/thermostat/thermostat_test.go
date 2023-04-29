package thermostat

import (
	"base/internal/devices"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
)

func TestSmarThermostat(t *testing.T) {
	logger = zaptest.NewLogger(t)

	thermostat := SmartThermostat{
		devices.DeviceBase{
			ID:   "123",
			Name: "Test Thermostat",
		},
		false,
		50,
	}

	// Test TurnOn
	err := thermostat.TurnOn()
	assert.NoError(t, err)
	assert.True(t, thermostat.IsOn)

	// Test TurnOff
	err = thermostat.TurnOff()
	assert.NoError(t, err)
	assert.False(t, thermostat.IsOn)

	// Test GetStatus
	status, err := thermostat.GetStatus()
	assert.NoError(t, err)
	expectedStatus := devices.Status{
		DeviceID:   "123",
		DeviceName: "Test Thermostat",
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
}
