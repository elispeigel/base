package speaker

import (
	"base/internal/devices"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
)

func TestSmartSpeaker(t *testing.T) {
	logger = zaptest.NewLogger(t)

	speaker := SmartSpeaker{
		devices.DeviceBase{
			ID:   "123",
			Name: "Test Speaker",
		},
		false,
		50,
	}

	//Test TurnOn
	err := speaker.TurnOn()
	assert.NoError(t, err)
	assert.True(t, speaker.IsOn)

	// Test TurnOff
	err = speaker.TurnOff()
	assert.NoError(t, err)
	assert.False(t, speaker.IsOn)

	// Test GetStatus
	status, err := speaker.GetStatus()
	assert.NoError(t, err)
	expectedStatus := devices.Status{
		DeviceID:   "123",
		DeviceName: "Test Speaker",
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
}
