// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package xdef

// Placeholder values standing in for build metadata a build could not
// determine. Tooling writes them where a value is required but unknown, so a
// label or variable is never left empty.
const (
	// PhDate is the placeholder for an unknown build date. It is the RFC3339
	// zero time, the value the OCI Image Spec prescribes for [LabImgCreated]
	// when the build date is unknown.
	PhDate = "0001-01-01T00:00:00Z"

	// PhHash is the placeholder for an unknown short SCM commit hash.
	PhHash = "0000000"

	// PhTag is the placeholder for an unknown SCM revision tag.
	PhTag = "v0.0.0"

	// PhUnknown is the placeholder for an unknown value of any other kind.
	PhUnknown = "unknown"

	// PhNotSet is the placeholder for a build-metadata variable that ldflags
	// never populated at compile time. Unlike the placeholders above it marks
	// a missing injection, not a value the build could not determine.
	PhNotSet = "<not set>"
)
