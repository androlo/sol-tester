package test

import (
    "github.com/androlo/sol-tester/linker"
    "github.com/stretchr/testify/assert"
    "path"
    "testing"
    "github.com/androlo/sol-tester/util"
    "strings"
    //"github.com/ethereum/go-ethereum/accounts/abi"
    //"github.com/ethereum/go-ethereum/common"
    //"github.com/ethereum/go-ethereum/accounts/abi/bind"
    "math/big"
    //"fmt"
)

const AdderAbi = `[{"constant":true,"inputs":[{"name":"a","type":"uint256"},{"name":"b","type":"uint256"}],"name":"add","outputs":[{"name":"sum","type":"uint256"}],"type":"function"}]`

const NoLibs = "6060604052600a8060106000396000f360606040526008565b00"

const Add = "606060405260da8060106000396000f360606040526000357c0100000000000000000000000000000000000000000000000000000000900480634f2be91f146037576035565b005b604260048050506058565b6040518082815260200191505060405180910390f35b600073__Adder_________________________________63771602f760026006604051837c010000000000000000000000000000000000000000000000000000000002815260040180838152602001828152602001925050506020604051808303818660325a03f41560025750505060405180519060200150905060d7565b9056"

const All = "6060604052610244806100126000396000f360606040526000357c0100000000000000000000000000000000000000000000000000000000900480634f2be91f1461004f578063718353bf14610072578063c54124be146100955761004d565b005b61005c60048050506100b8565b6040518082815260200191505060405180910390f35b61007f600480505061013c565b6040518082815260200191505060405180910390f35b6100a260048050506101c0565b6040518082815260200191505060405180910390f35b600073__Adder_________________________________63771602f760026006604051837c010000000000000000000000000000000000000000000000000000000002815260040180838152602001828152602001925050506020604051808303818660325a03f41561000257505050604051805190602001509050610139565b90565b600073__AdderDelegator________________________63771602f760026006604051837c010000000000000000000000000000000000000000000000000000000002815260040180838152602001828152602001925050506020604051808303818660325a03f415610002575050506040518051906020015090506101bd565b90565b600073__Subber________________________________63b67d77c560026006604051837c010000000000000000000000000000000000000000000000000000000002815260040180838152602001828152602001925050506020604051808303818660325a03f41561000257505050604051805190602001509050610241565b9056"

const AdderTag = "__Adder_________________________________"

const SubberTag = "__Subber________________________________"

const AdderDelegatorTag = "__AdderDelegator________________________"

func TestLinkNoLibs(t *testing.T) {
    moduleDir := path.Join(cwd, "lib_math")
    lk := linker.NewLinker(moduleDir)
    bytecode, lErr := lk.Link(NoLibs)

    assert.NoError(t, lErr)
    assert.Equal(t, bytecode, NoLibs)
}

func TestDeployLibraryNoLinking(t *testing.T) {
    moduleDir := path.Join(cwd, "lib_math")

    lk := linker.NewLinker(moduleDir)
    dErr := lk.DeployLibrary("Adder")

    assert.NoError(t, dErr)
    lib := lk.Libraries()["Adder"]
    assert.True(t, util.AddressRe.MatchString(lib.Address))
    assert.Equal(t, lib.Bytecode, AdderBin)
    assert.Equal(t, lib.Abi, AdderAbi)
}

func TestDeployLibraryLinking(t *testing.T) {
    moduleDir := path.Join(cwd, "lib_math")

    lk := linker.NewLinker(moduleDir)
    dErr := lk.DeployLibrary("AdderDelegator")

    assert.NoError(t, dErr)
    lib := lk.Libraries()["AdderDelegator"]
    assert.True(t, util.AddressRe.MatchString(lib.Address))
    lib = lk.Libraries()["Adder"]
    assert.True(t, util.AddressRe.MatchString(lib.Address))
}

func TestLinkOne(t *testing.T) {
    moduleDir := path.Join(cwd, "lib_math")

    lk := linker.NewLinker(moduleDir)
    bytecode, dErr := lk.Link(Add)

    assert.NoError(t, dErr)
    lib := lk.Libraries()["Adder"]
    assert.True(t, util.AddressRe.MatchString(lib.Address))
    addr := lib.Address[2:]
    rep := strings.Replace(Add, AdderTag, addr, -1)
    assert.Equal(t, rep, bytecode)
}

func TestLinkSeveral(t *testing.T) {
    moduleDir := path.Join(cwd, "lib_math")

    lk := linker.NewLinker(moduleDir)
    bytecode, dErr := lk.Link(All)

    assert.NoError(t, dErr)
    lib := lk.Libraries()["Adder"]
    adderAddr := lib.Address[2:]
    rep := strings.Replace(All, AdderTag, adderAddr, -1)
    lib = lk.Libraries()["Subber"]
    subberAddr := lib.Address[2:]
    rep = strings.Replace(rep, SubberTag, subberAddr, -1)
    lib = lk.Libraries()["AdderDelegator"]
    adAddr := lib.Address[2:]
    rep = strings.Replace(rep, AdderDelegatorTag, adAddr, -1)
    assert.Equal(t, rep, bytecode)
}

func TestRunLib(t *testing.T) {
    lk := linker.NewLinker("")
    auth := lk.Auth()
    backend := lk.Backend()

    _, _, adder, err := DeployAdder(auth, backend)
    assert.NoError(t, err)

    backend.Commit()

    addSum, addErr := adder.Add(nil, big.NewInt(10), big.NewInt(5))
    assert.NoError(t, addErr)
    assert.Equal(t, addSum, big.NewInt(15))
}


// Doesn't work yet, but problem is identified and will be solved.
// https://github.com/ethereum/go-ethereum/issues/2388

/*
func TestRunLinked(t *testing.T) {
    cwd, err := os.Getwd()
    assert.NoError(t, err)
    moduleDir := path.Join(cwd, "lib_math")

    lk := linker.NewLinker(moduleDir)
    auth := lk.Auth()
    backend := lk.Backend()
    bytecode, dErr := lk.Link(AdderTesterBin)
    assert.NoError(t, dErr)

    var ret0 = new(*big.Int)
    out := ret0
    cErr := lk.Libraries()["Adder"].Contract.Call(nil, out, "add", big.NewInt(10), big.NewInt(5))
    assert.NoError(t,cErr)
    fmt.Println((*ret0).String())
    parsed, jErr := abi.JSON(strings.NewReader(AdderTesterABI))
    assert.NoError(t, jErr)
    fmt.Println(bytecode)
    _, _, contract, dcErr := bind.DeployContract(auth, parsed, common.FromHex(bytecode), backend)
    assert.NoError(t, dcErr)

    backend.Commit()
    var ret1 = new(*big.Int)
    out = ret1
    c2Err := contract.Call(nil, out, "add")
    assert.NoError(t, c2Err)

}
*/