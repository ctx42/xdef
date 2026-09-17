// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package xdef

import (
	"testing"
	"time"
)

func Test_BldDate(t *testing.T) {
	t.Run("EnvBldDate empty", func(t *testing.T) {
		// --- Given ---
		env := []string{EnvBldDate + "="}

		// --- When ---
		have := BldDate(env)

		// --- Then ---
		tim, err := time.Parse(time.RFC3339Nano, have)
		if err != nil {
			t.Fatal(err)
		}
		if tim.Location().String() != "UTC" {
			t.Errorf("expected UTC timezone got %s", tim.Location())
		}
		since := time.Since(tim)
		if since > time.Second {
			t.Errorf("expected time diff to be less than 1s got: %s", since)
		}
	})

	t.Run("EnvBldDate set", func(t *testing.T) {
		// --- Given ---
		env := []string{EnvBldDate + "=2000-01-02T03:04:05Z"}

		// --- When ---
		have := BldDate(env)

		// --- Then ---
		if have != "2000-01-02T03:04:05Z" {
			t.Errorf("expected %q got %q", "2000-01-02T03:04:05Z", have)
		}
	})
}

func Test_BldDateStr(t *testing.T) {
	// --- When ---
	have := BldDateStr()

	// --- Then ---
	tim, err := time.Parse(time.RFC3339Nano, have)
	if err != nil {
		t.Fatal(err)
	}
	if tim.Location().String() != "UTC" {
		t.Errorf("expected UTC timezone got %s", tim.Location())
	}
	since := time.Since(tim)
	if since > time.Second {
		t.Errorf("expected time diff to be less than 1s got: %s", since)
	}
}
