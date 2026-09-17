# xdef

[![Go Reference](https://pkg.go.dev/badge/github.com/ctx42/xdef.svg)](https://pkg.go.dev/github.com/ctx42/xdef)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE.md)

Single source of truth for the names ctx42 modules share: the environment
variables carrying a binary's build provenance and a project's container build
inputs, the `go build -ldflags -X` variables that inject that provenance at
compile time, the OCI Image Spec labels stamped on built images, and the
placeholder values standing in for metadata a build could not determine. It
also provides helpers for reading those variables from a process environment
slice.

## Import

```go
import "github.com/ctx42/xdef/pkg/xdef"
```

## Naming

Each kind of name follows one mechanical rule, so a constant and the value it
carries cannot drift apart:

- `Env*` — `C42_` followed by the identifier without its prefix,
  upper-snake-cased. `EnvBldImgBase` is `C42_BLD_IMG_BASE`.
- `Var*` — the identifier without its prefix, lower-camel-cased. `VarBldDate`
  is `bldDate`, the Go identifier `-ldflags -X` targets.
- `Lab*` — an OCI Image Spec annotation key, whose spelling the spec fixes.
- `Ph*` — a placeholder value rather than a name.

The segment after `C42_` names the **family**: `BLD` (build), `SCM`, `PRJ`
(project), `REG` (registry), `CTR` (in-image layout). Each family owns one
file, `BLD` splitting its `BLD_IMG` container-build inputs into a second, so
where a name lives follows from the name itself:

| File              | Holds                                    |
|-------------------|------------------------------------------|
| `build.go`        | `C42_BLD_*` build provenance and readers |
| `scm.go`          | `C42_SCM_*` commit metadata              |
| `ldflags.go`      | `Var*` ldflags variable names            |
| `project.go`      | `C42_PRJ_*` project identity             |
| `registry.go`     | `C42_REG_*` registry coordinates         |
| `image.go`        | `C42_BLD_IMG_*` container build inputs   |
| `container.go`    | `C42_CTR_*` in-image directory layout    |
| `goenv.go`        | `C42_GO*` Go module configuration        |
| `labels.go`       | `Lab*` OCI Image Spec labels             |
| `placeholders.go` | `Ph*` placeholder values                 |
| `helpers.go`      | shared environment lookup                |

The Go module variables are the one exception to the family rule:
`C42_GOPROXY`, `C42_GOSUMDB`, and `C42_GOPRIVATE` drop the separator to read as
the toolchain variables they are copied into.

xdef owns `C42_` and the family segments above. A module building on it owns
`C42_` followed by its own tool segment, and defines those names itself.

## Build and SCM provenance

`C42_*` and the OCI labels overlap on several values on purpose: they describe
**two different build events**. The labels carry the **image's** provenance,
stamped when the image is built; `C42_*` carries the **binary's** provenance,
injected when the binary is built. A binary built from commit `A` and later
packaged into an image tagged `v1.2.4` reports the binary's values in `C42_*`
and the image's in its labels, and they may legitimately differ. Read the
labels to answer "which image is this?" and `C42_*` to answer "which binary is
this?".

| Concept            | C42 env         | ldflags var | OCI label                           |
|--------------------|-----------------|-------------|-------------------------------------|
| build date         | `C42_BLD_DATE`  | `bldDate`   | `org.opencontainers.image.created`  |
| revision tag       | `C42_SCM_REV`   | `scmRev`    | `org.opencontainers.image.version`  |
| commit hash        | `C42_SCM_HASH`  | `scmHash`   | `org.opencontainers.image.revision` |
| working-tree state | `C42_SCM_STATE` | `scmState`  | —                                   |
| source repository  | `C42_SCM_REPO`  | `scmRepo`   | `org.opencontainers.image.source`   |

Note the spec's wording for the revision-tag and commit-hash rows:
`image.version` carries the revision tag and `image.revision` carries the
commit hash, so `LabImgRev` pairs with `C42_SCM_HASH`, not with
`C42_SCM_REV`.

```go
xdef.VarScmRev // "scmRev"      -> -ldflags -X main.scmRev=v1.2.3
xdef.EnvScmRev // "C42_SCM_REV" -> environment variable read at runtime
```

## Project and container build variables

Read from the project configuration file to identify the project and drive its
image build. None of them has an ldflags or OCI counterpart.

| Concept                               | C42 env               | Family |
|---------------------------------------|-----------------------|--------|
| project name                          | `C42_PRJ_NAME`        | `PRJ`  |
| registry URL scheme                   | `C42_REG_SCHEME`      | `REG`  |
| registry host                         | `C42_REG_HOST`        | `REG`  |
| image repository                      | `C42_REG_REPO`        | `REG`  |
| build-stage base image                | `C42_BLD_IMG_BASE`    | `BLD`  |
| runtime-stage base image              | `C42_BLD_IMG_RUNTIME` | `BLD`  |
| build targets (comma-separated)       | `C42_BLD_IMG_TARGETS` | `BLD`  |
| in-image root directory               | `C42_CTR_ROOT`        | `CTR`  |
| in-image scripts/binaries directory   | `C42_CTR_BIN`         | `CTR`  |
| in-image project root directory       | `C42_CTR_PRJ_ROOT`    | `CTR`  |
| in-image entrypoint scripts directory | `C42_CTR_ENTRYPOINT`  | `CTR`  |
| Go module proxy URL                   | `C42_GOPROXY`         | —      |
| Go checksum database                  | `C42_GOSUMDB`         | —      |
| private Go module patterns            | `C42_GOPRIVATE`       | —      |

The registry variables assemble image references; `C42_REG_HOST` and
`C42_REG_REPO` together mark the remote as configured. The build variables are
passed to the container build as build arguments — `C42_BLD_IMG_TARGETS` is
optional and switches a project from one image to one image per listed target,
each of which must exist in the build file (`Containerfile`/`Dockerfile`). The
`C42_CTR_*` variables name the standard directory layout ctx42 base images
provide, for programs and entrypoint scripts running inside the image.

## Placeholders

| Constant    | Value                  | Stands in for                      |
|-------------|------------------------|------------------------------------|
| `PhDate`    | `0001-01-01T00:00:00Z` | an unknown build date              |
| `PhHash`    | `0000000`              | an unknown commit hash             |
| `PhTag`     | `v0.0.0`               | an unknown revision tag            |
| `PhUnknown` | `unknown`              | an unknown value of any other kind |
| `PhNotSet`  | `<not set>`            | a variable ldflags never populated |

`PhNotSet` is the odd one out: it marks a missing ldflags injection rather than
a value the build could not determine.

## Usage

Read build metadata from a process environment, falling back to sensible
defaults when the variables are absent:

```go
env := os.Environ()

// Returns C42_BLD_DATE, or the current UTC time in RFC3339Nano format
// (truncated to millisecond precision) if not set.
created := xdef.BldDate(env)
```

Stamp an image whose build metadata is unknown:

```go
labels := map[string]string{
    xdef.LabImgCreated: xdef.PhDate,
    xdef.LabImgRev:     xdef.PhHash,
    xdef.LabImgVer:     xdef.PhTag,
}
```

## License

MIT — see [LICENSE](LICENSE.md).
