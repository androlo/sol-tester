// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package test

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// AdderTesterABI is the input ABI used to generate the binding from.
const AdderTesterABI = `[{"constant":true,"inputs":[],"name":"add","outputs":[{"name":"","type":"uint256"}],"type":"function"}]`

// AdderTesterBin is the compiled bytecode used for deploying new contracts.
const AdderTesterBin = `606060405260da8060106000396000f360606040526000357c0100000000000000000000000000000000000000000000000000000000900480634f2be91f146037576035565b005b604260048050506058565b6040518082815260200191505060405180910390f35b600073__Adder_________________________________63771602f760026006604051837c010000000000000000000000000000000000000000000000000000000002815260040180838152602001828152602001925050506020604051808303818660325a03f41560025750505060405180519060200150905060d7565b9056`

// DeployAdderTester deploys a new Ethereum contract, binding an instance of AdderTester to it.
func DeployAdderTester(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AdderTester, error) {
	parsed, err := abi.JSON(strings.NewReader(AdderTesterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AdderTesterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AdderTester{AdderTesterCaller: AdderTesterCaller{contract: contract}, AdderTesterTransactor: AdderTesterTransactor{contract: contract}}, nil
}

// AdderTester is an auto generated Go binding around an Ethereum contract.
type AdderTester struct {
	AdderTesterCaller     // Read-only binding to the contract
	AdderTesterTransactor // Write-only binding to the contract
}

// AdderTesterCaller is an auto generated read-only Go binding around an Ethereum contract.
type AdderTesterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdderTesterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AdderTesterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdderTesterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AdderTesterSession struct {
	Contract     *AdderTester      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AdderTesterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AdderTesterCallerSession struct {
	Contract *AdderTesterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// AdderTesterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AdderTesterTransactorSession struct {
	Contract     *AdderTesterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// NewAdderTester creates a new instance of AdderTester, bound to a specific deployed contract.
func NewAdderTester(address common.Address, backend bind.ContractBackend) (*AdderTester, error) {
	contract, err := bindAdderTester(address, backend.(bind.ContractCaller), backend.(bind.ContractTransactor))
	if err != nil {
		return nil, err
	}
	return &AdderTester{AdderTesterCaller: AdderTesterCaller{contract: contract}, AdderTesterTransactor: AdderTesterTransactor{contract: contract}}, nil
}

// NewAdderTesterCaller creates a new read-only instance of AdderTester, bound to a specific deployed contract.
func NewAdderTesterCaller(address common.Address, caller bind.ContractCaller) (*AdderTesterCaller, error) {
	contract, err := bindAdderTester(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &AdderTesterCaller{contract: contract}, nil
}

// NewAdderTesterTransactor creates a new write-only instance of AdderTester, bound to a specific deployed contract.
func NewAdderTesterTransactor(address common.Address, transactor bind.ContractTransactor) (*AdderTesterTransactor, error) {
	contract, err := bindAdderTester(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &AdderTesterTransactor{contract: contract}, nil
}

// bindAdderTester binds a generic wrapper to an already deployed contract.
func bindAdderTester(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AdderTesterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Add is a free data retrieval call binding the contract method 0x4f2be91f.
//
// Solidity: function add() constant returns(uint256)
func (_AdderTester *AdderTesterCaller) Add(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AdderTester.contract.Call(opts, out, "add")
	return *ret0, err
}

// Add is a free data retrieval call binding the contract method 0x4f2be91f.
//
// Solidity: function add() constant returns(uint256)
func (_AdderTester *AdderTesterSession) Add() (*big.Int, error) {
	return _AdderTester.Contract.Add(&_AdderTester.CallOpts)
}

// Add is a free data retrieval call binding the contract method 0x4f2be91f.
//
// Solidity: function add() constant returns(uint256)
func (_AdderTester *AdderTesterCallerSession) Add() (*big.Int, error) {
	return _AdderTester.Contract.Add(&_AdderTester.CallOpts)
}
