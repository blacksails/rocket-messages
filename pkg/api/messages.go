package api

import (
	"encoding/json"
	"time"
)

type Message struct {
	Metadata MessageMetadata `json:"metadata"`
	Message  json.RawMessage `json:"message"`
}

type MessageMetadata struct {
	Channel       string      `json:"channel"`
	MessageNumber int         `json:"messageNumber"`
	MessageTime   time.Time   `json:"messageTime"`
	MessageType   MessageType `json:"messageType"`
}

type MessageType string

const (
	MessageTypeRocketLaunched       MessageType = "RocketLaunched"
	MessageTypeRocketSpeedIncreased MessageType = "RocketSpeedIncreased"
	MessageTypeRocketSpeedDecreased MessageType = "RocketSpeedDecreased"
	MessageTypeRocketExploded       MessageType = "RocketExploded"
	MessageTypeRocketMissionChanged MessageType = "RocketMissionChanged"
)

type MessageRocketLaunched struct {
	Type        string `json:"type"`
	LaunchSpeed int    `json:"launchSpeed"`
	Mission     string `json:"ARTEMIS"`
}

type MessageRocketSpeedChanged struct {
	By int `json:"by"`
}

type MessageRocketExploded struct {
	Reason string `json:"reason"`
}

type MessageRocketMissionChanged struct {
	NewMission string `json:"newMission"`
}
