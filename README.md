# xdef

[![Go Reference](https://pkg.go.dev/badge/github.com/ctx42/xdef.svg)](https://pkg.go.dev/github.com/ctx42/xdef)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE.md)

Single source of truth for shared definitions used across all ctx42
modules: label names, environment variable names, placeholder values,
and helpers for reading variables from a process environment slice.

## Import

```go
import "github.com/ctx42/xdef/pkg/xdef"
```

## Usage

Read image metadata from a container's environment, falling back to sensible
defaults when the variables are absent:

```go
env := os.Environ()

// Returns OCI_IMAGE_CREATED value, or current UTC time in RFC3339Nano
// format (truncated to millisecond precision) if not set.
created := xdef.Created(env)

// Returns OCI_IMAGE_REF_NAME value, or a unique "no-ccid-<timestamp>"
// tag if not set.
refName := xdef.ImgRefName(env)
```

Placeholder constants are provided for unknown build metadata:

```go
labels := map[string]string{
    xdef.LabImgRev:     xdef.PhHash,
    xdef.LabImgVer:     xdef.PhRev,
    xdef.LabImgCreated: xdef.PhTime,
}
```

## Labels and Environment Variables

| OCI label                              | Env variable            |
|----------------------------------------|-------------------------|
| `org.opencontainers.image.created`     | `OCI_IMAGE_CREATED`     |
| `org.opencontainers.image.title`       | `OCI_IMAGE_TITLE`       |
| `org.opencontainers.image.description` | `OCI_IMAGE_DESCRIPTION` |
| `org.opencontainers.image.authors`     | `OCI_IMAGE_AUTHORS`     |
| `org.opencontainers.image.source`      | `OCI_IMAGE_SOURCE`      |
| `org.opencontainers.image.version`     | `OCI_IMAGE_VERSION`     |
| `org.opencontainers.image.revision`    | `OCI_IMAGE_REVISION`    |
| `org.opencontainers.image.ref.name`    | `OCI_IMAGE_REF_NAME`    |
| `org.opencontainers.image.base.name`   | `OCI_IMAGE_BASE_NAME`   |

## License

MIT — see [LICENSE](LICENSE.md).
