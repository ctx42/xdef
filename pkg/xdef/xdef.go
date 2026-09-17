// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

// Package xdef is the single source of truth for the names ctx42 modules
// share:
//
//   - the environment variables that carry a binary's build metadata and a
//     project's container build inputs,
//   - the "go build -ldflags -X" variables that put that metadata into the
//     binary at compile time,
//   - the OCI Image Spec labels set on built images,
//   - the placeholder values used when a build cannot determine a value.
//
// It also provides helpers that read those variables from a process
// environment slice.
//
// # Naming
//
// Each kind of name follows one simple rule, so a constant and the value it
// holds always match:
//
//   - Env* is "C42_" followed by the identifier without its prefix, in upper
//     snake case, so [EnvBldImgBase] is "C42_BLD_IMG_BASE". The segment after
//     "C42_" names the family: BLD (build), SCM, PRJ (project), REG
//     (registry), CTR (in-image layout). Each family has its own file in this
//     package. BLD is the exception: its BLD_IMG container build inputs live
//     in a second file.
//   - Var* is the identifier without its prefix, in lower camel case, so
//     [VarBldDate] is "bldDate" — the Go variable that "-ldflags -X" sets.
//   - Lab* is an OCI Image Spec annotation key; the spec decides how it is
//     spelled.
//   - Ph* is a placeholder value rather than a name.
//
// The Go module variables are the one exception: [EnvGoProxy], [EnvGoSumDB],
// and [EnvGoPrivate] drop the family separator, so they read like the
// toolchain variables they are copied into.
//
// This package owns "C42_" and the family segments above. A module built on
// it owns "C42_" plus its own tool segment, and defines those names itself.
// It defines no vendor-specific label key: the OCI Image Spec decides the
// spelling of every key here, and a value the spec has no key for is named by
// the code that sets the labels, under that vendor's own prefix.
package xdef
