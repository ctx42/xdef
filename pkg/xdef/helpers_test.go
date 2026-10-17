// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package xdef

import "testing"

func Test_envLookup_tabular(t *testing.T) {
	tt := []struct {
		testN string

		env        []string
		findKey    string
		wantValue  string
		wantExists bool
	}{
		{
			"found",
			[]string{"key0=val0", "key1=val1"},
			"key1",
			"val1",
			true,
		},
		{
			"not found",
			[]string{"key0=val0", "key1=val1"},
			"key9",
			"",
			false,
		},
		{
			"partial",
			[]string{"key0=val0", "key1=val1"},
			"key",
			"",
			false,
		},
		{"empty env", []string{}, "key", "", false},
		{
			"empty key",
			[]string{"key0=val0", "key1=val1"},
			"",
			"",
			false,
		},
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
			format := "expected value %#q got %#q"
			if tc.wantValue != haveValue {
				t.Errorf(format, tc.wantValue, haveValue)
			}

			format = "expected exists %v got %v"
			if tc.wantExists != haveExists {
				t.Errorf(format, tc.wantExists, haveExists)
			}
		})
	}
}

func Test_envJoin_tabular(t *testing.T) {
	env := []string{"A=a", "B=b", "EMPTY="}

	tt := []struct {
		testN string

		prefix string
		keys   []string
		want   string
	}{
		{"one key", "", []string{"A"}, "a"},
		{"keys joined after prefix", "p-", []string{"A", "B"}, "p-a-b"},
		{"unset key", "p-", []string{"A", "MISSING"}, ""},
		{"empty key", "p-", []string{"A", "EMPTY"}, ""},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have := envJoin(env, tc.prefix, tc.keys...)

			// --- Then ---
			if have != tc.want {
				t.Errorf("expected %q got %q", tc.want, have)
			}
		})
	}
}
