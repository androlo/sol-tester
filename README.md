# sol-tester

## NOTICE: This lib is not release ready, and is for personal use as of now.

This library consists of a number of utilities for building and testing Solidity contracts.

### Installing and usage

No dependency management set up. Will add that before real release.

### Prerequisites

Need to have Go installed.

Need go-ethereum on latest develop.

Need `solc` on your path for the builder to work.

#### Installing

`go get -u -t github.com/androlo/sol-tester`

Tested with Go (1.6) on:

- `Ubuntu server 14.04, 64 bit`
- `Windows 10, 64 bit`.

#### Testing

`go test ./...`

### Usage

-

### Tools

#### Builder

The builder is used to compile Solidity-contracts. It features a simple build-file format, and `solc` bindings for calling the compiler from go.

#### Linker

The linker is used to deploy and link Solidity libraries, and can be called on to link the bytecode of a contract before deploying it. It will automatically find libraries that are referenced in the bytecode and deploy them, as well as any libraries they depend on.