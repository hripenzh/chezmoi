# Contributing to chezmoi

Thank you for your interest in contributing to chezmoi!

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) 1.21 or later
- [Git](https://git-scm.com/)

### Building from source

```console
$ git clone https://github.com/chezmoi/chezmoi.git
$ cd chezmoi
$ go build ./cmd/chezmoi
```

### Running tests

```console
$ go test ./...
```

## Making Changes

1. Fork the repository.
2. Create a new branch for your change: `git checkout -b my-feature`
3. Make your changes and add tests where appropriate.
4. Run `go test ./...` to ensure all tests pass.
5. Run `golangci-lint run` to check for lint issues.
6. Commit your changes following [Conventional Commits](https://www.conventionalcommits.org/).
7. Push your branch and open a pull request.

## Code Style

- Follow standard Go conventions (`gofmt`, `goimports`).
- Keep functions small and focused.
- Add comments for exported types and functions.
- Write table-driven tests where possible.

## Reporting Issues

Please use the [GitHub issue tracker](https://github.com/chezmoi/chezmoi/issues) to report bugs or request features. Include:

- chezmoi version (`chezmoi --version`)
- Operating system and version
- Steps to reproduce the issue
- Expected and actual behavior

## Code of Conduct

This project follows the [Contributor Covenant Code of Conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code.
