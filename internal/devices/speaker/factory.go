package speaker

import (
	"base/internal/devices"
)

// SpeakerFactory is an interface that defines the method for creating a smart speaker
type SpeakerFactory interface {
	CreateSmartSpeaker() *SmartSpeaker
}

// AbstractSpeakerFactory is an interface that defines the method for creating a speaker factory
type AbstractSpeakerFactory interface {
	CreateSpeakerFactory() SpeakerFactory
}

// ConcreteSpeakerFactory is a struct that implements the SpeakerFactory interface
type ConcreteSpeakerFactory struct{}

// CreateSmartSpeaker creates a new smart speaker with a unique ID, default name, off status, and volume
func (clf *ConcreteSpeakerFactory) CreateSmartSpeaker() *SmartSpeaker {
	return &SmartSpeaker{
		devices.DeviceBase{
			ID:   devices.GenerateUniqueID(),
			Name: "Default Speaker",
		},
		false,
		50,
	}
}

// ConcreteAbstractSpeakerFactory is a struct that implements the AbstractSpeakerFactory interface
type ConcreteAbstractSpeakerFactory struct{}

// CreateSpeakerFactory creates a new ConcreteSpeakerFactory instance
func (calf *ConcreteAbstractSpeakerFactory) CreateSpeakerFactory() SpeakerFactory {
	return &ConcreteSpeakerFactory{}
}
