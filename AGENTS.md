# AGENTS.md

`xdef` (module `github.com/ctx42/xdef`) is the single source of truth for the
names ctx42 modules share: build metadata variables, `-ldflags -X` variables,
OCI image labels, and placeholder values.

## Writing

- Call the values a build stamps into a binary or image "build metadata",
  never "build provenance".
- In Markdown prose write the module name as code (`xdef`); keep headings, the
  title, and commit scopes plain.
