// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package xdef

import (
	"strings"
	"testing"
	"time"
)

func Test_Created(t *testing.T) {
	t.Run("EnvImgCreated not set", func(t *testing.T) {
		// --- Given ---
		env := []string{EnvImgCreated + "="}

		// --- When ---
		have := Created(env)

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

	t.Run("EnvImgCreated set", func(t *testing.T) {
		// --- Given ---
		env := []string{EnvImgCreated + "=2000-01-02T03:04:05Z"}

		// --- When ---
		have := Created(env)

		// --- Then ---
		if have != "2000-01-02T03:04:05Z" {
			t.Errorf("expected %q got %q", "2000-01-02T03:04:05Z", have)
		}
	})
}

func Test_CreatedStr(t *testing.T) {
	// --- When ---
	have := CreatedStr()

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

func Test_ImgRefName(t *testing.T) {
	t.Run("EnvImgRefName not set", func(t *testing.T) {
		// --- Given ---
		env := []string{EnvImgRefName + "="}

		// --- When ---
		have := ImgRefName(env)

		// --- Then ---
		want := "no-cc-" + time.Now().UTC().Format("060102150405") + "-"
		if !strings.HasPrefix(have, want) {
			t.Errorf("expected \"unknown\" got: %q", have)
		}
	})

	t.Run("EnvImgRefName set", func(t *testing.T) {
		// --- Given ---
		env := []string{EnvImgRefName + "=cc-tag"}

		// --- When ---
		have := ImgRefName(env)

		// --- Then ---
		if have != "cc-tag" {
			t.Errorf(`expected "cc-tag" got: %q`, have)
		}
	})
}

// envLookupTests holds test cases for envLookup.
var envLookupTests = []struct {
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

func Test_EnvLookup_tabular(t *testing.T) {
	for _, tc := range envLookupTests {
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
