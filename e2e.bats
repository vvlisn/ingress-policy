#!/usr/bin/env bats

@test "reject when defaultBackend is configured" {
  run kwctl run annotated-policy.wasm -r test_data/ingress-with-default-backend.json --settings-json '{"denyDefaultBackend": true}'

  # this prints the output when one the checks below fails
  echo "output = ${output}"

  # request rejected
  [ "$status" -eq 0 ]
  [ $(expr "$output" : '.*allowed.*false') -ne 0 ]
  [ $(expr "$output" : '.*defaultBackend is not allowed.*') -ne 0 ]

}

@test "reject because invalid settings" {
  run kwctl run annotated-policy.wasm -r test_data/ingress-without-default-backend.json --settings-json '{"requireTLS": true}'

  # this prints the output when one the checks below fails
  echo "output = ${output}"

  # settings validation fails
  [ "$status" -eq 1 ]
}

@test "accept" {
  run kwctl run annotated-policy.wasm -r test_data/ingress-without-default-backend.json --settings-json '{"denyDefaultBackend": true}'
  # this prints the output when one the checks below fails
  echo "output = ${output}"

  # request accepted
  [ "$status" -eq 0 ]
  [ $(expr "$output" : '.*allowed.*true') -ne 0 ]
}

@test "accept when defaultBackend check is disabled" {
  run kwctl run annotated-policy.wasm -r test_data/ingress-with-default-backend.json --settings-json '{"denyDefaultBackend": false}'
  echo "output = ${output}"

  [ "$status" -eq 0 ]
  [ $(expr "$output" : '.*allowed.*true') -ne 0 ]
}
