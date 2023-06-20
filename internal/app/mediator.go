package app

import "base/internal/devices"

type Mediator interface {
	SendCommand(sender devices.SmartDevice, command string, targetDeviceID string) error
	RegisterDevice(device devices.SmartDevice) error
	DeregisterDevice(device devices.SmartDevice) error
}
