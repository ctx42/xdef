// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package xdef

// Environment variable names carrying the SCM metadata of the commit a binary
// was built from. Each pairs with the Var* ldflags variable of the same
// concept.
const (
	// EnvScmRev is the environment variable holding the SCM revision tag. It
	// carries the version tag, the value [LabImgVer] takes; the commit hash is
	// in [EnvScmHash].
	//
	// Example: v1.2.3
	EnvScmRev = "C42_SCM_REV"

	// EnvScmHash is the environment variable holding the short SCM commit
	// hash.
	//
	// Example: 1a12ec31
	EnvScmHash = "C42_SCM_HASH"

	// EnvScmState is the environment variable holding the working-tree state.
	//
	// Example: clean
	EnvScmState = "C42_SCM_STATE"

	// EnvScmRepo is the environment variable holding the SCM remote repository
	// URL.
	//
	// Example: https://github.com/ctx42/xdef.git
	EnvScmRepo = "C42_SCM_REPO"
)
