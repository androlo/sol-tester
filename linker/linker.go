package linker

import (
	"encoding/hex"
	"fmt"
	"github.com/androlo/sol-tester/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"io/ioutil"
	"math/big"
	"path"
	"regexp"
	"strings"
	"unicode"
)

type Linker struct {
	rootDir   string
	libraries map[string]*LibraryData
	vmenv     vm.Environment
	sender    vm.Account
	gas       *big.Int
}

func NewLinker(rootDir string, vmenv vm.Environment, sender vm.Account, gas *big.Int) *Linker {
	libraries := make(map[string]*LibraryData)
	return &Linker{rootDir, libraries, vmenv, sender, gas}
}

type LibraryData struct {
	Address  common.Address
	Bytecode string
}

func (this *Linker) AddLibrariesFromFile(filePath string) error {
	bts, rErr := ioutil.ReadFile(filePath)
	if rErr != nil {
		return rErr
	}
	libStr := strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, string(bts))
	if len(libStr) == 0 {
		return fmt.Errorf("Libraries file is empty.")
	}
	// Library file parsing in solc splits on comma or whitespace, but
	// in the docs it says format is <libname>: <address> [, or whitespace],
	// which suggests there might be a space after the colon. For this reason
	// we accept only comma separated lists here.
	libs := strings.Split(libStr, ",")
	for idx, lib := range libs {
		tmp := strings.Split(lib, ":")
		if len(tmp) != 2 {
			return fmt.Errorf("Invalid library-file. Data: %s", libStr)
		}
		libName := tmp[0]
		libAddr := tmp[1]

		if libName == "" {
			return fmt.Errorf("Library name missing. Index: %d, Data: %s", idx, libStr)
		}
		if !util.IdentifierRe.MatchString(libName) {
			return fmt.Errorf("Library name not an identifier: %s, Index: %d, Data: %s", libName, idx, libStr)
		}
		if libAddr == "" {
			return fmt.Errorf("Library address missing. Index: %d, Data: %s", idx, libStr)
		}
		if !util.AddressRe.MatchString(libAddr) {
			return fmt.Errorf("Invalid address. Address: %s, Index %d, Data: %s", libAddr, idx, libStr)
		}
		this.libraries[libName] = &LibraryData{common.HexToAddress(libAddr), ""}
	}
	return nil
}

func (this *Linker) Link(bytecode string) (string, error) {
	if !strings.Contains(bytecode, "_") {
		return bytecode, nil
	}
	re, reErr := regexp.Compile(`__([^_]{1,36})__`)
	if reErr != nil {
		return "", reErr
	}
	resultIdx := re.FindStringIndex(bytecode)

	if resultIdx == nil {
		return "", fmt.Errorf("Invalid bytecode format.")
	}
	libName := bytecode[resultIdx[0]+2 : resultIdx[1]-2]
	dErr := this.DeployLibrary(libName)
	if dErr != nil {
		return "", dErr
	}
	length := len(libName)
	padLen := 38 - length
	libLabel := "__" + libName + strings.Repeat("_", padLen)
	addr := this.libraries[libName].Address
	bytecode = strings.Replace(bytecode, libLabel, addr.Hex()[2:], -1)
	// Call recursively until all libraries has been linked.
	return this.Link(bytecode)
}

func (this *Linker) DeployLibrary(libName string) error {
	_, exists := this.libraries[libName]
	if exists {
		return nil
	}
	rErr := this.readLibrary(libName)
	if rErr != nil {
		return rErr
	}
	return this.deployLibrary(libName)
}

func (this *Linker) Libraries() map[string]*LibraryData {
	return this.libraries
}

func (this *Linker) SetRootDir(rootDir string) {
	this.rootDir = rootDir
}

func (this *Linker) RootDir() string {
	return this.rootDir
}

func (this *Linker) Environment() vm.Environment {
	return this.vmenv
}

func (this *Linker) Sender() vm.Account {
	return this.sender
}

func (this *Linker) readLibrary(libName string) error {
	_, exists := this.libraries[libName]
	if exists {
		fmt.Printf("Library %s was already deployed.\n", libName)
		return nil
	}
	bin, bErr := ioutil.ReadFile(path.Join(this.rootDir, libName+".bin"))
	if bErr != nil {
		return bErr
	}
	this.libraries[libName] = &LibraryData{common.Address{}, string(bin)}
	return nil
}

func (this *Linker) deployLibrary(libName string) error {
	lib := this.libraries[libName]
	bytecode, llErr := this.Link(lib.Bytecode)
	if llErr != nil {
		return llErr
	}
	lib.Bytecode = bytecode
	bts, decErr := hex.DecodeString(bytecode)
	if decErr != nil {
		return decErr
	}
	_, addr, cErr := this.vmenv.Create(this.sender, bts, this.gas, big.NewInt(0), big.NewInt(0))

	if cErr != nil {
		return cErr
	}
	lib.Address = addr
	return nil
}
