[![Kubewarden Policy Repository](https://github.com/kubewarden/community/blob/main/badges/kubewarden-policies.svg)](https://github.com/kubewarden/community/blob/main/REPOSITORIES.md#policy-scope)
[![Stable](https://img.shields.io/badge/status-stable-brightgreen?style=for-the-badge)](https://github.com/kubewarden/community/blob/main/REPOSITORIES.md#stable)

Kubewarden policy that rejects Ingress resources defining `spec.defaultBackend`.

# Behavior

This policy validates Kubernetes `Ingress` resources.

- If `spec.defaultBackend` is present, the policy rejects the resource.
- If `spec.defaultBackend` is not present, the policy accepts the resource.

# Settings

This policy exposes a single setting:

- `denyDefaultBackend`: `boolean`
  - Default: `true`
  - When `true`, the policy rejects Ingress resources defining `spec.defaultBackend`.
  - When `false`, the policy skips this validation.

Default configuration:

```json
{
  "denyDefaultBackend": true
}
```

To disable the check:

```json
{
  "denyDefaultBackend": false
}
```

Any unknown setting is rejected during settings validation.

# Example

The following Ingress is rejected because it defines `spec.defaultBackend`:

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: app
spec:
  defaultBackend:
    service:
      name: app-service
      port:
        number: 80
```
