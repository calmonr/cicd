# Go CI/CD Playground

This repository is a playground for experimenting with CI/CD in Go.

## Usage (CI/CD)

It's all based on pushes and pull requests. Here is the list of actions triggered based on each event:

- If you push a tag to the repository: `lint`, `release` and `security` actions are triggered.
- If you push a commit to the `main` branch or open a pull request: `build`, `lint` and `security` actions are triggered.

You need to configure the following secrets in the repository settings:

- `DOCKERHUB_USERNAME`: The username of the Docker Hub account.
- `DOCKERHUB_TOKEN`: The token of the Docker Hub account.
- `GITLEAKS_NOTIFY_USER_LIST`: The list of users to notify when a secret leak is found.

## Usage (CLI)

The application is a CLI with a gRPC server (graceful shutdown included) that has a greeting service.

You can run the application using one of the following commands:

```console
go run cmd/playground/main.go
```

```console
go run cmd/playground/main.go --config-file config/sample.yml
```

```console
PLAYGROUND_GRPC_SERVER_ADDRESS=:50051 PLAYGROUND_GRPC_SERVER_NETWORK=tcp go run cmd/playground/main.go
```

## gRPC

[Buf](https://docs.buf.build/introduction) is used to lint the protobuf definitions and generate source code from them.

> You only need to run the following commands if you made changes to the protobuf definitions

```console
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.0
go install github.com/bufbuild/buf/cmd/buf@v1.8.0
```

Before we continue, let's verify that everything is set up properly:

```console
buf build
```

You can run all of the configured lint rules by running this command:

```console
buf lint
```

Now we can generate the source code for our gRPC service:

```console
buf generate -o pkg/proto proto/internal
```

## Actions

- The [build action](.github/workflows/build.yml) builds the binaries and runs the tests.
- The [lint action](.github/workflows/lint.yml) runs `golangci-lint` on the code. ([example](https://github.com/calmonr/cicd/pull/1))
- The [release action](.github/workflows/release.yml) runs `goreleaser` to build and publish the binaries. ([examples](https://github.com/calmonr/cicd/releases))
- The [security action](.github/workflows/security.yml) runs `gitleaks` to check for secret leaks. ([example](https://github.com/calmonr/cicd/pull/2))

## Tools

> You can use [act](https://github.com/nektos/act) to run GitHub Actions locally.

- [x] [GitHub Actions](https://docs.github.com/pt/actions)
- [x] [GoReleaser](https://goreleaser.com/)
- [x] [golangci-lint](https://golangci-lint.run/)
- [x] [EditorConfig](https://editorconfig.org/)
- [x] [Gitleaks](https://gitleaks.io/)
- [x] [Cobra](https://cobra.dev/)
- [x] [Viper](https://github.com/spf13/viper)
- [x] [zap](https://github.com/uber-go/zap)
- [x] [gRPC](https://grpc.io/)
  - [x] [Buf](https://buf.build/)
- [ ] [Dagger](https://dagger.io/)
- [ ] [Task](https://taskfile.dev/)
- [ ] [UPX](https://github.com/upx/upx)

## Plans

- [ ] [CLI completion](https://carlosbecker.com/posts/golang-completions-cobra/)
- [ ] [Versioning](https://goreleaser.com/cookbooks/using-main.version)
- [x] [Docker images](https://goreleaser.com/customization/docker/)
  - [x] GitHub Container Registry
  - [x] Docker Hub
- [ ] Sign [checksums, artifacts](https://goreleaser.com/customization/sign/), [Docker images and manifests](https://goreleaser.com/customization/docker_sign/)
- [x] Tests
- [ ] Improve the documentation

> Probably more to come.

## Tips

You should use [pre-commit](https://pre-commit.com/) to run linting and testing on your project, this way you can catch problems early.

> I'm not going to do that here (we want action :eyes:), but you should.

```yaml
---
repos:
  - repo: https://github.com/golangci/golangci-lint
    rev: v1.48.0
    hooks:
      - id: golangci-lint
  - repo: https://github.com/zricethezav/gitleaks
    rev: v8.12.0
    hooks:
      - id: gitleaks
default_install_hook_types: [pre-commit]
minimum_pre_commit_version: 2.20.0
```
