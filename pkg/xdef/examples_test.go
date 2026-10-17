// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package xdef

import "fmt"

func ExampleBldDate() {
	env := []string{EnvBldDate + "=2000-01-02T03:04:05.678Z"}

	fmt.Println(BldDate(env))
	// Output: 2000-01-02T03:04:05.678Z
}

func ExampleCCRunID() {
	env := []string{
		"GITHUB_ACTIONS=true",
		"GITHUB_REPOSITORY=ctx42/xctr",
		"GITHUB_RUN_ID=1234567890",
		"GITHUB_RUN_ATTEMPT=1",
	}

	fmt.Println(CCRunID(env))
	// Output: github-ctx42/xctr-1234567890-1
}
