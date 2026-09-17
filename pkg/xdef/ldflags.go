// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package xdef

// Build and SCM metadata variable names for "go build -ldflags -X". Each is
// the exact Go identifier that "-ldflags -X <pkg>.<name>=<value>" targets,
// populated at build time, and pairs with the Env* variable of the same
// concept.
const (
	// VarBldDate is the variable name holding the build date in RFC3339
	// format.
	//
	// Example: 2000-01-02T03:04:05.678Z
	VarBldDate = "bldDate"

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

	// VarScmRepo is the variable name holding the SCM remote repository URL.
	//
	// Example: https://github.com/ctx42/xdef.git
	VarScmRepo = "scmRepo"
)
