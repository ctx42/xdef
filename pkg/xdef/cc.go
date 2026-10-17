// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package xdef

// Names tying the Docker resources a CI/CD run creates to that run, so the run
// can remove them when it ends without any tool knowing which CI/CD system
// runs it.
const (
	// EnvCCRunID is the environment variable holding the ID of the CI/CD run
	// the process belongs to: one Jenkins build, one GitHub Actions run
	// attempt, one GitLab job. The value is opaque, compared only for
	// equality, and must be unique among the runs sharing a Docker daemon.
	// Read it with [CCRunID], which derives it when the variable is not set.
	//
	// Example: jenkins-skw-dki-go-master-12
	EnvCCRunID = "C42_CC_RUN_ID"

	// LabCCRunID is the label holding the [CCRunID] of the CI/CD run that
	// created a Docker resource: a container, a built image, a network, or a
	// volume. A run removes the resources labeled with its own ID when it
	// ends.
	//
	// Example: dev.ctx42.cc.run.id=jenkins-skw-dki-go-master-12
	LabCCRunID = "dev.ctx42.cc.run.id"
)

// CCRunID returns the ID of the CI/CD run env belongs to: the value of
// [EnvCCRunID] when it is not empty, otherwise one derived from the variables
// of the CI/CD system it detects, or "" when it detects none, as outside
// CI/CD. It checks in this order, the first match winning:
//
//   - Jenkins, when JENKINS_URL is not empty: BUILD_TAG;
//   - GitHub Actions, when GITHUB_ACTIONS is "true": "github-" followed by
//     GITHUB_REPOSITORY, GITHUB_RUN_ID, and GITHUB_RUN_ATTEMPT joined with "-";
//   - GitLab CI, when GITLAB_CI is "true": "gitlab-" followed by
//     CI_PROJECT_PATH and CI_JOB_ID joined with "-".
//
// A detected system with any of those variables empty yields "", never a
// partial ID. The values are used as they are: a label value may hold any
// string.
//
// The GitHub ID leaves out the job, because a workflow-level environment
// cannot read it; GitHub-hosted runners give every job its own Docker daemon.
// Runs that share a daemon and get the same derived ID, such as matrix jobs on
// a self-hosted runner, set [EnvCCRunID] themselves.
func CCRunID(env []string) string {
	if val, _ := envLookup(env, EnvCCRunID); val != "" {
		return val
	}
	if val, _ := envLookup(env, "JENKINS_URL"); val != "" {
		return envJoin(env, "", "BUILD_TAG")
	}
	if val, _ := envLookup(env, "GITHUB_ACTIONS"); val == "true" {
		return envJoin(
			env,
			"github-",
			"GITHUB_REPOSITORY",
			"GITHUB_RUN_ID",
			"GITHUB_RUN_ATTEMPT",
		)
	}
	if val, _ := envLookup(env, "GITLAB_CI"); val == "true" {
		return envJoin(env, "gitlab-", "CI_PROJECT_PATH", "CI_JOB_ID")
	}
	return ""
}
