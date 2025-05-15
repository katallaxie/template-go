# :partying_face: Template Go

[![Open in GitHub Codespaces](https://github.com/codespaces/badge.svg)](https://codespaces.new/katallaxie/template-go?quickstart=1)

> This is a GitHub Template Repository. You can use the green button to create a new repository based on this template. Read more about [GitHub Template Repositories](https://help.github.com/en/github/creating-cloning-and-archiving-repositories/creating-a-repository-from-a-template).
> Have a look at the [Standard Go Project Layout](https://github.com/golang-standards/project-layout) for setup conventions.

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
