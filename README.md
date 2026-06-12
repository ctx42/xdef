# xdef

[![Go Reference](https://pkg.go.dev/badge/github.com/ctx42/xdef.svg)](https://pkg.go.dev/github.com/ctx42/xdef)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE.md)

Shared definitions for ctx42 container images: OCI Image Spec label names,
their matching container environment variable names, placeholder values for
unknown build metadata, and helpers for reading those variables from a process
environment slice.

## Import

```go
import "github.com/ctx42/xdef/pkg/xdef"
```

## Usage

Read OCI image metadata from a container's environment, falling back to
sensible defaults when the variables are absent:

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

## License

MIT — see [LICENSE](LICENSE.md).
