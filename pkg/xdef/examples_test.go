// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package xdef

import "fmt"

func ExampleBldDate() {
	env := []string{EnvBldDate + "=2000-01-02T03:04:05.678Z"}

	fmt.Println(BldDate(env))
	// Output: 2000-01-02T03:04:05.678Z
}

func ExampleCCID() {
	env := []string{EnvBldCCID + "=project-master-29"}

	fmt.Println(CCID(env))
	// Output: project-master-29
}
