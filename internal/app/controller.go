package app

import (
	"base/internal/devices"
	"sync"
)

type BaseController struct {
	devices map[string]devices.SmartDevice
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
	hc.devices[device.GetID()] = device
}

// RemoveDevice removes a smart device from the base controller.
func (hc *BaseController) RemoveDevice(deviceID string) {
	delete(hc.devices, deviceID)
}

// GetDevice returns a smart device from the base controller by its ID.
func (hc *BaseController) GetDevice(deviceID string) devices.SmartDevice {
	return hc.devices[deviceID]
}
