package test

import (
	"github.com/androlo/sol-tester/builder"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"path"
	"testing"
)

func TestBuildFailNoModuleName(t *testing.T) {
	moduleDir := path.Join(cwd, "no_module_name")
	bd := builder.NewBuilder()
	bErr := bd.Build(moduleDir, "")
	assert.Error(t, bErr)
}

func TestBuildFailNoUnits(t *testing.T) {
	moduleDir := path.Join(cwd, "no_units")
	bd := builder.NewBuilder()
	bErr := bd.Build(moduleDir, "")
	assert.Error(t, bErr)
}

func TestBuildFailCannotFindContractSourceFile(t *testing.T) {
	moduleDir := path.Join(cwd, "math_no_sol")
	bd := builder.NewBuilder()
	bErr := bd.Build(moduleDir, "")
	assert.Error(t, bErr)
}

func TestBuildSuccess(t *testing.T) {

	moduleDir := path.Join(cwd, "math")
	bdta, lErr := builder.LoadBuildData(path.Join(moduleDir, "build.json"))
	assert.NoError(t, lErr)

	outPath, tcErr := ioutil.TempDir("", "sol_tester_test_")
	assert.NoError(t, tcErr)
	bdta.CompilerOptions.OutputPath = outPath
	bd := builder.NewBuilder()
	bErr := bd.BuildFromBuildData(moduleDir, bdta)
	assert.NoError(t, bErr)

	outPathContent, rdErr := ioutil.ReadDir(outPath)
	assert.NoError(t, rdErr)
	assert.NotNil(t, outPathContent)
	assert.Len(t, outPathContent, 2)

	opcAbi := outPathContent[0]
	pbAbi := path.Base(opcAbi.Name())
	assert.Equal(t, pbAbi, "Adder.abi")

	opcBin := outPathContent[1]
	pbBin := path.Base(opcBin.Name())
	assert.Equal(t, pbBin, "Adder.bin")

}
