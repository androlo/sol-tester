# sol-tester

## NOTICE: This lib is not release ready. It needs a few new features of go-ethereum that involves DELEGATECALL. It will not work fully. It's only uploaded since I use parts of it in other code.

This library consists of a number of utilities for testing Solidity contracts. It makes use of [go-ethereum](https://github.com/ethereum/go-ethereum)'s auto-generated contract proxies and simulated chain.

### Builder

The builder is used to compile Solidity-contracts. It features a simple build-file format, and `solc` bindings for calling the compiler from go.

You must have `solc` on your path for the builder to work.

### Linker

The linker is used to deploy and link Solidity libraries, and can be called on to link the bytecode of a contract before deploying it. It will automatically find libraries that are referenced in the bytecode and deploy them, as well as any libraries they depend on.

When linking, the entire list of libraries has to be provided