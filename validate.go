package main

import (
	"encoding/json"

	"github.com/kubewarden/gjson"
	kubewarden "github.com/kubewarden/policy-sdk-go"
	kubewarden_protocol "github.com/kubewarden/policy-sdk-go/protocol"
)

const defaultBackendForbiddenMessage = "defaultBackend is not allowed"
const badRequestStatusCode = 400

func validate(payload []byte) ([]byte, error) {
	validationRequest := kubewarden_protocol.ValidationRequest{}
	err := json.Unmarshal(payload, &validationRequest)
	if err != nil {
		return kubewarden.RejectRequest(
			kubewarden.Message(err.Error()),
			kubewarden.Code(badRequestStatusCode))
	}

	settings, err := NewSettingsFromValidationReq(&validationRequest)
	if err != nil {
		return kubewarden.RejectRequest(
			kubewarden.Message(err.Error()),
			kubewarden.Code(badRequestStatusCode))
	}

	if settings.shouldDenyDefaultBackend() && hasDefaultBackend(payload) {
		return kubewarden.RejectRequest(
			kubewarden.Message(defaultBackendForbiddenMessage),
			kubewarden.NoCode)
	}

	return kubewarden.AcceptRequest()
}

func hasDefaultBackend(payload []byte) bool {
	return gjson.GetBytes(payload, "request.object.spec.defaultBackend").Exists()
}
