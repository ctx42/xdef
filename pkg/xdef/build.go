// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package xdef

import "time"

// Environment variable names carrying a binary's build provenance. Each pairs
// with the Var* ldflags variable of the same concept, which injects the value
// at compile time.
const (
	// EnvBldDate is the environment variable holding the build date in RFC3339
	// format.
	//
	// Example: 2000-01-02T03:04:05.678Z
	EnvBldDate = "C42_BLD_DATE"
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
