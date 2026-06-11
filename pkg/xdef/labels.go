// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package xdef

// Docker image labels following OCI Image Spec annotation keys.
// See: https://specs.opencontainers.org/image-spec/annotations/
const (
	// LabImgCreated is the OCI Image Spec label for the date and time on which
	// the image was built. The value must be in RFC3339 format. If the build
	// date is unknown, set it to "0001-01-01T00:00:00Z".
	//
	// Example: 2000-01-02T03:04:05.678Z
	LabImgCreated = "org.opencontainers.image.created"

	// LabImgTitle is the OCI Image Spec label for the human-readable name of
	// the application packaged in the image.
	//
	// Example: echo-server
	LabImgTitle = "org.opencontainers.image.title"

	// LabImgDesc is the OCI Image Spec label for a human-readable description
	// of the software packaged in the image.
	//
	// Example: Minimal HTTP echo server.
	LabImgDesc = "org.opencontainers.image.description"

	// LabImgAuthors is the OCI Image Spec label for the contact details of the
	// people or organization responsible for the image.
	LabImgAuthors = "org.opencontainers.image.authors"

	// LabImgSrc is the OCI Image Spec label for the URL to the source code
	// repository used to build the image.
	//
	// Example: https://github.com/ctx42/xdef.git
	LabImgSrc = "org.opencontainers.image.source"

	// LabImgVer is the OCI Image Spec label for the version of the packaged
	// software. It should match the SCM revision tag.
	//
	// Example: v1.2.3
	LabImgVer = "org.opencontainers.image.version"

	// LabImgRev is the OCI Image Spec label for the SCM commit hash from which
	// the image was built.
	//
	// Example: 1a12ec31
	LabImgRev = "org.opencontainers.image.revision"

	// LabImgRefName is the OCI Image Spec annotation used to store the CI/CD
	// build tag that uniquely identifies the pipeline job which built the
	// image. If the image was not built by a CI/CD pipeline, it should be set
	// to "unknown".
	//
	// Example: jenkins-project-master-29
	LabImgRefName = "org.opencontainers.image.ref.name"

	// LabImgBaseName is the OCI Image Spec label for the reference of the base
	// image used to build this image.
	//
	// Example: busybox:1.36.1-uclibc
	LabImgBaseName = "org.opencontainers.image.base.name"
)
