// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package xdef

// Environment variable names driving a project's container build. Project
// tooling reads them from the project configuration file, or derives them
// from it, and passes them to the build as build arguments.
const (
	// EnvBldImgBase is the environment variable holding the base image
	// reference the build stage of the container build is based on.
	//
	// Example: almalinux:9.5-minimal
	EnvBldImgBase = "C42_BLD_IMG_BASE"

	// EnvBldImgRuntime is the environment variable holding the image
	// reference the final runtime stage of the container build is based on.
	// The build stage uses the image in [EnvBldImgBase].
	//
	// Example: almalinux:9.5-micro
	EnvBldImgRuntime = "C42_BLD_IMG_RUNTIME"

	// EnvBldImgTarget is the environment variable holding the name of the
	// build target the image was built from. It is one of the names listed
	// in [EnvBldImgTargets] and is passed to the build as a build argument.
	// The OCI Image Spec has no annotation for it, so the label stamping
	// it, if any, is named by the code building the image.
	//
	// Example: second
	EnvBldImgTarget = "C42_BLD_IMG_TARGET"

	// EnvBldImgTargets is the environment variable holding the
	// comma-separated list of build targets. Each target must exist in the
	// build file (Containerfile/Dockerfile).
	//
	// It is optional and switches a project from building a single image,
	// with no target selected, to building one image per listed target.
	//
	// Example: first,second,third
	EnvBldImgTargets = "C42_BLD_IMG_TARGETS"
)
