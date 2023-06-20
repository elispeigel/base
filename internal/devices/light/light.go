package light

import (
	"base/internal/app"
	"base/internal/devices"
	"fmt"
	"image/color"

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
	devices.DeviceBase
	mediator   app.Mediator
	IsOn       bool
	Brightness int
	Color      color.RGBA
}

// GetID returns the ID of the smart light
func (l *SmartLight) GetID() string {
	return l.ID
}

// GetName returns the name of the smart light
func (l *SmartLight) GetName() string {
	return l.Name
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

// GetBrightness returns the brightness of the smart light
func (l *SmartLight) GetBrightness() int {
	return l.Brightness
}

// SetBrightness sets the brightness of the smart light and logs a message
func (l *SmartLight) SetBrightness(brightness int) error {
	if brightness < 0 || brightness > 100 {
		return fmt.Errorf("Invalid brightness value. Brightness must be between 0 and 100.")
	}

	l.Brightness = brightness

	logger.Info("Changed the brightness of the smart light", zap.String("lightID", l.ID), zap.String("lightName", l.Name), zap.Int("brightness", l.Brightness))

	return nil
}

// GetColor returns the color of the smart light
func (l *SmartLight) GetColor() color.RGBA {
	return l.Color
}

// SetColor sets the RGBA color of the smart light and logs a message
func (l *SmartLight) SetColor(r int, g int, b int, a int) error {
	for _, num := range [4]int{r, g, b, a} {
		if num < 0 || num > 255 {
			return fmt.Errorf("Invalid Color value. Color must be between 0 and 255 for RGBA.")
		}
	}

	l.Color = color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
	logger.Info("Changed the color of the smart light", zap.String("lightID", l.ID), zap.String("lightName", l.Name), zap.Uint8("lightColor Red", l.Color.R), zap.Uint8("lightColor Green", l.Color.G), zap.Uint8("lightColor Blue", l.Color.B))

	return nil
}
