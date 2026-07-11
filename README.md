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

## Build Metadata

Build and SCM metadata has two matching name constants per value: the `Var*`
ldflags variable (for `go build -ldflags -X`) and the `Env*` environment
variable that carries the same value at runtime.

| Build variable | Env variable     |
|----------------|------------------|
| `buildDate`    | `C42_BUILD_DATE` |
| `scmRev`       | `C42_SCM_REV`    |
| `scmHash`      | `C42_SCM_HASH`   |
| `scmState`     | `C42_SCM_STATE`  |
| `scmRepo`      | `C42_SCM_REPO`   |
| `ccid`         | `C42_CCID`       |

```go
xdef.VarScmRev // "scmRev"      -> -ldflags -X main.scmRev=v1.2.3
xdef.EnvScmRev // "C42_SCM_REV" -> environment variable read at runtime
```

## Image vs. binary provenance

The `C42_*` (`Env*`) and `OCI_IMAGE_*` (`EnvImg*`) variables overlap on several
values on purpose: they describe **two different build events**.

- `OCI_IMAGE_*` carries the **image's** provenance, stamped when the image is
  built (the OCI Image Spec labels, surfaced to the container as environment
  variables).

- `C42_*` carries the **binary's** provenance, injected when the binary is
  built (`go build -ldflags -X`).

Those events can happen at different times: a binary built from commit `A` and
later packaged into an image tagged `v1.2.4` reports the binary's values in
`C42_*` and the image's values in `OCI_IMAGE_*`, and they may legitimately
differ. Read `OCI_IMAGE_*` to answer "which image is this?" and `C42_*` to
answer "which binary is this?".

| Concept            | OCI env (image build)   | C42 env (binary build) |
|--------------------|-------------------------|------------------------|
| build date         | `OCI_IMAGE_CREATED`     | `C42_BUILD_DATE`       |
| version tag        | `OCI_IMAGE_VERSION`     | `C42_SCM_REV`          |
| commit hash        | `OCI_IMAGE_REVISION`    | `C42_SCM_HASH`         |
| source repository  | `OCI_IMAGE_SOURCE`      | `C42_SCM_REPO`         |
| CI/CD identifier   | `OCI_IMAGE_REF_NAME`    | `C42_CCID`             |
| working-tree state | —                       | `C42_SCM_STATE`        |
| title / name       | `OCI_IMAGE_TITLE`       | —                      |
| authors            | `OCI_IMAGE_AUTHORS`     | —                      |
| description        | `OCI_IMAGE_DESCRIPTION` | —                      |
| base image         | `OCI_IMAGE_BASE_NAME`   | —                      |

## License

MIT — see [LICENSE](LICENSE.md).
