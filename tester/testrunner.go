package tester

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	// "fmt"
	"github.com/fatih/color"
	"math/big"
	"strings"
	"time"
)

var hiYellow *color.Color

func init() {
	color.NoColor = false
	hiYellow = color.New(color.FgHiYellow)
}

const TEST_EVENT_ID = "0xd204e27263771793b3472e4c07118500eb1ab892c2b2f0a4b90b33c33d4df42f"

const MESSAGE_ID = "0x51a7f65c6325882f237d4aeb43228179cfad48b868511d508e24b4437a819137"

// const RUNTIME_EVENT_ID = ""

// const STACK_EVENT_ID = ""

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

type TestRunner struct {
	envWrapper *EnvWrapper
	testContracts []*TestContract
	testData      *TestData
}

func NewTestRunner(testContracts []*TestContract, testOptions *TestOptions) (*TestRunner, error) {
	ew, newErr := NewEnvWrapper(testOptions)
	if newErr != nil {
		return nil, newErr
	}
	return &TestRunner{
		envWrapper: ew,
		testContracts: testContracts,
		testData:      &TestData{make(map[string]*ContractTestData)},
	}, nil
}

func (self *TestRunner) Run() (*TestData, error) {
	for i, _ := range self.testContracts {
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

func (self *TestRunner) runTest(index int) (ctd *ContractTestData) {
	ctd = &ContractTestData{}
	testContract := self.testContracts[index]
	self.testData.contractTestData[testContract.name] = ctd
	contract, _, createErr := self.envWrapper.DeployContract(testContract.bytecode)
	if createErr != nil {
		ctd.err = createErr
		return
	}

	color.Cyan("\nStarting tests for contract '%s'\n", testContract.name)

	testMethods := testContract.abi.Methods

	methodDataArr := make(map[string]*MethodTestData)

	for mName, method := range testMethods {
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

		pre := time.Now()
		_, _, logs, cErr := contract.Run(method.Id(), nil)
		tDur := time.Since(pre)

		if cErr != nil {
			methodData.success = false
			methodData.errors = append(methodData.errors, cErr)
		} else {
			if len(logs) > 0 {
				for i := 0; i < len(logs); i++ {
					log := logs[i]
					topics := log.Topics
					if topics == nil || len(topics) == 0 {
						break;
					}
					eventId := log.Topics[0].Hex()
					if eventId == MESSAGE_ID {
						n := new(big.Int)
						n.SetBytes(log.Data[32:64])
						if n.Uint64() == 0 {
							hiYellow.Printf("Message received: (empty)\n")
						} else {
							//fmt.Printf("Data-length: %d\n", len(log.Data))
							//fmt.Printf("String length: %s\n", n.String())
							msgLen := n.Uint64()
							msg := string(log.Data[64 : 64+msgLen])
							hiYellow.Printf("Message received: %s\n", msg)
						}
					} else if len(topics) == 2 && TEST_EVENT_ID == eventId{
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
			color.Yellow("Duration: %d (ms)", tDur / 1000000)
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
