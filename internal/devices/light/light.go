package light

import (
	"base/internal/devices"

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

// SmartLight is a struct that represents a smart light
type SmartLight struct {
	ID         string
	Name       string
	IsOn       bool
	Brightness int
	Color      string
}

// TurnOn turns on the smart light and logs a message
func (l *SmartLight) TurnOn() error {
	l.IsOn = true
	logger.Info("Turned on the smart light", zap.String("lightID", l.ID), zap.String("lightName", l.Name))

	return nil
}

// TurnOff turns off the smart light and logs a message
func (l *SmartLight) TurnOff() error {
	l.IsOn = false
	logger.Info("Turned off the smart light", zap.String("lightID", l.ID), zap.String("lightName", l.Name))

	return nil
}

// GetStatus returns the current status of the smart light and logs a message
func (l *SmartLight) GetStatus() (devices.Status, error) {
	status := devices.Status{
		DeviceID:   l.ID,
		DeviceName: l.Name,
		IsOn:       l.IsOn,
	}

	logger.Info("Retrieved status of the smart light", zap.String("lightID", l.ID), zap.String("lightName", l.Name))

	return status, nil
}

// GetID returns the ID of the smart light
func (l *SmartLight) GetID() string {
	return l.ID
}

// GetName returns the name of the smart light
func (l *SmartLight) GetName() string {
	return l.Name
}
