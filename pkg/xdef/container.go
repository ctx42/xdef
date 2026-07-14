// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package xdef

// Environment variable names describing the private container registry a
// project publishes to and the inputs that drive its container build. Project
// tooling reads them from the project configuration to assemble image
// references and run the build.
const (
	// EnvRegHost is the environment variable holding the host of the private
	// container registry. Together with [EnvRegRepo] it marks the remote as
	// configured.
	//
	// Example: my.nexus.dev
	EnvRegHost = "C42_REG_HOST"

	// EnvRegScheme is the environment variable holding the URL scheme used to
	// reach the private container registry named by [EnvRegHost].
	//
	// Example: https
	EnvRegScheme = "C42_REG_SCHEME"

	// EnvRegRepo is the environment variable holding the private repository that
	// image references are built from and images are pushed to.
	//
	// Example: my.nexus.dev/repo
	EnvRegRepo = "C42_REG_REPO"

	// EnvBldImgBase is the environment variable holding the base image
	// reference passed to the container build as a build argument.
	//
	// Example: almalinux:9.5-minimal
	EnvBldImgBase = "C42_BLD_IMG_BASE"

	// EnvBldTargets is the environment variable holding the comma-separated
	// list of build targets. Each target must exist in the build file
	// (Containerfile/Dockerfile).
	//
	// Example: first,second,third
	EnvBldTargets = "C42_BLD_TARGETS"
)
