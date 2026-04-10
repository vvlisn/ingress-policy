package main

import (
	"encoding/json"
	"testing"

	kubewarden_protocol "github.com/kubewarden/policy-sdk-go/protocol"
	kubewarden_testing "github.com/kubewarden/policy-sdk-go/testing"
)

func TestHasDefaultBackend(t *testing.T) {
	payload, err := kubewarden_testing.BuildValidationRequestFromFixture(
		"test_data/ingress-with-default-backend.json",
		&Settings{})
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	if !hasDefaultBackend(payload) {
		t.Errorf("Expected defaultBackend to be detected")
	}
}

func TestValidationRejectionDueToInvalidJSON(t *testing.T) {
	payload := []byte(`boom baby!`)

	responsePayload, err := validate(payload)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	var response kubewarden_protocol.ValidationResponse
	if unmarshalErr := json.Unmarshal(responsePayload, &response); unmarshalErr != nil {
		t.Errorf("Unexpected error: %+v", unmarshalErr)
	}

	if response.Accepted != false {
		t.Error("Unexpected approval")
	}

	expectedMessage := "invalid character 'b' looking for beginning of value"
	if *response.Message != expectedMessage {
		t.Errorf("Got '%s' instead of '%s'", *response.Message, expectedMessage)
	}
}

func TestValidationRejectionDueToUnknownSettings(t *testing.T) {
	payload, err := kubewarden_testing.BuildValidationRequestFromFixture(
		"test_data/ingress-without-default-backend.json",
		json.RawMessage(`{"requireTLS": true}`))
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	responsePayload, err := validate(payload)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	var response kubewarden_protocol.ValidationResponse
	if unmarshalErr := json.Unmarshal(responsePayload, &response); unmarshalErr != nil {
		t.Errorf("Unexpected error: %+v", unmarshalErr)
	}

	if response.Accepted != false {
		t.Error("Unexpected approval")
	}

	expectedMessage := `json: unknown field "requireTLS"`
	if *response.Message != expectedMessage {
		t.Errorf("Got '%s' instead of '%s'", *response.Message, expectedMessage)
	}
	if response.Code == nil || *response.Code != 400 {
		t.Errorf("Unexpected response code: %+v", response.Code)
	}
}

func TestValidationRejectsIngressWithDefaultBackend(t *testing.T) {
	payload, err := kubewarden_testing.BuildValidationRequestFromFixture(
		"test_data/ingress-with-default-backend.json",
		&Settings{})
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	responsePayload, err := validate(payload)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	var response kubewarden_protocol.ValidationResponse
	if unmarshalErr := json.Unmarshal(responsePayload, &response); unmarshalErr != nil {
		t.Errorf("Unexpected error: %+v", unmarshalErr)
	}

	if response.Accepted != false {
		t.Error("Unexpected approval")
	}

	if *response.Message != defaultBackendForbiddenMessage {
		t.Errorf("Got '%s' instead of '%s'", *response.Message, defaultBackendForbiddenMessage)
	}
}

func TestValidationAcceptsIngressWithoutDefaultBackend(t *testing.T) {
	payload, err := kubewarden_testing.BuildValidationRequestFromFixture(
		"test_data/ingress-without-default-backend.json",
		&Settings{})
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	responsePayload, err := validate(payload)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	var response kubewarden_protocol.ValidationResponse
	if unmarshalErr := json.Unmarshal(responsePayload, &response); unmarshalErr != nil {
		t.Errorf("Unexpected error: %+v", unmarshalErr)
	}

	if response.Accepted != true {
		t.Error("Unexpected rejection")
	}
}

func TestValidationAcceptsIngressWithDefaultBackendWhenCheckDisabled(t *testing.T) {
	payload, err := kubewarden_testing.BuildValidationRequestFromFixture(
		"test_data/ingress-with-default-backend.json",
		json.RawMessage(`{"denyDefaultBackend": false}`))
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	responsePayload, err := validate(payload)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	var response kubewarden_protocol.ValidationResponse
	if unmarshalErr := json.Unmarshal(responsePayload, &response); unmarshalErr != nil {
		t.Errorf("Unexpected error: %+v", unmarshalErr)
	}

	if response.Accepted != true {
		t.Error("Unexpected rejection")
	}
}
