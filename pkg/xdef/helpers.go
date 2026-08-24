// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package xdef

import (
	"fmt"
	"strings"
	"time"
)

// BldDate returns the value of the [EnvBldDate] environment variable. Returns
// the current date, formatted the way [BldDateStr] formats it, if the variable
// is not set or is empty.
func BldDate(env []string) string {
	if val, _ := envLookup(env, EnvBldDate); val != "" {
		return val
	}
	return BldDateStr()
}

// BldDateStr returns the current date in UTC formatted as [time.RFC3339Nano],
// truncated to millisecond precision.
func BldDateStr() string {
	return time.Now().UTC().Truncate(time.Millisecond).Format(time.RFC3339Nano)
}

// CCID returns the value of the [EnvBldCCID] environment variable. When the
// environment variable is not set or is empty, it returns a fallback
// identifier formatted like "ccid260913203327674": the "ccid" prefix, the
// current UTC date as twelve digits (YYMMDDhhmmss), and the milliseconds
// within that second as three, zero padded. Both fields are fixed width, so
// they can be read back by position.
//
// The fallback is alphanumeric on purpose. It reaches SemVerBuild as the
// value of the "ccid" pair, where "." separates the pairs and the ctx42
// convention bars "-", which leaves no separator to spend inside the value.
func CCID(env []string) string {
	if val, _ := envLookup(env, EnvBldCCID); val != "" {
		return val
	}
	now := time.Now().UTC()
	msec := now.Nanosecond() / int(time.Millisecond)
	return fmt.Sprintf("ccid%s%03d", now.Format("060102150405"), msec)
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
