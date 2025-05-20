# :partying_face: Template Go

[![Go Reference](https://pkg.go.dev/badge/github.com/katallaxie/template-go.svg)](https://pkg.go.dev/github.com/katallaxie/template-go)
[![go.mod](https://img.shields.io/github/go-mod/go-version/katallaxie/template-go)](go.mod)
[![LICENSE](https://img.shields.io/github/license/katallaxie/template-go)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/katallaxie/template-go)](https://goreportcard.com/report/github.com/katallaxie/template-go)
[![Codecov](https://codecov.io/gh/katallaxie/template-go/branch/main/graph/badge.svg)](https://codecov.io/gh/katallaxie/template-go)


[![Open in GitHub Codespaces](https://github.com/codespaces/badge.svg)](https://codespaces.new/katallaxie/template-go?quickstart=1)

> This is a GitHub Template Repository. You can use the green button to create a new repository based on this template. Read more about [GitHub Template Repositories](https://help.github.com/en/github/creating-cloning-and-archiving-repositories/creating-a-repository-from-a-template).
> Have a look at the [Standard Go Project Layout](https://github.com/golang-standards/project-layout) for setup conventions.

## 🔋 Batteries included

- Continuous integration via [GitHub Actions](https://github.com/features/actions),
- Build automation via [Make](https://www.gnu.org/software/make),
- Dependency management using [Go Modules](https://github.com/golang/go/wiki/Modules),
- Code formatting using [gofumpt](https://github.com/mvdan/gofumpt),
- Linting with [golangci-lint](https://github.com/golangci/golangci-lint)
  and [misspell](https://github.com/client9/misspell),
- Unit testing with [race detector](https://blog.golang.org/race-detector), code coverage [HTML report](https://blog.golang.org/cover) and [Codecov report](https://codecov.io/),
- releasing using [GoReleaser](https://github.com/goreleaser/goreleaser),
- dependencies scanning and updating thanks to [Dependabot](https://dependabot.com),
- security code analysis using [CodeQL Action](https://docs.github.com/en/github/finding-security-vulnerabilities-and-errors-in-your-code/about-code-scanning),
  and [govulncheck](https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck),
- [Visual Studio Code](https://code.visualstudio.com) configuration with [Go](https://code.visualstudio.com/docs/languages/go) support.

## Get Started

This template supports `Makefile` to run tooling.

> `make` is choosen as it is available on most systems.

```bash
# show `help`
make help
```

Other available targets are

```bash
build                          Build the binary file.
clean                          Remove previous build.
fmt                            Run go fmt against code.
generate                       Generate code.
help                           Display this help screen.
lint                           Run lint.
mocks                          Generate mocks.
release                        Release the project.
test                           Run tests.
vet                            Run go vet against code.
```

The convention is to use `make` to run the build.

Happy coding!
