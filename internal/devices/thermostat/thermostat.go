package thermostat

import (
	"base/internal/devices"
	"fmt"

	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
}

type SmartThermostat struct {
    devices.DeviceBase
	IsOn bool
	Temperature int
}

// TurnOn turns on the smart thermostat and logs a message
func (t *SmartThermostat) TurnOn() error {
	t.IsOn = true
	logger.Info("Turned on the smart thermostat", zap.String("thermostatID", t.ID), zap.String("thermostatName", t.Name))

	return nil
}

// TurnOff turns off the smart thermostat and logs a message
func (t *SmartThermostat) TurnOff() error {
	t.IsOn = false
	logger.Info("Turned off the smart thermostat", zap.String("thermostatID", t.ID), zap.String("thermostatName", t.Name))

	return nil
}

// GetStatus returns the current status of the smart thermostat and logs a message
func (t *SmartThermostat) GetStatus() (devices.Status, error) {
	status := devices.Status{
		DeviceID:   t.ID,
		DeviceName: t.Name,
		IsOn:       t.IsOn,
	}

	logger.Info("Retrieved status of the smart thermostat", zap.String("thermostatID", t.ID), zap.String("thermostatName", t.Name))

	return status, nil
}

// GetTemperature returns the current temperature of the smart thermostat
func (t *SmartThermostat) GetTemperature() int {
	return t.Temperature
}

func (t *SmartThermostat) SetTemperature(temperature int) error {
	if temperature < 0 || temperature > 100 {
		return fmt.Errorf("Invalid temperature value. Temperature must be between 0 and 100.")
	}

	t.Temperature = temperature

	logger.Info("Changed the volume of the smart thermostat", zap.String("thermostatID", t.ID), zap.String("thermostatName", t.Name), zap.Int("temeprature", t.Temperature))

	return nil
}