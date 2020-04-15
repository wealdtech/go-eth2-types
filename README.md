# go-eth2-types

[![Tag](https://img.shields.io/github/tag/wealdtech/go-eth2-types.svg)](https://github.com/wealdtech/go-eth2-types/releases/)
[![License](https://img.shields.io/github/license/wealdtech/go-eth2-types.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/wealdtech/go-eth2-types?status.svg)](https://godoc.org/github.com/wealdtech/go-eth2-types)
[![Travis CI](https://img.shields.io/travis/wealdtech/go-eth2-types.svg)](https://travis-ci.org/wealdtech/go-eth2-types)
[![codecov.io](https://img.shields.io/codecov/c/github/wealdtech/go-eth2-types.svg)](https://codecov.io/github/wealdtech/go-eth2-types)

Go library providing Ethereum 2 types.

**Please note that this library uses standards that are not yet final, and as such may result in changes that alter public and private keys.  Do not use this library for production use just yet**

## Table of Contents

- [Install](#install)
- [Usage](#usage)
- [Maintainers](#maintainers)
- [Contribute](#contribute)
- [License](#license)

## Install

`go-eth2-types` is a standard Go module which can be installed with:

```sh
go get github.com/wealdtech/go-eth2-types
```

## Usage

**Before using any cryptographic features you must call `InitBLS()`.**

Please read the [Go documentation for this library](https://godoc.org/github.com/wealdtech/go-eth2-types) for interface information.

## Maintainers

Jim McDonald: [@mcdee](https://github.com/mcdee).

## Contribute

Contributions welcome. Please check out [the issues](https://github.com/wealdtech/go-eth2-types/issues).

## License

[Apache-2.0](LICENSE) © 2019 Weald Technology Trading Ltd
