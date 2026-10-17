// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package xdef

import "testing"

func Test_CCRunID_tabular(t *testing.T) {
	jenkins := []string{
		"JENKINS_URL=https://jenkins.example.com/",
		"BUILD_TAG=jenkins-skw-dki-go-master-12",
	}
	github := []string{
		"GITHUB_ACTIONS=true",
		"GITHUB_REPOSITORY=ctx42/xctr",
		"GITHUB_RUN_ID=1234567890",
		"GITHUB_RUN_ATTEMPT=1",
	}
	gitlab := []string{
		"GITLAB_CI=true",
		"CI_PROJECT_PATH=acme/api",
		"CI_JOB_ID=987654",
	}

	tt := []struct {
		testN string

		env  []string
		want string
	}{
		{"explicit", []string{EnvCCRunID + "=my-run"}, "my-run"},
		{
			"explicit wins over a detected system",
			append([]string{EnvCCRunID + "=my-run"}, jenkins...),
			"my-run",
		},
		{
			"empty explicit falls back to detection",
			append([]string{EnvCCRunID + "="}, jenkins...),
			"jenkins-skw-dki-go-master-12",
		},
		{"jenkins", jenkins, "jenkins-skw-dki-go-master-12"},
		{"jenkins without build tag", jenkins[:1], ""},
		{"github", github, "github-ctx42/xctr-1234567890-1"},
		{"github without run attempt", github[:3], ""},
		{
			"github actions not true",
			append([]string{"GITHUB_ACTIONS=false"}, github[1:]...),
			"",
		},
		{"gitlab", gitlab, "gitlab-acme/api-987654"},
		{"gitlab without job ID", gitlab[:2], ""},
		{
			"jenkins checked before github",
			append(github, jenkins...),
			"jenkins-skw-dki-go-master-12",
		},
		{"outside CI", []string{"HOME=/root"}, ""},
		{"empty environment", nil, ""},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have := CCRunID(tc.env)

			// --- Then ---
			if have != tc.want {
				t.Errorf("expected %q got %q", tc.want, have)
			}
		})
	}
}
