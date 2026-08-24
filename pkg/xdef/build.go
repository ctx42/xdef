// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package xdef

// Environment variable names carrying a binary's build provenance. Each pairs
// with the Var* ldflags variable of the same concept, which injects the value
// at compile time.
const (
	// EnvBldDate is the environment variable holding the build date in RFC3339
	// format.
	//
	// Example: 2000-01-02T03:04:05.678Z
	EnvBldDate = "C42_BLD_DATE"

	// EnvBldCCID is the environment variable holding the CI/CD job identifier.
	//
	// Example: project-master-29
	EnvBldCCID = "C42_BLD_CCID"
)
