package linker

import (
    "crypto/rand"
    "fmt"
    "github.com/androlo/sol-tester/util"
    "github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core"
    "github.com/ethereum/go-ethereum/crypto"
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
    backend   *backends.SimulatedBackend
    auth      *bind.TransactOpts
}

func NewLinker(rootDir string) *Linker {
    // Generate a new random account and a funded simulator
    key := crypto.NewKey(rand.Reader)
    backend := backends.NewSimulatedBackend(core.GenesisAccount{key.Address, big.NewInt(10000000000)})
    // Convert the tester key to an authorized transactor for ease of use
    auth := bind.NewKeyedTransactor(key)
    libraries := make(map[string]*LibraryData)
    return &Linker{rootDir, libraries, backend, auth}
}

type LibraryData struct {
    Address  string
    Bytecode string
    Abi      string
    Contract *bind.BoundContract
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
        this.libraries[libName] = &LibraryData{libAddr, "", "", nil}
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
    libName := bytecode[resultIdx[0] + 2 : resultIdx[1] - 2]

    dErr := this.DeployLibrary(libName)
    if dErr != nil {
        return "", dErr
    }
    length := len(libName)
    padLen := 38 - length
    libLabel := "__" + libName + strings.Repeat("_", padLen)
    addr := this.libraries[libName].Address
    if addr[0:2] == "0x" {
        addr = addr[2:]
    }
    bytecode = strings.Replace(bytecode, libLabel, addr, -1)
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

func (this *Linker) Backend() *backends.SimulatedBackend {
    return this.backend
}

func (this *Linker) Auth() *bind.TransactOpts {
    return this.auth
}

func (this *Linker) readLibrary(libName string) error {
    _, exists := this.libraries[libName]
    if exists {
        fmt.Printf("Library %s was already deployed.\n", libName)
        return nil
    }
    bin, bErr := ioutil.ReadFile(path.Join(this.rootDir, libName + ".bin"))
    if bErr != nil {
        return bErr
    }
    abi, aErr := ioutil.ReadFile(path.Join(this.rootDir, libName + ".abi"))
    if aErr != nil {
        return aErr
    }
    this.libraries[libName] = &LibraryData{"", string(bin), string(abi), nil}
    return nil
}

func (this *Linker) deployLibrary(libName string) error {
    lib := this.libraries[libName]
    bytecode, llErr := this.Link(lib.Bytecode)
    if llErr != nil {
        return llErr
    }
    lib.Bytecode = bytecode
    parsed, jsonErr := abi.JSON(strings.NewReader(lib.Abi))
    if jsonErr != nil {
        return jsonErr
    }
    address, _, contract, depErr := bind.DeployContract(this.auth, parsed, common.FromHex(lib.Bytecode), this.backend)
    if depErr != nil {
        return depErr
    }
    lib.Address = address.Hex()
    lib.Contract = contract
    this.backend.Commit()
    return nil
}
