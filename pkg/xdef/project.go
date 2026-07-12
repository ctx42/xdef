// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package xdef

// Environment variable names describing a project's name, on-disk layout, and
// Go identity. Unlike the build and SCM metadata, these have no ldflags
// counterpart; project tooling sets them to point commands at the project.
const (
	// EnvProjName is the environment variable holding the project name. It
	// carries the same information as [EnvImgTitle], the human-readable
	// application name, but identifies the project on disk rather than inside a
	// container image.
	//
	// Example: echo-server
	EnvProjName = "C42_PROJ_NAME"

	// EnvProjRootDir is the environment variable holding the absolute path to
	// the project root directory.
	//
	// Example: /home/user/projects/echo-server
	EnvProjRootDir = "C42_PROJ_ROOT_DIR"

	// EnvProjDistDir is the environment variable holding the absolute path to
	// the project distribution directory.
	//
	// Example: /home/user/projects/echo-server/dist
	EnvProjDistDir = "C42_PROJ_DIST_DIR"

	// EnvProjGoImpSpec is the environment variable holding the Go project
	// import spec.
	//
	// Example: github.com/ctx42/echo-server
	EnvProjGoImpSpec = "C42_PROJ_GO_IMP_SPEC"
)
