# xdef

[![Go Reference](https://pkg.go.dev/badge/github.com/ctx42/xdef.svg)](https://pkg.go.dev/github.com/ctx42/xdef)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE.md)

Single source of truth for the names ctx42 modules share:

- the environment variables that carry a binary's build metadata and a
  project's container build inputs,
- the `go build -ldflags -X` variables that put that metadata into the binary
  at compile time,
- the OCI Image Spec labels set on built images,
- the labels ctx42 modules must agree on, such as the one tying Docker
  resources to the CI/CD run that created them,
- the placeholder values used when a build cannot determine a value.

`xdef` also provides helpers that read those variables from a process
environment slice, such as `CCRunID`.

## Import

```go
import "github.com/ctx42/xdef/pkg/xdef"
```

## Naming

Each kind of name follows one simple rule, so a constant and the value it
holds always match:

- `Env*` — `C42_` followed by the identifier without its prefix, in upper
  snake case. `EnvBldImgBase` is `C42_BLD_IMG_BASE`.
- `Var*` — the identifier without its prefix, in lower camel case.
  `VarBldDate` is `bldDate`, the Go variable that `-ldflags -X` sets.
- `Lab*` — an OCI Image Spec annotation key, which the spec spells, or a
  ctx42 key: `dev.ctx42.` followed by the family and the identifier without
  its prefix, in lower case and dot-separated. `LabCCRunID` is
  `dev.ctx42.cc.run.id`.
- `Ph*` — a placeholder value rather than a name.

The segment after `C42_` names the **family**: `BLD` (build), `SCM`, `PRJ`
(project), `REG` (registry), `CTR` (in-image layout), `CC` (CI/CD run). Each
family has its own file, which also holds the family's ctx42 label keys.
`BLD` is the exception: its `BLD_IMG` container build inputs live in a second
file. The name itself therefore tells you where it lives:

| File              | Holds                                     |
|-------------------|-------------------------------------------|
| `build.go`        | `C42_BLD_*` build metadata and readers    |
| `scm.go`          | `C42_SCM_*` commit metadata               |
| `ldflags.go`      | `Var*` ldflags variable names             |
| `project.go`      | `C42_PRJ_*` project identity              |
| `registry.go`     | `C42_REG_*` registry coordinates          |
| `image.go`        | `C42_BLD_IMG_*` container build inputs    |
| `container.go`    | `C42_CTR_*` in-image directory layout     |
| `goenv.go`        | `C42_GO*` Go module configuration         |
| `cc.go`           | `C42_CC_*` and `dev.ctx42.cc.*` CI/CD run |
| `labels.go`       | `Lab*` OCI Image Spec labels              |
| `placeholders.go` | `Ph*` placeholder values                  |
| `helpers.go`      | shared environment lookup                 |

The Go module variables are the one exception to the family rule.
`C42_GOPROXY`, `C42_GOSUMDB`, and `C42_GOPRIVATE` drop the separator, so they
read like the toolchain variables they are copied into.

`xdef` owns `C42_` and the family segments above, and `dev.ctx42.` with the
same families. A module built on it owns `C42_` plus its own tool segment, and
defines those names itself.

## Build and SCM metadata

`C42_*` and the OCI labels hold several of the same values on purpose: they
describe **two different build events**. The labels carry the **image's**
metadata, set when the image is built. `C42_*` carries the **binary's**
metadata, set when the binary is built. Take a binary built from commit `A`
and later packed into an image tagged `v1.2.4`: `C42_*` reports the binary's
values, the labels report the image's, and the two may differ for a good
reason. Read the labels to answer "which image is this?", and `C42_*` to
answer "which binary is this?".

| Concept            | C42 env         | ldflags var | OCI label                           |
|--------------------|-----------------|-------------|-------------------------------------|
| build date         | `C42_BLD_DATE`  | `bldDate`   | `org.opencontainers.image.created`  |
| revision tag       | `C42_SCM_REV`   | `scmRev`    | `org.opencontainers.image.version`  |
| commit hash        | `C42_SCM_HASH`  | `scmHash`   | `org.opencontainers.image.revision` |
| working-tree state | `C42_SCM_STATE` | `scmState`  | —                                   |
| source repository  | `C42_SCM_REPO`  | `scmRepo`   | `org.opencontainers.image.source`   |

Watch the spec's wording in the revision-tag and commit-hash rows:
`image.version` holds the revision tag and `image.revision` holds the commit
hash. `LabImgRev` therefore pairs with `C42_SCM_HASH`, not with
`C42_SCM_REV`.

```go
xdef.VarScmRev // "scmRev"      -> -ldflags -X main.scmRev=v1.2.3
xdef.EnvScmRev // "C42_SCM_REV" -> environment variable read at runtime
```

## Image labels

The metadata table above pairs four values with the OCI Image Spec keys that
hold them. The `org.opencontainers` prefix belongs to the spec, and the spec
has no key for some values a build carries — the build target in
`C42_BLD_IMG_TARGET` is one of them. `xdef` names no key for those on purpose.
Such a key belongs to the code that sets the labels, under that vendor's own
reverse domain prefix.

The one exception is a label several ctx42 modules must agree on. `xdef`
defines those under `dev.ctx42.`, in the family's file, so the modules that
set a label and the ones that read it cannot drift apart. Today that is
`LabCCRunID` (see [CI/CD run ID](#cicd-run-id)). No user of this module is
required to set a ctx42 label.

## Project and container build variables

These are read from the project configuration file. They identify the project
and drive its image build. None of them has an ldflags counterpart, and the
spec has no key for any of them (see [Image labels](#image-labels)).

| Concept                               | C42 env               | Family |
|---------------------------------------|-----------------------|--------|
| project name                          | `C42_PRJ_NAME`        | `PRJ`  |
| registry URL scheme                   | `C42_REG_SCHEME`      | `REG`  |
| registry host                         | `C42_REG_HOST`        | `REG`  |
| image repository                      | `C42_REG_REPO`        | `REG`  |
| build-stage base image                | `C42_BLD_IMG_BASE`    | `BLD`  |
| runtime-stage base image              | `C42_BLD_IMG_RUNTIME` | `BLD`  |
| build target                          | `C42_BLD_IMG_TARGET`  | `BLD`  |
| build targets (comma-separated)       | `C42_BLD_IMG_TARGETS` | `BLD`  |
| in-image root directory               | `C42_CTR_ROOT`        | `CTR`  |
| in-image scripts/binaries directory   | `C42_CTR_BIN`         | `CTR`  |
| in-image project root directory       | `C42_CTR_PRJ_ROOT`    | `CTR`  |
| in-image entrypoint scripts directory | `C42_CTR_ENTRYPOINT`  | `CTR`  |
| Go module proxy URL                   | `C42_GOPROXY`         | —      |
| Go checksum database                  | `C42_GOSUMDB`         | —      |
| private Go module patterns            | `C42_GOPRIVATE`       | —      |

The registry variables build image references. `C42_REG_REPO` carries the host
and the repository together (`my.nexus.dev/repo`), the prefix every image
reference starts with. `C42_REG_HOST` and `C42_REG_REPO` together mark the
remote as configured.

The build variables are passed to the container build as build arguments.
`C42_BLD_IMG_TARGETS` is optional: it switches a project from one image to one
image per listed target. Every listed target must exist in the build file
(`Containerfile`/`Dockerfile`). `C42_BLD_IMG_TARGET` names the single target a
given image was built from.

The `C42_CTR_*` variables name the standard directory layout that ctx42 base
images provide, for programs and entrypoint scripts that run inside the image.

## CI/CD run ID

`C42_CC_RUN_ID` identifies one CI/CD run: one Jenkins build, one GitHub
Actions run attempt, one GitLab job. Tools label the Docker resources a run
creates with `dev.ctx42.cc.run.id=<run ID>`, so the run removes its own
leftovers when it ends without any tool knowing which CI/CD system runs it.
The value is opaque, compared only for equality, and must be unique among the
runs that share a Docker daemon.

Tools read it with `CCRunID`, never from the variable directly. It returns
the first match:

| Detected when                | Run ID                                                                                 | Example                          |
|------------------------------|----------------------------------------------------------------------------------------|----------------------------------|
| `C42_CC_RUN_ID` is not empty | its value                                                                              | `my-run`                         |
| `JENKINS_URL` is not empty   | `BUILD_TAG`                                                                            | `jenkins-skw-dki-go-master-12`   |
| `GITHUB_ACTIONS` is `true`   | `github-` + `GITHUB_REPOSITORY`, `GITHUB_RUN_ID`, `GITHUB_RUN_ATTEMPT` joined with `-` | `github-ctx42/xctr-1234567890-1` |
| `GITLAB_CI` is `true`        | `gitlab-` + `CI_PROJECT_PATH`, `CI_JOB_ID` joined with `-`                             | `gitlab-acme/api-987654`         |
| none of these                | empty: no label, nothing removed                                                       |                                  |

A detected system with any of its variables empty yields an empty ID, never a
partial one. The GitHub ID leaves out the job, because a workflow-level `env:`
cannot read it, and GitHub-hosted runners give every job its own Docker
daemon. Runs that share a daemon and get the same derived ID, such as matrix
jobs on a self-hosted runner, set `C42_CC_RUN_ID` themselves. In a GitHub
workflow:

```yaml
env:
  C42_CC_RUN_ID: github-${{ github.repository }}-${{ github.run_id }}-${{ github.run_attempt }}
```

## Placeholders

| Constant    | Value                  | Used for                           |
|-------------|------------------------|------------------------------------|
| `PhDate`    | `0001-01-01T00:00:00Z` | an unknown build date              |
| `PhHash`    | `0000000`              | an unknown commit hash             |
| `PhTag`     | `v0.0.0`               | an unknown revision tag            |
| `PhUnknown` | `unknown`              | an unknown value of any other kind |
| `PhNotSet`  | `<not set>`            | a variable that ldflags never set  |

`PhNotSet` is different from the others: it marks a value that ldflags never
set, not a value the build could not determine.

## Usage

Read build metadata from a process environment. The readers fall back to
sensible defaults when a variable is not set:

```go
env := os.Environ()

// Returns C42_BLD_DATE, or the current UTC time in RFC3339 format with
// millisecond precision if not set.
created := xdef.BldDate(env)

// Returns the CI/CD run ID, or "" outside CI/CD.
runID := xdef.CCRunID(env)
```

Set the labels on an image whose build metadata is unknown:

```go
labels := map[string]string{
    xdef.LabImgCreated: xdef.PhDate,
    xdef.LabImgRev:     xdef.PhHash,
    xdef.LabImgVer:     xdef.PhTag,
}
```

## License

MIT — see [LICENSE](LICENSE.md).
