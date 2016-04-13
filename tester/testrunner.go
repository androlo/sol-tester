package tester

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/ethdb"
	// "fmt"
	"encoding/hex"
	"github.com/androlo/sol-tester/linker"
	"github.com/fatih/color"
	"math/big"
	"strings"
)

func init() {
	color.NoColor = false
}

const TEST_EVENT_ID = "D204E27263771793B3472E4C07118500EB1AB892C2B2F0A4B90B33C33D4DF42F"

// const RUNTIME_EVENT_ID = ""

// const STACK_EVENT_ID = ""

// import "fmt"

type TestContract struct {
	name     string
	bytecode string
	abi      abi.ABI
}

func NewTestContract(name string, bytecode string, abi abi.ABI) *TestContract {
	return &TestContract{name, bytecode, abi}
}

type MethodTestData struct {
	success  bool
	messages []string
	errors   []error
}

type ContractTestData struct {
	methodTestData map[string]*MethodTestData
	err            error
}

type TestData struct {
	contractTestData map[string]*ContractTestData
}

type TestOptions struct {
	gas           string
	value         string
	price         string
	debug         bool
	forcejit      bool
	disablejit    bool
	callerAddress string
}

func NewTestOptions(gas, value, price string, debug, forcejit, disablejit bool, callerAddress string) *TestOptions {
	return &TestOptions{gas, value, price, debug, forcejit, disablejit, callerAddress}
}

type TestRunner struct {
	testContracts []*TestContract
	testOptions   *TestOptions
	dir           string
	testData      *TestData
	statedb       *state.StateDB
	sender        vm.Account
	linker        *linker.Linker
}

func NewTestRunner(testContracts []*TestContract, testOptions *TestOptions, dir string) *TestRunner {
	return &TestRunner{
		testContracts: testContracts,
		testOptions:   testOptions,
		dir:           dir,
		testData:      &TestData{make(map[string]*ContractTestData)},
	}
}

func (self *TestRunner) Run() (*TestData, error) {
	for i, _ := range self.testContracts {
		self.init()
		ctd := self.runTest(i)
		// self.statedb.Commit()
		// color.White("Dump:\n%s\n",string(self.statedb.Dump()))
		if ctd.err != nil {
			color.Red("Failed to deploy contract, skipping. (%s)", ctd.err.Error())
		} else {
			successful := 0
			total := 0
			for _, data := range ctd.methodTestData {
				if data.success {
					successful++
				}
				total++
			}
			color.Magenta("\nSuccessful tests: %d / %d\n", successful, total)
			if successful < total {
				color.Yellow("Some tests failed!\n")
			}
		}
	}

	return self.testData, nil
}

func (self *TestRunner) init() error {
	db, _ := ethdb.NewMemDatabase()
	self.statedb, _ = state.New(common.Hash{}, db)

	var callerAddr common.Address
	if self.testOptions.callerAddress == "" {
		callerAddr = common.StringToAddress("sender")
	} else {
		callerAddr = common.HexToAddress(self.testOptions.callerAddress)
	}

	self.sender = self.statedb.CreateAccount(callerAddr)

	return nil
}

func (self *TestRunner) newVMEnv() *VMEnv {
	return NewEnv(self.statedb, self.sender.Address(), common.Big(self.testOptions.value), vm.Config{
		Debug:     self.testOptions.debug,
		ForceJit:  self.testOptions.forcejit,
		EnableJit: !self.testOptions.disablejit,
	})
}

func (self *TestRunner) runTest(index int) (ctd *ContractTestData) {
	ctd = &ContractTestData{}
	testContract := self.testContracts[index]
	self.testData.contractTestData[testContract.name] = ctd

	vmenv := self.newVMEnv()
	lk := linker.NewLinker(self.dir, vmenv, self.sender, common.Big(self.testOptions.gas))

	linkedByteCode, linkErr := lk.Link(testContract.bytecode)
	if linkErr != nil {
		ctd.err = linkErr
		return
	}
	bts, decErr := hex.DecodeString(linkedByteCode)
	if decErr != nil {
		ctd.err = decErr
		return
	}
	_, addr, createErr := vmenv.Create(self.sender, bts, common.Big(self.testOptions.gas), common.Big(self.testOptions.price), common.Big(self.testOptions.value))
	if createErr != nil {
		ctd.err = createErr
		return
	}

	color.Cyan("\nStarting tests for contract '%s'\n", testContract.name)

	testMethods := testContract.abi.Methods

	methodDataArr := make(map[string]*MethodTestData)

	for mName, method := range testMethods {
		logIndex := len(vmenv.state.Logs())
		methodData := &MethodTestData{true, make([]string, 0), make([]error, 0)}

		// Only run methods that start with 'test'
		if len(mName) < 4 || mName[0:4] != "test" {
			continue
		}
		// Test-methods can have neither input nor output params.
		if len(method.Inputs) != 0 || len(method.Outputs) != 0 {
			continue
		}

		color.White("\nMethod: %s\n", mName)
		// Call the function.
		vmenv = self.newVMEnv()
		_, cErr := vmenv.Call(
			self.sender,
			addr,
			method.Id(),
			common.Big(self.testOptions.gas),
			common.Big(self.testOptions.price),
			common.Big(self.testOptions.value),
		)
		if cErr != nil {
			methodData.success = false
			methodData.errors = append(methodData.errors, cErr)
		} else {
			logs := vmenv.state.Logs()
			if len(logs) > logIndex {
				for i := logIndex; i < len(logs); i++ {

					log := logs[i]
					// fmt.Println(log.String())
					topics := log.Topics
					if len(topics) == 2 {
						eventId := strings.ToUpper(topics[0].Hex()[2:])
						if eventId != TEST_EVENT_ID {
							continue
						}
						res := topics[1].Big().Uint64() == 1
						if !res {
							methodData.success = false
							n := new(big.Int)
							n.SetBytes(log.Data[32:64])
							if n.Uint64() == 0 {
								methodData.messages = append(methodData.messages, "Unknown assertion failure")
							} else {
								//fmt.Printf("Data-length: %d\n", len(log.Data))
								//fmt.Printf("String length: %s\n", n.String())
								msgLen := n.Uint64()
								msg := string(log.Data[64 : 64+msgLen])
								methodData.messages = append(methodData.messages, msg)
							}
						}
					}

				}
			}
		}

		if methodData.success {
			color.Green("SUCCESS\n")
		} else {
			color.Red("FAILED\n")
			for _, err := range methodData.errors {
				color.Red("%s\n", err.Error())
			}
			for _, msg := range methodData.messages {
				color.Red("%s\n", msg)
			}
		}
		methodDataArr[mName] = methodData
	}
	ctd.methodTestData = methodDataArr

	return
}
