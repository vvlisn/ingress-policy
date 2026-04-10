package main

import (
	"bytes"
	"encoding/json"

	kubewarden "github.com/kubewarden/policy-sdk-go"
	kubewarden_protocol "github.com/kubewarden/policy-sdk-go/protocol"
)

type Settings struct {
	DenyDefaultBackend *bool `json:"denyDefaultBackend,omitempty"`
}

func NewSettingsFromValidationReq(validationReq *kubewarden_protocol.ValidationRequest) (Settings, error) {
	return parseSettings(validationReq.Settings)
}

func parseSettings(payload []byte) (Settings, error) {
	if len(bytes.TrimSpace(payload)) == 0 {
		return Settings{}, nil
	}

	decoder := json.NewDecoder(bytes.NewReader(payload))
	decoder.DisallowUnknownFields()

	settings := Settings{}
	if err := decoder.Decode(&settings); err != nil {
		return Settings{}, err
	}

	return settings, nil
}

func validateSettings(payload []byte) ([]byte, error) {
	if _, err := parseSettings(payload); err != nil {
		return []byte{}, err
	}

	return kubewarden.AcceptSettings()
}

func (s Settings) shouldDenyDefaultBackend() bool {
	return s.DenyDefaultBackend == nil || *s.DenyDefaultBackend
}
