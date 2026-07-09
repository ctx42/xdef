// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package xdef

// Build-metadata ldflags variable names. Each value is the exact Go identifier
// that "go build -ldflags -X <pkg>.<name>=<value>" targets. They are the
// variable names populated at build time.
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

	// VarCcid is the variable name holding the CI/CD job identifier.
	//
	// Example: jenkins-skw-tst-master-29
	VarCcid = "ccid"
)
