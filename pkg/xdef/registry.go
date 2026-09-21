// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package xdef

// Environment variable names describing the private container registry a
// project publishes to. Project tooling reads them from the project
// configuration to assemble image references.
const (
	// EnvRegScheme is the environment variable holding the URL scheme used
	// to reach the private container registry named by [EnvRegHost].
	//
	// Example: https
	EnvRegScheme = "C42_REG_SCHEME"

	// EnvRegHost is the environment variable holding the host of the
	// private container registry. Together with [EnvRegRepo] it marks the
	// remote as configured.
	//
	// Example: my.nexus.dev
	EnvRegHost = "C42_REG_HOST"

	// EnvRegRepo is the environment variable holding the private
	// repository that image references are built from and images are
	// pushed to. It carries the repository alone: the host it lives on
	// is in [EnvRegHost], and a value that repeats the host there builds
	// a reference naming it twice.
	//
	// Example: repo
	EnvRegRepo = "C42_REG_REPO"
)
