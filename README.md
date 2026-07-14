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

## Variables

Each value has up to four matching name constants, one per column below:

- **C42 env** (`Env*`) — the runtime environment variable carrying the
  binary's provenance; pairs one-to-one with the ldflags var.
- **OCI env** (`EnvImg*`) — the container environment variable mirroring the
  OCI label.
- **OCI label** (`Lab*`) — the OCI Image Spec label stamped on the image.
- **ldflags var** (`Var*`) — the Go identifier `go build -ldflags -X` targets
  to inject the value at build time; build/SCM metadata only.

The `C42_*` and `OCI_IMAGE_*` variables overlap on several values on purpose:
they describe **two different build events**. `OCI_IMAGE_*` carries the
**image's** provenance, stamped when the image is built; `C42_*` carries the
**binary's** provenance, injected when the binary is built. A binary built from
commit `A` and later packaged into an image tagged `v1.2.4` reports the
binary's values in `C42_*` and the image's in `OCI_IMAGE_*`, and they may
legitimately differ. Read `OCI_IMAGE_*` to answer "which image is this?" and
`C42_*` to answer "which binary is this?".

| Concept            | C42 env          | OCI env                 | OCI label                              | ldflags var |
|--------------------|------------------|-------------------------|----------------------------------------|-------------|
| build date         | `C42_BUILD_DATE` | `OCI_IMAGE_CREATED`     | `org.opencontainers.image.created`     | `buildDate` |
| version tag        | `C42_SCM_REV`    | `OCI_IMAGE_VERSION`     | `org.opencontainers.image.version`     | `scmRev`    |
| commit hash        | `C42_SCM_HASH`   | `OCI_IMAGE_REVISION`    | `org.opencontainers.image.revision`    | `scmHash`   |
| source repository  | `C42_SCM_REPO`   | `OCI_IMAGE_SOURCE`      | `org.opencontainers.image.source`      | `scmRepo`   |
| working-tree state | `C42_SCM_STATE`  | —                       | —                                      | `scmState`  |
| CI/CD identifier   | `C42_CCID`       | `OCI_IMAGE_REF_NAME`    | `org.opencontainers.image.ref.name`    | `ccid`      |
| title / name       | `C42_PROJ_NAME`  | `OCI_IMAGE_TITLE`       | `org.opencontainers.image.title`       | —           |
| authors            | —                | `OCI_IMAGE_AUTHORS`     | `org.opencontainers.image.authors`     | —           |
| description        | —                | `OCI_IMAGE_DESCRIPTION` | `org.opencontainers.image.description` | —           |
| base image         | —                | `OCI_IMAGE_BASE_NAME`   | `org.opencontainers.image.base.name`   | —           |

The project layout variables (`C42_PROJ_ROOT_DIR`, `C42_PROJ_DIST_DIR`,
`C42_PROJ_GO_IMP_SPEC`) have no ldflags or OCI counterpart; project tooling
sets them to point commands at the project.

```go
xdef.VarScmRev // "scmRev"      -> -ldflags -X main.scmRev=v1.2.3
xdef.EnvScmRev // "C42_SCM_REV" -> environment variable read at runtime
```

## Container build variables

Environment variable names read from the project configuration to drive image
builds. Like the project layout variables, they have no ldflags or OCI
counterpart.

| Concept                               | C42 env                |
|---------------------------------------|------------------------|
| registry hostname                     | `C42_REG_HOST`         |
| registry URL scheme                   | `C42_REG_SCHEME`       |
| image repository                      | `C42_REG_REPO`         |
| base image reference (build arg)      | `C42_BLD_IMG_BASE`     |
| build targets (comma-separated)       | `C42_BLD_TARGETS`      |
| in-image root directory               | `C42_CTR_ROOT`         |
| in-image scripts/binaries directory   | `C42_CTR_BIN`          |
| in-image project root directory       | `C42_CTR_PROJECT_ROOT` |
| in-image entrypoint scripts directory | `C42_CTR_ENTRYPOINT`   |
| Go module proxy URL                   | `C42_GOPROXY`          |
| Go checksum database                  | `C42_GOSUMDB`          |
| private Go module patterns            | `C42_GOPRIVATE`        |

The registry (`EnvReg*`) and build (`EnvBld*`) variables drive image
references and the container build; `C42_BLD_IMG_BASE` is passed as a build
argument and `C42_BLD_TARGETS` is a comma-separated list of targets in the build
file (`Containerfile`/`Dockerfile`). The
in-image layout variables (`EnvCtr*`) name the standard directory layout that
ctx42 base images provide. The Go module variables (`EnvGo*`) are copied into the
toolchain's `GOPROXY`, `GOSUMDB`, and `GOPRIVATE`.

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

## License

MIT — see [LICENSE](LICENSE.md).
