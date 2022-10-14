<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [Golang Template Project](#golang-template-project)
    - [About the project](#about-the-project)
        - [API docs](#api-docs)
        - [Design](#design)
        - [Status](#status)
        - [See also](#see-also)
    - [Getting started](#getting-started)
        - [Layout](#layout)
    - [Notes](#notes)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Golang Template Project

## About the project

Describe what the project is about here.

### API docs

The application is using redoc to generate api documentation.
You can access it by running the app and hitting the url `[HOST]/api/docs`.
You can also download the openapi specification file from `[HOST]/api/docs/openapi.yaml`.

### Design



### Status

This project is a starter one, and is an attempt to create a CRUD app using the hexagonal architecture.

### See also

## Getting started

This project ships with a `makefile` and a `docker-compose` files.

### Prerequisites

- go >= v1.18
- docker and docker-compose

### Install dependencies:
```shell
$ go mod tidy
```

### Run the app:
```shell
$ make run
```

### Build the app:
```shell
$ make build
```

### Run the linter:
```shell
$ make lint
```

### Run the tests (no tests as of now):
```shell
$ make test
```

### Or with `docker-compose`
```shell
$ docker-compose up
```

### Build a command line interface
```shell
$ make build-cli
```

### Layout

```tree
├── Dockerfile
├── README.md
├── cmd
│   ├── cli
│   │   └── main.go
│   ├── graphql
│   │   └── README.md
│   └── http
│       └── main.go
├── docker-compose.yaml
├── docs
│   ├── assets
│   │   ├── index.css
│   │   └── translate-logo.png
│   ├── docs.go
│   ├── index.html
│   └── openapi.yaml
├── go.mod
├── go.sum
├── internal
│   ├── core
│   │   ├── adapters
│   │   │   ├── README.md
│   │   │   └── right
│   │   │       └── repositories
│   │   │           ├── sentencessInMemory.go
│   │   │           └── wordsInMemory.go
│   │   ├── domain
│   │   │   ├── README.md
│   │   │   └── component.go
│   │   ├── dtos
│   │   │   ├── README.md
│   │   │   └── api_dtos.gen.go
│   │   ├── http
│   │   │   └── http_spec.gen.go
│   │   ├── ports
│   │   │   ├── README.md
│   │   │   └── right
│   │   │       └── repositories
│   │   │           ├── sentences.go
│   │   │           └── words.go
│   │   ├── server
│   │   └── usecases
│   │       └── README.md
│   └── infrastructure
│       ├── cli
│       │   └── cobra
│       │       ├── app.go
│       │       ├── commands
│       │       ├── listCommand.go
│       │       └── wordsCommand.go
│       ├── databases
│       │   ├── README.md
│       │   ├── inMemory.go
│       │   └── sqlite.go
│       └── http-server
│           ├── README.md
│           ├── chi
│           ├── echo
│           └── gin
│               └── app.go
└── makefile
```

A brief description of the layout:

* `.github` has two template files for creating PR and issue. Please see the files for more details.
* `.gitignore` varies per project, but all projects need to ignore `bin` directory.
* `.golangci.yml` is the golangci-lint config file.
* `Makefile` is used to build the project. **You need to tweak the variables based on your project**.
* `CHANGELOG.md` contains auto-generated changelog information.
* `OWNERS` contains owners of the project.
* `README.md` is a detailed description of the project.
* `bin` is to hold build outputs.
* `cmd` contains main packages, each subdirecoty of `cmd` is a main package.
* `build` contains scripts, yaml files, dockerfiles, etc, to build and package the project.
* `docs` for project documentations.
* `hack` contains scripts used to manage this repository, e.g. codegen, installation, verification, etc.
* `pkg` places most of project business logic and locate `api` package.
* `release` [chart](https://github.com/caicloud/charts) for production deployment.
* `test` holds all tests (except unit tests), e.g. integration, e2e tests.
* `third_party` for all third party libraries and tools, e.g. swagger ui, protocol buf, etc.
* `vendor` contains all vendored code.

## Notes

* Makefile **MUST NOT** change well-defined command semantics, see Makefile for details.
* Every project **MUST** use `dep` for vendor management and **MUST** checkin `vendor` direcotry.
* `cmd` and `build` **MUST** have the same set of subdirectories for main targets
    * For example, `cmd/admin,cmd/controller` and `build/admin,build/controller`.
    * Dockerfile **MUST** be put under `build` directory even if you have only one Dockerfile.
