// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

// Package xdef provides shared definitions used across ctx42 software: OCI
// Image Spec label names and their matching container environment variable
// names, build and SCM metadata variable names (both the ldflags variable and
// the environment variable that carry each value), project layout environment
// variable names, placeholder values for unknown build metadata, and helpers
// for reading those variables from a process environment slice.
package xdef

// Placeholders.
const (
	// PhHash defines placeholder unknown hash value.
	PhHash = "0000000"

	// PhRev defines placeholder for unknown revision value.
	PhRev = "v0.0.0"

	// PhUnknown defines placeholder for unknown values.
	PhUnknown = "unknown"

	// PhTime defines placeholder for unknown (zero value) time.
	PhTime = "0001-01-01T00:00:00Z"

	// NotSet is the placeholder for a build-metadata variable that was never
	// populated via ldflags at compile time.
	NotSet = "<not set>"
)
