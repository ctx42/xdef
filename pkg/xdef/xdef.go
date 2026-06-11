// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

// Package xdef provides shared definitions for ctx42 container images:
// OCI Image Spec label names, their matching container environment variable
// names, placeholder values for unknown build metadata, and helpers for
// reading those variables from a process environment slice.
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
)
