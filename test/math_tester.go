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

// MathTesterABI is the input ABI used to generate the binding from.
const MathTesterABI = `[{"constant":true,"inputs":[{"name":"a","type":"uint256"},{"name":"b","type":"uint256"}],"name":"addD","outputs":[{"name":"","type":"uint256"}],"type":"function"},{"constant":true,"inputs":[{"name":"a","type":"uint256"},{"name":"b","type":"uint256"}],"name":"add","outputs":[{"name":"","type":"uint256"}],"type":"function"},{"constant":true,"inputs":[{"name":"a","type":"uint256"},{"name":"b","type":"uint256"}],"name":"sub","outputs":[{"name":"","type":"uint256"}],"type":"function"}]`

// MathTesterBin is the compiled bytecode used for deploying new contracts.
const MathTesterBin = `6060604052610281806100126000396000f360606040526000357c01000000000000000000000000000000000000000000000000000000009004806373531b591461004f578063771602f714610084578063b67d77c5146100b95761004d565b005b61006e6004808035906020019091908035906020019091905050610173565b6040518082815260200191505060405180910390f35b6100a360048080359060200190919080359060200190919050506100ee565b6040518082815260200191505060405180910390f35b6100d860048080359060200190919080359060200190919050506101fa565b6040518082815260200191505060405180910390f35b600073__Adder_________________________________63771602f78484604051837c010000000000000000000000000000000000000000000000000000000002815260040180838152602001828152602001925050506020604051808303818660325a03f4156100025750505060405180519060200150905061016d565b92915050565b600073__AdderDelegator________________________63771602f760026006604051837c010000000000000000000000000000000000000000000000000000000002815260040180838152602001828152602001925050506020604051808303818660325a03f415610002575050506040518051906020015090506101f4565b92915050565b600073__Subber________________________________63b67d77c560026006604051837c010000000000000000000000000000000000000000000000000000000002815260040180838152602001828152602001925050506020604051808303818660325a03f4156100025750505060405180519060200150905061027b565b9291505056`

// DeployMathTester deploys a new Ethereum contract, binding an instance of MathTester to it.
func DeployMathTester(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MathTester, error) {
	parsed, err := abi.JSON(strings.NewReader(MathTesterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MathTesterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MathTester{MathTesterCaller: MathTesterCaller{contract: contract}, MathTesterTransactor: MathTesterTransactor{contract: contract}}, nil
}

// MathTester is an auto generated Go binding around an Ethereum contract.
type MathTester struct {
	MathTesterCaller     // Read-only binding to the contract
	MathTesterTransactor // Write-only binding to the contract
}

// MathTesterCaller is an auto generated read-only Go binding around an Ethereum contract.
type MathTesterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathTesterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MathTesterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathTesterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MathTesterSession struct {
	Contract     *MathTester       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathTesterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MathTesterCallerSession struct {
	Contract *MathTesterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MathTesterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MathTesterTransactorSession struct {
	Contract     *MathTesterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// NewMathTester creates a new instance of MathTester, bound to a specific deployed contract.
func NewMathTester(address common.Address, backend bind.ContractBackend) (*MathTester, error) {
	contract, err := bindMathTester(address, backend.(bind.ContractCaller), backend.(bind.ContractTransactor))
	if err != nil {
		return nil, err
	}
	return &MathTester{MathTesterCaller: MathTesterCaller{contract: contract}, MathTesterTransactor: MathTesterTransactor{contract: contract}}, nil
}

// NewMathTesterCaller creates a new read-only instance of MathTester, bound to a specific deployed contract.
func NewMathTesterCaller(address common.Address, caller bind.ContractCaller) (*MathTesterCaller, error) {
	contract, err := bindMathTester(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &MathTesterCaller{contract: contract}, nil
}

// NewMathTesterTransactor creates a new write-only instance of MathTester, bound to a specific deployed contract.
func NewMathTesterTransactor(address common.Address, transactor bind.ContractTransactor) (*MathTesterTransactor, error) {
	contract, err := bindMathTester(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &MathTesterTransactor{contract: contract}, nil
}

// bindMathTester binds a generic wrapper to an already deployed contract.
func bindMathTester(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MathTesterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Add is a free data retrieval call binding the contract method 0x771602f7.
//
// Solidity: function add(a uint256, b uint256) constant returns(uint256)
func (_MathTester *MathTesterCaller) Add(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MathTester.contract.Call(opts, out, "add", a, b)
	return *ret0, err
}

// Add is a free data retrieval call binding the contract method 0x771602f7.
//
// Solidity: function add(a uint256, b uint256) constant returns(uint256)
func (_MathTester *MathTesterSession) Add(a *big.Int, b *big.Int) (*big.Int, error) {
	return _MathTester.Contract.Add(&_MathTester.CallOpts, a, b)
}

// Add is a free data retrieval call binding the contract method 0x771602f7.
//
// Solidity: function add(a uint256, b uint256) constant returns(uint256)
func (_MathTester *MathTesterCallerSession) Add(a *big.Int, b *big.Int) (*big.Int, error) {
	return _MathTester.Contract.Add(&_MathTester.CallOpts, a, b)
}

// AddD is a free data retrieval call binding the contract method 0x73531b59.
//
// Solidity: function addD(a uint256, b uint256) constant returns(uint256)
func (_MathTester *MathTesterCaller) AddD(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MathTester.contract.Call(opts, out, "addD", a, b)
	return *ret0, err
}

// AddD is a free data retrieval call binding the contract method 0x73531b59.
//
// Solidity: function addD(a uint256, b uint256) constant returns(uint256)
func (_MathTester *MathTesterSession) AddD(a *big.Int, b *big.Int) (*big.Int, error) {
	return _MathTester.Contract.AddD(&_MathTester.CallOpts, a, b)
}

// AddD is a free data retrieval call binding the contract method 0x73531b59.
//
// Solidity: function addD(a uint256, b uint256) constant returns(uint256)
func (_MathTester *MathTesterCallerSession) AddD(a *big.Int, b *big.Int) (*big.Int, error) {
	return _MathTester.Contract.AddD(&_MathTester.CallOpts, a, b)
}

// Sub is a free data retrieval call binding the contract method 0xb67d77c5.
//
// Solidity: function sub(a uint256, b uint256) constant returns(uint256)
func (_MathTester *MathTesterCaller) Sub(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MathTester.contract.Call(opts, out, "sub", a, b)
	return *ret0, err
}

// Sub is a free data retrieval call binding the contract method 0xb67d77c5.
//
// Solidity: function sub(a uint256, b uint256) constant returns(uint256)
func (_MathTester *MathTesterSession) Sub(a *big.Int, b *big.Int) (*big.Int, error) {
	return _MathTester.Contract.Sub(&_MathTester.CallOpts, a, b)
}

// Sub is a free data retrieval call binding the contract method 0xb67d77c5.
//
// Solidity: function sub(a uint256, b uint256) constant returns(uint256)
func (_MathTester *MathTesterCallerSession) Sub(a *big.Int, b *big.Int) (*big.Int, error) {
	return _MathTester.Contract.Sub(&_MathTester.CallOpts, a, b)
}
