// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package xdef

// Build and SCM metadata variable names for "go build -ldflags -X". Each is
// the exact Go identifier that "-ldflags -X <pkg>.<name>=<value>" targets,
// populated at build time. Each pairs with the Env* variable of the same
// concept below.
const (
	// VarBuildDate is the variable name holding the build date in RFC3339
	// format.
	//
	// Example: 2000-01-02T03:04:05.678Z
	VarBuildDate = "buildDate"

	// VarScmRev is the variable name holding the SCM revision tag.
	//
	// Example: v1.2.3
	VarScmRev = "scmRev"

	// VarScmHash is the variable name holding the short SCM commit hash.
	//
	// Example: 1a12ec31
	VarScmHash = "scmHash"

	// VarScmState is the variable name holding the working-tree state.
	//
	// Example: clean
	VarScmState = "scmState"

	// VarScmRepo is the variable name holding the SCM remote repository
	// URL.
	//
	// Example: https://github.com/ctx42/xdef.git
	VarScmRepo = "scmRepo"

	// VarCCID is the variable name holding the CI/CD job identifier.
	//
	// Example: jenkins-tst-master-29
	VarCCID = "ccid"
)

// Environment variable names carrying the build and SCM metadata. Each pairs
// with the Var* ldflags variable of the same concept above.
const (
	// EnvBuildDate is the environment variable holding the build date in
	// RFC3339 format.
	//
	// Example: 2000-01-02T03:04:05.678Z
	EnvBuildDate = "C42_BUILD_DATE"

	// EnvScmRev is the environment variable holding the SCM revision tag.
	//
	// Example: v1.2.3
	EnvScmRev = "C42_SCM_REV"

	// EnvScmHash is the environment variable holding the short SCM commit
	// hash.
	//
	// Example: 1a12ec31
	EnvScmHash = "C42_SCM_HASH"

	// EnvScmState is the environment variable holding the working-tree
	// state.
	//
	// Example: clean
	EnvScmState = "C42_SCM_STATE"

	// EnvScmRepo is the environment variable holding the SCM remote
	// repository URL.
	//
	// Example: https://github.com/ctx42/xdef.git
	EnvScmRepo = "C42_SCM_REPO"

	// EnvCCID is the environment variable holding the CI/CD job identifier.
	//
	// Example: jenkins-tst-master-29
	EnvCCID = "C42_CCID"
)
