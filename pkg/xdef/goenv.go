// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package xdef

// Environment variable names carrying Go module configuration for builds run
// inside ctx42 images. Tooling copies them into the matching GOPROXY, GOSUMDB,
// and GOPRIVATE variables of the Go toolchain, which is why they drop the
// family separator every other C42_ name carries.
const (
	// EnvGoProxy is the environment variable holding the Go module proxy
	// URL.
	//
	// Example: https://proxy.golang.org
	EnvGoProxy = "C42_GOPROXY"

	// EnvGoSumDB is the environment variable holding the Go checksum
	// database configuration.
	//
	// Example: sum.golang.org
	EnvGoSumDB = "C42_GOSUMDB"

	// EnvGoPrivate is the environment variable holding the comma-separated
	// list of module path patterns that must not go through the proxy or
	// the checksum database.
	//
	// Example: github.com/ctx42/*
	EnvGoPrivate = "C42_GOPRIVATE"
)
