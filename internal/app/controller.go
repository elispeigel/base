package app

import (
	"base/internal/devices"
	"fmt"
	"sync"
)

type BaseController struct {
    sync.RWMutex
    devices map[string]devices.SmartDevice
}

func (hc *BaseController) SendCommand(sender devices.SmartDevice, command string, targetDeviceID string) error {
	targetDevice, ok := hc.devices[targetDeviceID]
	if !ok {
		return fmt.Errorf("Target device not found: %s", targetDeviceID)
	}

	if command == "turnOn" {
		return targetDevice.TurnOn()
	} else if command == "turnOff" {
		return targetDevice.TurnOff()
	}

	return nil
}

func (hc *BaseController) RegisterDevice(device devices.SmartDevice) error {
    hc.AddDevice(device)
    return nil
}

func (hc *BaseController) DeregisterDevice(device devices.SmartDevice) error {
    deviceID := device.GetID()
    if _, ok := hc.devices[deviceID]; !ok {
        return fmt.Errorf("Device not found: %s", deviceID)
    }

    hc.RemoveDevice(deviceID)
    return nil
}


// NewBaseController initializes a new BaseController instance.
func NewBaseController() *BaseController {
	return &BaseController{
		devices: make(map[string]devices.SmartDevice),
	}
}

var baseControllerInstance *BaseController
var once sync.Once

// GetBaseControllerInstance returns the singleton instance of BaseController.
func GetBaseControllerInstance() *BaseController {
	once.Do(func() {
		baseControllerInstance = NewBaseController()
	})
	return baseControllerInstance
}

// AddDevice adds a new smart device to the base controller.
func (hc *BaseController) AddDevice(device devices.SmartDevice) {
	hc.Lock()
    defer hc.Unlock()
	if hc.devices == nil {
        hc.devices = make(map[string]devices.SmartDevice)
    }
	hc.devices[device.GetID()] = device
}

// RemoveDevice removes a smart device from the base controller.
func (hc *BaseController) RemoveDevice(deviceID string) {
	hc.Lock()
    defer hc.Unlock()
	delete(hc.devices, deviceID)
}

// GetDevice returns a smart device from the base controller by its ID.
func (hc *BaseController) GetDevice(deviceID string) devices.SmartDevice {
	return hc.devices[deviceID]
}
