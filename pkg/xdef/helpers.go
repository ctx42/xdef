// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package xdef

import "strings"

// envLookup retrieves the value of the "env" variable named by the key. If
// the variable is present in the "env", the value (which may be empty) is
// returned and the boolean is true. Otherwise, the returned value will be
// empty, and the boolean will be false. When the key appears more than once,
// the last occurrence wins.
func envLookup(env []string, key string) (string, bool) {
	var exists bool
	var value string
	for _, val := range env {
		if strings.HasPrefix(val, key+"=") {
			value = val[len(key)+1:]
			exists = true
		}
	}
	return value, exists
}
