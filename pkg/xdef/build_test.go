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
		format := "expected time diff to be less than 1s got: %s"
		if since > time.Second {
			t.Errorf(format, since)
		}
	})

	t.Run("EnvBldDate set", func(t *testing.T) {
		// --- Given ---
		env := []string{EnvBldDate + "=2000-01-02T03:04:05Z"}

		// --- When ---
		have := BldDate(env)

		// --- Then ---
		want := "2000-01-02T03:04:05Z"
		if have != want {
			t.Errorf("expected %q got %q", want, have)
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

func Test_bldDateStr_tabular(t *testing.T) {
	tt := []struct {
		testN string

		tim  time.Time
		want string
	}{
		{
			"millisecond",
			time.Date(2000, 1, 2, 3, 4, 5, 678_000_000, time.UTC),
			"2000-01-02T03:04:05.678Z",
		},
		{
			"millisecond ending in zero",
			time.Date(2000, 1, 2, 3, 4, 5, 670_000_000, time.UTC),
			"2000-01-02T03:04:05.670Z",
		},
		{
			"whole second",
			time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC),
			"2000-01-02T03:04:05.000Z",
		},
		{
			"below millisecond precision",
			time.Date(2000, 1, 2, 3, 4, 5, 678_999_999, time.UTC),
			"2000-01-02T03:04:05.678Z",
		},
		{
			"non-UTC zone",
			time.Date(
				2000, 1, 2, 3, 4, 5, 678_000_000,
				time.FixedZone("CET", 2*60*60),
			),
			"2000-01-02T01:04:05.678Z",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have := bldDateStr(tc.tim)

			// --- Then ---
			if have != tc.want {
				t.Errorf("expected %q got %q", tc.want, have)
			}
		})
	}
}
