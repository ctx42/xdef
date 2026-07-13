// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package xdef

// Environment variable names describing the private Docker registry a project
// publishes to and the Dockerfile targets it builds. Project tooling reads
// them from the project configuration to assemble image references and drive
// "docker build".
const (
	// EnvDkrRegHost is the environment variable holding the Docker registry
	// host of the private repository. Together with [EnvDkrRepo] it marks the
	// remote as configured.
	//
	// Example: my.nexus.dev
	EnvDkrRegHost = "C42_DKR_REG_HOST"

	// EnvDkrRepo is the environment variable holding the private Docker
	// repository the image reference is built from and images are pushed to.
	//
	// Example: my.nexus.dev/repo
	EnvDkrRepo = "C42_DKR_REPO"

	// EnvDkfTargets is the environment variable holding the comma-separated
	// Dockerfile targets to build. Each target must exist in the Dockerfile.
	//
	// Example: first,second,third
	EnvDkfTargets = "C42_DKF_TARGETS"
)
