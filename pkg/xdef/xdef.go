// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

// Package xdef is the single source of truth for the names ctx42 modules
// share: the environment variables carrying a binary's build provenance and a
// project's container build inputs, the "go build -ldflags -X" variables that
// inject that provenance at compile time, the OCI Image Spec labels stamped on
// built images, and the placeholder values standing in for metadata a build
// could not determine. It also provides helpers for reading those variables
// from a process environment slice.
//
// # Naming
//
// Each kind of name follows one mechanical rule, so a constant and the value
// it carries cannot drift apart:
//
//   - Env* is "C42_" followed by the identifier without its prefix,
//     upper-snake-cased, so [EnvBldImgBase] is "C42_BLD_IMG_BASE". The segment
//     after "C42_" names the family: BLD (build), SCM, PRJ (project), REG
//     (registry), CTR (in-image layout). Each family owns one file in this
//     package, BLD splitting its BLD_IMG container-build inputs into a
//     second.
//   - Var* is the identifier without its prefix, lower-camel-cased, so
//     [VarBldDate] is "bldDate" — the Go identifier "-ldflags -X" targets.
//   - Lab* is an OCI Image Spec annotation key, whose spelling the spec fixes.
//   - Ph* is a placeholder value rather than a name.
//
// The Go module variables are the one exception: [EnvGoProxy], [EnvGoSumDB],
// and [EnvGoPrivate] drop the family separator to read as the toolchain
// variables they are copied into.
//
// This package owns "C42_" and the family segments above. A module building on
// it owns "C42_" followed by its own tool segment, and defines those names
// itself. It defines no vendor-specific label key: the OCI Image Spec fixes
// the spelling of every key here, and a value the spec has no annotation for
// is named by the code stamping it, under that vendor's own prefix.
package xdef
