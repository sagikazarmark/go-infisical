# Wrapper around the official [Infisical Go SDK](https://github.com/Infisical/go-sdk)

[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/sagikazarmark/go-infisical/ci.yaml?style=flat-square)](https://github.com/sagikazarmark/go-infisical/actions/workflows/ci.yaml)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/mod/github.com/sagikazarmark/go-infisical)
![Go Version](https://img.shields.io/badge/go%20version-%3E=1.22-61CFDD.svg?style=flat-square)
[![built with nix](https://builtwithnix.org/badge.svg)](https://builtwithnix.org)

This is an opinionated (aka. "better") wrapper around the official [Infisical Go SDK](https://github.com/Infisical/go-sdk).

It fixes some of the shortcomings of the original SDK, such as:

- No builtin token refresh
- Mutable client

## Usage

```shell
go get github.com/sagikazarmark/go-infisical
```

## Development

**For an optimal developer experience, it is recommended to install [Nix](https://nixos.org/download.html) and [direnv](https://direnv.net/docs/installation.html).**

TODO

## License

The project is licensed under the [MIT License](LICENSE).
