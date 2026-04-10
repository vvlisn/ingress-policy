package main

import (
	"encoding/json"
	"testing"

	networkingv1 "github.com/kubewarden/k8s-objects/api/networking/v1"
	metav1 "github.com/kubewarden/k8s-objects/apimachinery/pkg/apis/meta/v1"
	kubewarden_protocol "github.com/kubewarden/policy-sdk-go/protocol"
	kubewarden_testing "github.com/kubewarden/policy-sdk-go/testing"
)

func TestParsingEmptySettingsFromValidationReq(t *testing.T) {
	ingress := networkingv1.Ingress{
		Metadata: &metav1.ObjectMeta{
			Name:      "test",
			Namespace: "default",
		},
	}

	validationReqRaw, err := kubewarden_testing.BuildValidationRequest(ingress, &Settings{})
	if err != nil {
		t.Errorf("Unexpected error %+v", err)
	}

	validationReq := kubewarden_protocol.ValidationRequest{}
	err = json.Unmarshal(validationReqRaw, &validationReq)
	if err != nil {
		t.Errorf("Unexpected error %+v", err)
	}

	settings, err := NewSettingsFromValidationReq(&validationReq)
	if err != nil {
		t.Errorf("Unexpected error %+v", err)
	}

	if !settings.shouldDenyDefaultBackend() {
		t.Errorf("Expected denyDefaultBackend to default to true")
	}
}

func TestParsingEmptySettingsPayload(t *testing.T) {
	settings, err := parseSettings([]byte(`{}`))
	if err != nil {
		t.Errorf("Unexpected error %+v", err)
	}

	if !settings.shouldDenyDefaultBackend() {
		t.Errorf("Expected denyDefaultBackend to default to true")
	}
}

func TestParsingSettingsWithDisabledCheck(t *testing.T) {
	settings, err := parseSettings([]byte(`{"denyDefaultBackend": false}`))
	if err != nil {
		t.Errorf("Unexpected error %+v", err)
	}

	if settings.shouldDenyDefaultBackend() {
		t.Errorf("Expected denyDefaultBackend to be disabled")
	}
}

func TestRejectingUnknownSettingsFields(t *testing.T) {
	if _, err := parseSettings([]byte(`{"requireTLS": true}`)); err == nil {
		t.Errorf("Expected settings parsing to fail")
	}
}
