package test

import (
	"encoding/hex"
	"github.com/androlo/sol-tester/util"
	"github.com/stretchr/testify/assert"
	"path"
	"strings"
	"testing"
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
	lk := newLinker(moduleDir)
	bytecode, lErr := lk.Link(NoLibs)

	assert.NoError(t, lErr)
	assert.Equal(t, bytecode, NoLibs)
}

func TestDeployLibraryNoLinking(t *testing.T) {
	moduleDir := path.Join(cwd, "lib_math")

	lk := newLinker(moduleDir)
	dErr := lk.DeployLibrary("Adder")

	assert.NoError(t, dErr)
	lib := lk.Libraries()["Adder"]
	assert.True(t, util.AddressRe.MatchString(lib.Address.Hex()[2:]))
}

func TestDeployLibraryLinking(t *testing.T) {
	moduleDir := path.Join(cwd, "lib_math")

	lk := newLinker(moduleDir)
	dErr := lk.DeployLibrary("AdderDelegator")

	assert.NoError(t, dErr)
	lib := lk.Libraries()["AdderDelegator"]
	assert.True(t, util.AddressRe.MatchString(lib.Address.Hex()[2:]))
	lib = lk.Libraries()["Adder"]
	assert.True(t, util.AddressRe.MatchString(lib.Address.Hex()[2:]))
}

func TestLinkOne(t *testing.T) {
	moduleDir := path.Join(cwd, "lib_math")

	lk := newLinker(moduleDir)
	bytecode, dErr := lk.Link(Add)

	assert.NoError(t, dErr)
	lib := lk.Libraries()["Adder"]
	assert.True(t, util.AddressRe.MatchString(lib.Address.Hex()[2:]))
	addr := lib.Address.Hex()[2:]
	rep := strings.Replace(Add, AdderTag, addr, -1)
	assert.Equal(t, rep, bytecode)
}

func TestLinkSeveral(t *testing.T) {
	moduleDir := path.Join(cwd, "lib_math")

	lk := newLinker(moduleDir)
	bytecode, dErr := lk.Link(All)

	assert.NoError(t, dErr)
	lib := lk.Libraries()["Adder"]
	adderAddr := lib.Address.Hex()[2:]
	rep := strings.Replace(All, AdderTag, adderAddr, -1)
	lib = lk.Libraries()["Subber"]
	subberAddr := lib.Address.Hex()[2:]
	rep = strings.Replace(rep, SubberTag, subberAddr, -1)
	lib = lk.Libraries()["AdderDelegator"]
	adAddr := lib.Address.Hex()[2:]
	rep = strings.Replace(rep, AdderDelegatorTag, adAddr, -1)
	assert.Equal(t, rep, bytecode)
}

func TestRunLib(t *testing.T) {
	moduleDir := path.Join(cwd, "lib_math")
	lk := newLinker(moduleDir)
	lk.DeployLibrary("Adder")
	vmenv := lk.Environment()
	sender := lk.Sender()
	libAddr := lk.Libraries()["Adder"].Address
	dataStr := "771602f7" +
		"0000000000000000000000000000000000000000000000000000000000000003" +
		"0000000000000000000000000000000000000000000000000000000000000008"
	data, _ := hex.DecodeString(dataStr)
	ret, err := vmenv.Call(sender, libAddr, data, vmGas, vmPrice, vmValue)
	assert.NoError(t, err)
	retStr := hex.EncodeToString(ret)
	assert.Equal(t, retStr, "000000000000000000000000000000000000000000000000000000000000000b")
}

func TestRunLinked(t *testing.T) {
	moduleDir := path.Join(cwd, "lib_math")
	lk := newLinker(moduleDir)
	lk.DeployLibrary("MathTester")
	vmenv := lk.Environment()
	sender := lk.Sender()
	libAddr := lk.Libraries()["MathTester"].Address
	dataStr := "771602f7" +
		"0000000000000000000000000000000000000000000000000000000000000003" +
		"0000000000000000000000000000000000000000000000000000000000000008"
	data, _ := hex.DecodeString(dataStr)
	ret, err := vmenv.Call(sender, libAddr, data, vmGas, vmPrice, vmValue)
	assert.NoError(t, err)
	retStr := hex.EncodeToString(ret)
	assert.Equal(t, retStr, "000000000000000000000000000000000000000000000000000000000000000b")
}
