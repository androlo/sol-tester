package test

import (
	"github.com/androlo/sol-tester/linker"
	"github.com/androlo/sol-tester/tester"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/ethdb"
	"math/big"
	"os"
)

var cwd string

var vmGas = big.NewInt(10000000000)
var vmPrice = big.NewInt(0)
var vmValue = big.NewInt(0)

func init() {
	cwd, _ = os.Getwd()
}

func newTestEnv() (vm.Environment, vm.Account) {
	db, _ := ethdb.NewMemDatabase()
	statedb, _ := state.New(common.Hash{}, db)

	sender := statedb.CreateAccount(common.StringToAddress("sender"))

	return tester.NewEnv(statedb, common.StringToAddress("evmuser"), big.NewInt(0), vm.Config{
		Debug:     false,
		ForceJit:  false,
		EnableJit: false,
	}), sender
}

func newLinker(rootDir string) *linker.Linker {
	vmenv, acc := newTestEnv()
	return linker.NewLinker(rootDir, vmenv, acc, vmGas)
}
