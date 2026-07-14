// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package xdef

// Environment variable names describing the standard in-image directory layout
// that ctx42 base images provide. Programs and entrypoint scripts running
// inside the image read them to locate project files and supporting tooling.
const (
	// EnvCtrRoot is the environment variable holding the absolute path to the
	// root directory inside the image under which files are placed, such as
	// project sources, scripts, and supporting tooling.
	//
	// Example: /ctx42
	EnvCtrRoot = "C42_CTR_ROOT"

	// EnvCtrBin is the environment variable holding the absolute path to the
	// directory with project-related scripts and binary files. It is usually a
	// subdirectory of [EnvCtrRoot].
	//
	// Example: /ctx42/bin
	EnvCtrBin = "C42_CTR_BIN"

	// EnvCtrProjectRoot is the environment variable holding the absolute path
	// inside the image where a project's files are bound or copied. It is
	// usually a subdirectory of [EnvCtrRoot].
	//
	// Example: /ctx42/project
	EnvCtrProjectRoot = "C42_CTR_PROJECT_ROOT"

	// EnvCtrEntrypoint is the environment variable holding the absolute path to
	// the directory with entrypoint scripts run during container startup. It is
	// usually a subdirectory of [EnvCtrRoot].
	//
	// Example: /ctx42/entrypoint
	EnvCtrEntrypoint = "C42_CTR_ENTRYPOINT"
)
