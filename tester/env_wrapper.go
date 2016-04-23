package tester

import (
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/common"
	"github.com/androlo/sol-tester/linker"
	"github.com/ethereum/go-ethereum/ethdb"
	"encoding/hex"
	"math/big"
)

type TestOptions struct {
	gas           string
	value         string
	price         string
	debug         bool
	forcejit      bool
	disablejit    bool
	callerAddress string
	rootDir	      string
}

func NewTestOptions(gas, value, price string, debug, forcejit, disablejit bool, callerAddress, rootDir string) *TestOptions {
	return &TestOptions{gas, value, price, debug, forcejit, disablejit, callerAddress, rootDir}
}

type EnvWrapper struct {
	testOptions   *TestOptions
	statedb       *state.StateDB
	sender        vm.Account
	linker        *linker.Linker
}

func NewEnvWrapper(testOptions *TestOptions) (*EnvWrapper, error) {
	db, _ := ethdb.NewMemDatabase()
	statedb, _ := state.New(common.Hash{}, db)
	var callerAddr common.Address
	if testOptions.callerAddress == "" {
		callerAddr = common.StringToAddress("sender")
	} else {
		callerAddr = common.HexToAddress(testOptions.callerAddress)
	}
	sender := statedb.CreateAccount(callerAddr)
	ew := &EnvWrapper{
		testOptions: testOptions,
		statedb: statedb,
		sender: sender,
	}
	ew.linker = linker.NewLinker(testOptions.rootDir, ew.newVMEnv(), sender, common.Big(testOptions.gas))
	return ew, nil
}

func (self *EnvWrapper) newVMEnv() *VMEnv {
	return NewEnv(self.statedb, self.sender.Address(), common.Big(self.testOptions.value), vm.Config{
		Debug:     self.testOptions.debug,
		ForceJit:  self.testOptions.forcejit,
		EnableJit: !self.testOptions.disablejit,
	})
}

func (self *EnvWrapper) DeployContract(code string) (*Contract, uint64, error) {
	vmenv := self.newVMEnv()
	lk := linker.NewLinker(self.testOptions.rootDir, vmenv, self.sender, common.Big(self.testOptions.gas))
	linkedByteCode, linkErr := lk.Link(code)
	if linkErr != nil {
		return nil, 0, linkErr
	}
	bts, decErr := hex.DecodeString(linkedByteCode)
	if decErr != nil {
		return nil, 0, decErr
	}
	gas := common.Big(self.testOptions.gas)
	_, addr, createErr := vmenv.Create(self.sender, bts, gas, common.Big(self.testOptions.price), common.Big(self.testOptions.value))
	if createErr != nil {
		return nil, 0, createErr
	}
	gasTotal := common.Big(self.testOptions.gas).Uint64() - gas.Uint64()
	return &Contract{addr, self}, gasTotal, nil
}

type Contract struct {
	address common.Address
	envWrapper *EnvWrapper
}

func (self *Contract) Run(data []byte, value *big.Int) ([]byte, uint64, vm.Logs, error) {
	vmenv := self.envWrapper.newVMEnv();
	gas := common.Big(self.envWrapper.testOptions.gas)
	price := common.Big(self.envWrapper.testOptions.gas)
	if value == nil {
		value = common.Big(self.envWrapper.testOptions.value)
	}
	logsPreLen := len(vmenv.state.Logs())
	ret, cErr := vmenv.Call(self.envWrapper.sender, self.address, data, gas, price, value)
	gasUsed := common.Big(self.envWrapper.testOptions.gas).Uint64() - gas.Uint64()
	logsPost := vmenv.state.Logs()
	var logsRet vm.Logs
	if len(logsPost) > logsPreLen {
		logsRet = logsPost[logsPreLen: ]
	} else {
		logsRet = vm.Logs{}
	}
	return ret, gasUsed, logsRet, cErr
}