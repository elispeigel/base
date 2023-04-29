package app

import (
	"base/internal/devices"
	"testing"

	"github.com/stretchr/testify/assert"
)


type MockDevice struct {
	devices.DeviceBase
}

func (d *MockDevice) TurnOn() error {
	return nil
}

func (d *MockDevice) TurnOff() error {
	return nil
}

func (d *MockDevice) GetStatus() (devices.Status, error) {
	return devices.Status{
		DeviceID:   d.ID,
		DeviceName: d.Name,
		IsOn:       false,
	}, nil
}

func TestBaseController(t *testing.T) {
	device1 := &MockDevice{devices.DeviceBase{ID: "device1", Name: "Mock Device 1"}}
	device2 := &MockDevice{devices.DeviceBase{ID: "device2", Name: "Mock Device 2"}}

	// Test GetBaseControllerInstance (singleton)
	controller1 := GetBaseControllerInstance()
	controller2 := GetBaseControllerInstance()
	assert.Equal(t, controller1, controller2)

	// Test AddDevice
	controller1.AddDevice(device1)
	controller1.AddDevice(device2)

	// Test GetDevice
	retrievedDevice1 := controller1.GetDevice("device1")
	retrievedDevice2 := controller1.GetDevice("device2")
	assert.Equal(t, device1, retrievedDevice1)
	assert.Equal(t, device2, retrievedDevice2)

	// Test RemoveDevice
	controller1.RemoveDevice("device1")
	retrievedDevice1 = controller1.GetDevice("device1")
	assert.Nil(t, retrievedDevice1)
}

