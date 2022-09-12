# Go CI/CD Playground

This repository is a playground for experimenting with CI/CD in Go.

## Usage

It's all based on pushes and pull requests. Here is the list of actions triggered based on each event:

- If you push a tag to the repository: `lint`, `release` and `security` actions are triggered.
- If you push a commit to the `main` branch or open a pull request: `build`, `lint` and `security` actions are triggered.

You need to configure the following secrets in the repository settings:

- `DOCKERHUB_USERNAME`: The username of the Docker Hub account.
- `DOCKERHUB_TOKEN`: The token of the Docker Hub account.
- `GITLEAKS_NOTIFY_USER_LIST`: The list of users to notify when a secret leak is found.

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
- [ ] [Cobra](https://cobra.dev/)
- [ ] [gRPC](https://grpc.io/)
  - [ ] [Buf](https://buf.build/)
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
