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

// AdderABI is the input ABI used to generate the binding from.
const AdderABI = `[{"constant":true,"inputs":[{"name":"a","type":"uint256"},{"name":"b","type":"uint256"}],"name":"add","outputs":[{"name":"sum","type":"uint256"}],"type":"function"}]`

// AdderBin is the compiled bytecode used for deploying new contracts.
const AdderBin = `606060405260848060106000396000f3650300e759a2485060606040526000357c010000000000000000000000000000000000000000000000000000000090048063771602f714604157603d565b6007565b605e60048080359060200190919080359060200190919050506074565b6040518082815260200191505060405180910390f35b6000818301905080505b9291505056`

// DeployAdder deploys a new Ethereum contract, binding an instance of Adder to it.
func DeployAdder(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Adder, error) {
	parsed, err := abi.JSON(strings.NewReader(AdderABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AdderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Adder{AdderCaller: AdderCaller{contract: contract}, AdderTransactor: AdderTransactor{contract: contract}}, nil
}

// Adder is an auto generated Go binding around an Ethereum contract.
type Adder struct {
	AdderCaller     // Read-only binding to the contract
	AdderTransactor // Write-only binding to the contract
}

// AdderCaller is an auto generated read-only Go binding around an Ethereum contract.
type AdderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AdderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AdderSession struct {
	Contract     *Adder            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AdderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AdderCallerSession struct {
	Contract *AdderCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AdderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AdderTransactorSession struct {
	Contract     *AdderTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NewAdder creates a new instance of Adder, bound to a specific deployed contract.
func NewAdder(address common.Address, backend bind.ContractBackend) (*Adder, error) {
	contract, err := bindAdder(address, backend.(bind.ContractCaller), backend.(bind.ContractTransactor))
	if err != nil {
		return nil, err
	}
	return &Adder{AdderCaller: AdderCaller{contract: contract}, AdderTransactor: AdderTransactor{contract: contract}}, nil
}

// NewAdderCaller creates a new read-only instance of Adder, bound to a specific deployed contract.
func NewAdderCaller(address common.Address, caller bind.ContractCaller) (*AdderCaller, error) {
	contract, err := bindAdder(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &AdderCaller{contract: contract}, nil
}

// NewAdderTransactor creates a new write-only instance of Adder, bound to a specific deployed contract.
func NewAdderTransactor(address common.Address, transactor bind.ContractTransactor) (*AdderTransactor, error) {
	contract, err := bindAdder(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &AdderTransactor{contract: contract}, nil
}

// bindAdder binds a generic wrapper to an already deployed contract.
func bindAdder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AdderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Add is a free data retrieval call binding the contract method 0x771602f7.
//
// Solidity: function add(a uint256, b uint256) constant returns(sum uint256)
func (_Adder *AdderCaller) Add(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Adder.contract.Call(opts, out, "add", a, b)
	return *ret0, err
}

// Add is a free data retrieval call binding the contract method 0x771602f7.
//
// Solidity: function add(a uint256, b uint256) constant returns(sum uint256)
func (_Adder *AdderSession) Add(a *big.Int, b *big.Int) (*big.Int, error) {
	return _Adder.Contract.Add(&_Adder.CallOpts, a, b)
}

// Add is a free data retrieval call binding the contract method 0x771602f7.
//
// Solidity: function add(a uint256, b uint256) constant returns(sum uint256)
func (_Adder *AdderCallerSession) Add(a *big.Int, b *big.Int) (*big.Int, error) {
	return _Adder.Contract.Add(&_Adder.CallOpts, a, b)
}
