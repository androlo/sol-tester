package test

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestCreateLinker(t *testing.T) {
	lk := newLinker("foo")
	libs := lk.Libraries()
	assert.NotNil(t, libs)
	assert.Len(t, libs, 0)
	assert.Equal(t, lk.RootDir(), "foo")
}

func TestSetRootDir(t *testing.T) {
	lk := newLinker("")
	lk.SetRootDir("bar")
	assert.Equal(t, lk.RootDir(), "bar")
}

func TestAddLibrariesFailNoLibraryFile(t *testing.T) {
	baseDir, err := ioutil.TempDir("", "st")
	assert.NoError(t, err)
	lk := newLinker("")
	fileName := path.Join(baseDir, "baz")
	_, fErr := os.Stat(fileName)
	assert.True(t, os.IsNotExist(fErr))
	aErr := lk.AddLibrariesFromFile(fileName)
	assert.Error(t, aErr)
}

func TestAddLibrariesFailLibFileEmpty(t *testing.T) {
	moduleDir := path.Join(cwd, "lib_parse_file")
	lk := newLinker(moduleDir)

	aErr := lk.AddLibrariesFromFile(path.Join(moduleDir, "libraries_empty.csv"))
	assert.Error(t, aErr)
}

func TestAddLibrariesFailMissingColon(t *testing.T) {
	moduleDir := path.Join(cwd, "lib_parse_file")
	lk := newLinker(moduleDir)
	aErr := lk.AddLibrariesFromFile(path.Join(moduleDir, "libraries_missing_colon.csv"))
	assert.Error(t, aErr)
}

func TestAddLibrariesFailNameEmpty(t *testing.T) {
	moduleDir := path.Join(cwd, "lib_parse_file")
	lk := newLinker(moduleDir)
	aErr := lk.AddLibrariesFromFile(path.Join(moduleDir, "libraries_name_empty.csv"))
	assert.Error(t, aErr)
}

func TestAddLibrariesFailNameNotId(t *testing.T) {
	moduleDir := path.Join(cwd, "lib_parse_file")
	lk := newLinker(moduleDir)
	aErr := lk.AddLibrariesFromFile(path.Join(moduleDir, "libraries_name_not_id.csv"))
	assert.Error(t, aErr)
}

func TestAddLibrariesFailAddressEmpty(t *testing.T) {
	moduleDir := path.Join(cwd, "lib_parse_file")
	lk := newLinker(moduleDir)
	aErr := lk.AddLibrariesFromFile(path.Join(moduleDir, "libraries_address_empty.csv"))
	assert.Error(t, aErr)
}

func TestAddLibrariesFailAddressNotValid(t *testing.T) {
	moduleDir := path.Join(cwd, "lib_parse_file")
	lk := newLinker(moduleDir)
	aErr := lk.AddLibrariesFromFile(path.Join(moduleDir, "libraries_address_not_valid.csv"))
	assert.Error(t, aErr)
}

func TestAddLibrariesFailAddressHexButWrong(t *testing.T) {
	moduleDir := path.Join(cwd, "lib_parse_file")
	lk := newLinker(moduleDir)
	aErr := lk.AddLibrariesFromFile(path.Join(moduleDir, "libraries_address_hex_but_wrong.csv"))
	assert.Error(t, aErr)
}

func TestAddLibrariesSuccessOneWorking(t *testing.T) {
	moduleDir := path.Join(cwd, "lib_parse_file")
	lk := newLinker(moduleDir)
	aErr := lk.AddLibrariesFromFile(path.Join(moduleDir, "libraries_one_working.csv"))
	assert.NoError(t, aErr)
	libs := lk.Libraries()
	assert.Len(t, libs, 1)
	lib, exists := libs["first"]
	assert.True(t, exists)
	assert.Equal(t, lib.Address.Hex(), "0x1234567812345678123456781234567812345678")
	assert.Empty(t, lib.Bytecode)
}

func TestAddLibrariesSuccessTwoWorking(t *testing.T) {
	moduleDir := path.Join(cwd, "lib_parse_file")
	lk := newLinker(moduleDir)
	aErr := lk.AddLibrariesFromFile(path.Join(moduleDir, "libraries_two_working.csv"))
	assert.NoError(t, aErr)
	libs := lk.Libraries()
	assert.Len(t, libs, 2)
	lib, exists := libs["first"]
	assert.True(t, exists)
	assert.Equal(t, lib.Address.Hex(), "0x1234567812345678123456781234567812345678")
	assert.Empty(t, lib.Bytecode)
	lib, exists = libs["second"]
	assert.True(t, exists)
	assert.Equal(t, lib.Address.Hex(), "0x0012afbd0012afbd0012afbd0012afbd0012afbd")
	assert.Empty(t, lib.Bytecode)
}
