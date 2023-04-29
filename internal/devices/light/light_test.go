package light

import (
	"base/internal/devices"
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
)

func TestSmartLight(t *testing.T) {
	logger = zaptest.NewLogger(t)

	light := SmartLight{
		ID:         "123",
		Name:       "Test Light",
		IsOn:       false,
		Brightness: 50,
		Color:      color.RGBA{R: 255, G: 255, B: 255, A: 255},
	}

	// Test TurnOn
	err := light.TurnOn()
	assert.NoError(t, err)
	assert.True(t, light.IsOn)

	// Test TurnOff
	err = light.TurnOff()
	assert.NoError(t, err)
	assert.False(t, light.IsOn)

	// Test GetStatus
	status, err := light.GetStatus()
	assert.NoError(t, err)
	expectedStatus := devices.Status{
		DeviceID:   "123",
		DeviceName: "Test Light",
		IsOn:       false,
	}
	assert.Equal(t, expectedStatus, status)

	// Test GetBrightness
	assert.Equal(t, 50, light.GetBrightness())

	// Test SetBrightness
	err = light.SetBrightness(80)
	assert.NoError(t, err)
	assert.Equal(t, 80, light.GetBrightness())

	// Test SetBrightness with invalid value
	err = light.SetBrightness(-1)
	assert.Error(t, err)

	// Test GetColor
	assert.Equal(t, color.RGBA{R: 255, G: 255, B: 255, A: 255}, light.GetColor())

	// Test SetColor
	err = light.SetColor(0, 255, 0, 255)
	assert.NoError(t, err)
	assert.Equal(t, color.RGBA{R: 0, G: 255, B: 0, A: 255}, light.GetColor())

	// Test SetColor with invalid value
	err = light.SetColor(0, -1, 0, 255)
	assert.Error(t, err)
}
