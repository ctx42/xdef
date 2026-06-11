// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package xdef

import (
	"strings"
	"time"
)

// Created returns the value of the [EnvImgCreated] environment variable.
// Returns the current UTC date in RFC3339 format if the variable is not set or
// is empty.
func Created(env []string) string {
	if val, _ := envLookup(env, EnvImgCreated); val != "" {
		return val
	}
	return CreatedStr()
}

// CreatedStr returns the current date in UTC formatted as [time.RFC3339Nano],
// truncated to millisecond precision.
func CreatedStr() string {
	return time.Now().UTC().Truncate(time.Millisecond).Format(time.RFC3339Nano)
}

// ImgRefName returns the value of [EnvImgRefName] environment variable. When
// the environment variable is not set, it will return a string formatted like
// "no-cc-240808100435-885410365" where the first set of numbers is the current
// date and the second is the number of nanoseconds.
func ImgRefName(env []string) string {
	if val, _ := envLookup(env, EnvImgRefName); val != "" {
		return val
	}
	tim := time.Now().UTC().Format("060102150405-.999999999")
	tag := "no-cc-" + strings.ReplaceAll(tim, ".", "")
	return tag
}

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
