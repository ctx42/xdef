// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package xdef

// Container image labels following OCI Image Spec annotation keys. They carry
// the image's metadata, stamped when the image is built, where the Env*
// variables carry the binary's.
// See: https://specs.opencontainers.org/image-spec/annotations/
const (
	// LabImgCreated is the OCI Image Spec label for the date and time on which
	// the image was built. The value must be in RFC3339 format. If the build
	// date is unknown, set it to [PhDate].
	//
	// Example: 2000-01-02T03:04:05.678Z
	LabImgCreated = "org.opencontainers.image.created"

	// LabImgRev is the OCI Image Spec label for the SCM commit hash from which
	// the image was built. Despite the spec's "revision" wording it takes the
	// hash in [EnvScmHash], not the revision tag in [EnvScmRev].
	//
	// Example: 1a12ec31
	LabImgRev = "org.opencontainers.image.revision"

	// LabImgVer is the OCI Image Spec label for the version of the packaged
	// software. It should match the SCM revision tag in [EnvScmRev].
	//
	// Example: v1.2.3
	LabImgVer = "org.opencontainers.image.version"

	// LabImgSrc is the OCI Image Spec label for the URL to the source code
	// repository used to build the image.
	//
	// Example: https://github.com/ctx42/xdef.git
	LabImgSrc = "org.opencontainers.image.source"
)
