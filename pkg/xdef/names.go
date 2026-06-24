// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package xdef

// Environment variables set in Docker images. The OCI_IMAGE_* variables
// mirror the OCI Image Spec labels inside the container. The CTX42_*
// variables are ctx42-specific.
const (
	// EnvImgCreated is the container environment variable mirroring the
	// [LabImgCreated] label. It holds the image build date in RFC3339 format.
	//
	// Example: 2000-01-02T03:04:05.678Z
	EnvImgCreated = "OCI_IMAGE_CREATED"

	// EnvImgTitle is the container environment variable mirroring the
	// [LabImgTitle] label. It holds the human-readable application name.
	//
	// Example: echo-server
	EnvImgTitle = "OCI_IMAGE_TITLE"

	// EnvImgAuthors is the container environment variable mirroring the
	// [LabImgAuthors]. It holds the contact details of the people or
	// organization responsible for the image.
	EnvImgAuthors = "OCI_IMAGE_AUTHORS"

	// EnvImgDesc is the container environment variable mirroring the
	// [LabImgDesc] label. It holds a human-readable description of the
	// software packaged in the image.
	EnvImgDesc = "OCI_IMAGE_DESCRIPTION"

	// EnvImgSrc is the container environment variable mirroring the
	// [LabImgSrc]. It holds the URL to the source code repository used to
	// build the image.
	//
	// Example: https://github.com/ctx42/testkit.git
	EnvImgSrc = "OCI_IMAGE_SOURCE"

	// EnvImgVer is the container environment variable mirroring the
	// [LabImgVer] label. It holds the version of the packaged software,
	// matching the SCM revision tag.
	//
	// Example: v1.2.3
	EnvImgVer = "OCI_IMAGE_VERSION"

	// EnvImgRev is the container environment variable mirroring the
	// [LabImgRev] label. It holds the SCM commit hash from which the image was
	// built.
	//
	// Example: 1a12ec31
	EnvImgRev = "OCI_IMAGE_REVISION"

	// EnvImgRefName is the container environment variable mirroring the
	// [LabImgRefName] annotation (e.g., a registry tag). OCI does not
	// prescribe a specific format. It may be used, for example, as a CI/CD
	// build tag that uniquely identifies the pipeline job which produced the
	// image.
	//
	// Example: jenkins-skw-tst-master-29
	EnvImgRefName = "OCI_IMAGE_REF_NAME"

	// EnvImgBaseName is the container environment variable mirroring the
	// [LabImgBaseName] label.
	//
	// Example: busybox:1.36.1-uclibc
	EnvImgBaseName = "OCI_IMAGE_BASE_NAME"
)
