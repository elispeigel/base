package speaker

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

type SmartSpeaker struct {
	ID     string
	Name   string
	IsOn   bool
	Volume int
}

// TurnOn turns on the smart speaker and logs a message
func (s *SmartSpeaker) TurnOn() error {
	s.IsOn = true
	logger.Info("Turned on the smart speaker", zap.String("speakerID", s.ID), zap.String("speakerName", s.Name))

	return nil
}

// TurnOff turns off the smart speaker and logs a message
func (s *SmartSpeaker) TurnOff() error {
	s.IsOn = false
	logger.Info("Turned off the smart speaker", zap.String("speakerID", s.ID), zap.String("speakerName", s.Name))

	return nil
}

// GetStatus returns the current status of the smart speaker and logs a message
func (s *SmartSpeaker) GetStatus() (devices.Status, error) {
	status := devices.Status{
		DeviceID:   s.ID,
		DeviceName: s.Name,
		IsOn:       s.IsOn,
	}

	logger.Info("Retrieved status of the smart speaker", zap.String("speakerID", s.ID), zap.String("speakerName", s.Name))

	return status, nil
}

// GetVolume returns the current volume of the smart speaker
func (s *SmartSpeaker) GetVolume() int {
	return s.Volume
}

// SetVolume sets the volume of the smart speaker and logs a message
func (s *SmartSpeaker) SetVolume(volume int) error {
	if volume < 0 || volume > 100 {
		return fmt.Errorf("Invalid volume value. Volume must be between 0 and 100.")
	}

	s.Volume = volume

	logger.Info("Changed the volume of the smart speaker", zap.String("speakerID", s.ID), zap.String("speakerName", s.Name), zap.Int("volume", s.Volume))

	return nil
}
