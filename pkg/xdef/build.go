// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package xdef

import "time"

// Environment variable names carrying a binary's build metadata. Each pairs
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

// bldDateLayout is the layout [BldDateStr] renders the build date with: an
// RFC3339 timestamp whose fractional second is always three digits.
// [time.RFC3339Nano] drops trailing zeros, which makes the width of the
// rendered date vary with the value.
const bldDateLayout = "2006-01-02T15:04:05.000Z07:00"

// BldDateStr returns the current date in UTC formatted as RFC3339 with
// millisecond precision, for example "2000-01-02T03:04:05.678Z". The
// fractional second is always three digits, so every date it returns is the
// same width.
func BldDateStr() string {
	return time.Now().UTC().Format(bldDateLayout)
}
