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

// envJoin returns prefix followed by the values of the "env" variables named
// by keys, joined with "-". It returns an empty string when any of them is
// unset or empty.
func envJoin(env []string, prefix string, keys ...string) string {
	vals := make([]string, 0, len(keys))
	for _, key := range keys {
		val, _ := envLookup(env, key)
		if val == "" {
			return ""
		}
		vals = append(vals, val)
	}
	return prefix + strings.Join(vals, "-")
}
