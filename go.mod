module github.com/vvlisn/ingress-policy

// TinyGo does not support Go v1.26. Therefore, let's keep go and toolchain
// versions to 1.25 to ensure that all tinygo and the standard go command (used
// in makefile) behavie in the same way
go 1.25

toolchain go1.25.7

require (
	github.com/kubewarden/gjson v1.7.2
	github.com/kubewarden/k8s-objects v1.29.0-kw1
	github.com/kubewarden/policy-sdk-go v0.12.0
	github.com/wapc/wapc-guest-tinygo v0.3.3
)

require (
	github.com/go-openapi/strfmt v0.21.3 // indirect
	github.com/tidwall/match v1.0.3 // indirect
	github.com/tidwall/pretty v1.0.2 // indirect
)

replace github.com/go-openapi/strfmt => github.com/kubewarden/strfmt v0.1.3
