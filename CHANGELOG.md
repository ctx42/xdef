## v0.10.1 (Sat, 17 Oct 2026 22:00:08 UTC)
- ci: set C42_CC_RUN_ID in the workflow environment.

## v0.10.0 (Sat, 17 Oct 2026 18:32:56 UTC)
- feat(xdef): add CI/CD run ID variable, label, and reader.

## v0.9.1 (Fri, 09 Oct 2026 12:34:33 UTC)
- docs(xdef): document C42_REG_REPO as host plus repository.

## v0.9.0 (Mon, 21 Sep 2026 11:17:45 UTC)
- ci: add GitHub Actions workflow to run tests.

## v0.8.0 (Mon, 21 Sep 2026 10:58:10 UTC)
- refactor(xdef)!: derive every name from one rule.
- refactor(xdef)!: drop the CI/CD job identifier.
- refactor(xdef): move the build date readers to build.go.
- feat(xdef): add the build target variable.
- test(xdef): stop on an unparsable build date.
- refactor(xdef)!: render the build date at a fixed width.
- docs(xdef): reword the package docs and README in plain English.
- refactor(xdef): make the build date format testable.

## v0.7.0 (Tue, 14 Jul 2026 08:16:43 UTC)
- feat(xdef)!: use container-generic env var names.

## v0.6.0 (Mon, 13 Jul 2026 20:05:57 UTC)
- feat(xdef): add Docker registry and build env vars.

## v0.5.0 (Sun, 12 Jul 2026 12:31:05 UTC)
- feat(xdef): add project layout environment variables.
- chore(idea): rename run config, fix module to xdef.

## v0.4.0 (Sat, 11 Jul 2026 21:54:36 UTC)
- feat(xdef)!: add C42_ build-metadata env vars.

## v0.3.0 (Thu, 09 Jul 2026 14:58:06 UTC)
- feat(xdef): add NotSet placeholder for unpopulated build metadata.

## v0.2.0 (Thu, 09 Jul 2026 14:50:31 UTC)
- docs: improve godoc style.
- feat(xdef): add build-metadata ldflags variable names.

## v0.1.0 (Fri, 12 Jun 2026 10:21:02 UTC)
- feat!: update unique tag prefix in ImgRefName from "no-cc" to "no-ccid".
- docs: clarify description for LabImgRefName and EnvImgRefName annotations.
- docs: refine README for clarity, add label/env variable table.

## v0.0.0 (Thu, 11 Jun 2026 19:08:39 UTC)
- Initial commit.

