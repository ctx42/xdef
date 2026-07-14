// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package xdef

// Environment variable names carrying Go module configuration for builds run
// inside ctx42 images. Tooling copies them into the matching GOPROXY, GOSUMDB,
// and GOPRIVATE variables of the Go toolchain.
const (
	// EnvGoProxy is the environment variable holding the Go module proxy URL.
	//
	// Example: https://proxy.golang.org
	EnvGoProxy = "C42_GOPROXY"

	// EnvGoSumDB is the environment variable holding the Go checksum database
	// configuration.
	//
	// Example: sum.golang.org
	EnvGoSumDB = "C42_GOSUMDB"

	// EnvGoPrivate is the environment variable holding the glob patterns
	// matching module path prefixes served from private repositories.
	//
	// Example: github.com/ctx42
	EnvGoPrivate = "C42_GOPRIVATE"
)
