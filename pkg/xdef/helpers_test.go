// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package xdef

import (
	"strings"
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
			t.Error(err)
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
		t.Error(err)
	}
	if tim.Location().String() != "UTC" {
		t.Errorf("expected UTC timezone got %s", tim.Location())
	}
	since := time.Since(tim)
	if since > time.Second {
		t.Errorf("expected time diff to be less than 1s got: %s", since)
	}
}

func Test_CCID(t *testing.T) {
	t.Run("EnvBldCCID empty", func(t *testing.T) {
		// --- Given ---
		env := []string{EnvBldCCID + "="}
		before := time.Now().UTC().Format("060102150405")

		// --- When ---
		have := CCID(env)

		// --- Then ---
		after := time.Now().UTC().Format("060102150405")
		wBefore, wAfter := "ccid"+before, "ccid"+after
		if !strings.HasPrefix(have, wBefore) &&
			!strings.HasPrefix(have, wAfter) {

			t.Errorf("expected prefix %q or %q got: %q",
				wBefore, wAfter, have)
		}
		if len(have) != len(wBefore)+3 {
			t.Errorf("expected length %d got: %d (%q)",
				len(wBefore)+3, len(have), have)
		}
		if strings.ContainsAny(have, ".-") {
			t.Errorf(`expected no "." or "-" in the fallback got: %q`, have)
		}
	})

	t.Run("EnvBldCCID set", func(t *testing.T) {
		// --- Given ---
		env := []string{EnvBldCCID + "=project-master-29"}

		// --- When ---
		have := CCID(env)

		// --- Then ---
		if have != "project-master-29" {
			t.Errorf(`expected "project-master-29" got: %q`, have)
		}
	})
}

func Test_envLookup_tabular(t *testing.T) {
	tt := []struct {
		testN string

		env        []string
		findKey    string
		wantValue  string
		wantExists bool
	}{
		{"found", []string{"key0=val0", "key1=val1"}, "key1", "val1", true},
		{"not found", []string{"key0=val0", "key1=val1"}, "key9", "", false},
		{"partial", []string{"key0=val0", "key1=val1"}, "key", "", false},
		{"empty env", []string{}, "key", "", false},
		{"empty key", []string{"key0=val0", "key1=val1"}, "", "", false},
		{
			"last value counts",
			[]string{"key0=val0", "key1=val1", "key0=abc"},
			"key0",
			"abc",
			true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			haveValue, haveExists := envLookup(tc.env, tc.findKey)

			// --- Then ---
			if tc.wantValue != haveValue {
				t.Errorf("expected value %#q got %#q", tc.wantValue, haveValue)
			}
			if tc.wantExists != haveExists {
				t.Errorf(
					"expected value `%v` got `%v`",
					tc.wantExists,
					haveExists,
				)
			}
		})
	}
}
