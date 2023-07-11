# :partying_face: Template Go

> This is the standard template for Go projects of @katallaxie.

> This is a GitHub Template Repository. You can use the green button to create a new repository based on this template. Read more about [GitHub Template Repositories](https://help.github.com/en/github/creating-cloning-and-archiving-repositories/creating-a-repository-from-a-template).

## Get Started

```bash
make setup MODULE_NAME=github.com/username/repo
```

Features

* [Development Containers](https://containers.dev/)
* [Editorconfig](https://editorconfig.org)
* [GoReleaser](https://goreleaser.com)
* [Hexagonal Architecture](https://en.wikipedia.org/wiki/Hexagonal_architecture_(software))
* [GitHub Actions](https://github.com/features/actions) (lint, test, build, release)

> You can `sh scripts/postCreateCommand.sh` if you are not running in a remote container or on [Codespaces](https://github.com/features/codespaces).

## Usage

This template supports `Makefile` to run tooling.

> `make` is choosen as it is available on most systems.

```bash
# show `help`
make help
```

## Setup

Setup the project.

```bash
make setup MODULE_NAME=github.com/username/repo
```

Other available targets are

* `build`
* `fmt`
* `lint`
* `vet`
* `generate`
* `clean`

The convention is to use `make` to run the build.

Happy coding!
