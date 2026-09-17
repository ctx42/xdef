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
