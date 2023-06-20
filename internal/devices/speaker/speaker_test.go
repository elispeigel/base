package speaker

import (
	"base/internal/app"
	"base/internal/devices"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
)

func TestSmartSpeaker(t *testing.T) {
	logger = zaptest.NewLogger(t)

	// Get BaseController as mediator
	mediator := app.GetBaseControllerInstance()

	// Factory for creating SmartSpeaker
	factory := ConcreteSpeakerFactory{}

	// Create SmartSpeaker using the factory
	speaker := factory.CreateSmartSpeaker(mediator)

	// Test RegisterDevice
	err := mediator.RegisterDevice(speaker)
	assert.NoError(t, err)

	//Test TurnOn
	err = mediator.SendCommand(speaker, "turnOn", speaker.GetID())
	assert.NoError(t, err)
	assert.True(t, speaker.IsOn)

	// Test TurnOff
	err = mediator.SendCommand(speaker, "turnOff", speaker.GetID())
	assert.NoError(t, err)
	assert.False(t, speaker.IsOn)

	// Test GetStatus
	status, err := speaker.GetStatus()
	assert.NoError(t, err)
	expectedStatus := devices.Status{
		DeviceID:   speaker.GetID(),
		DeviceName: "Default Speaker",
		IsOn:       false,
	}
	assert.Equal(t, expectedStatus, status)

	// Test GetVolume
	assert.Equal(t, 50, speaker.GetVolume())

	// Test SetVolume
	err = speaker.SetVolume(80)
	assert.NoError(t, err)
	assert.Equal(t, 80, speaker.GetVolume())

	// Test SetVolume with invalid value
	err = speaker.SetVolume(-1)
	assert.Error(t, err)

	// Test DeregisterDevice
	err = mediator.DeregisterDevice(speaker)
	assert.NoError(t, err)
}
