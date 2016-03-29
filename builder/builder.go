package builder

import (
	"fmt"
	"github.com/androlo/sol-tester/util"
	"path"
	"strconv"
)

const BUILD_FILE_DEFAULT_NAME = "build.json"

const SOURCE_DIR_DEFAULT_NAME = "src"

const OUTPUT_DIR_DEFAULT_NAME = "build"

type Builder struct {
	solc *Solc
}

func NewBuilder() *Builder {
	return &Builder{&Solc{}}
}

func (this *Builder) BuildFromBuildData(moduleRoot string, bd *BuildData) error {
	if bd.Module == "" {
		return fmt.Errorf("No module name.")
	}
	if bd.Units == nil || len(bd.Units) == 0 {
		return fmt.Errorf("No compilation units.")
	}
	if bd.CompilerOptions == nil {
		bd.CompilerOptions = &CompilerOptions{}
	}
	if bd.CompilerOptions.SourcePath == "" {
		bd.CompilerOptions.SourcePath = path.Join(moduleRoot, SOURCE_DIR_DEFAULT_NAME)
	}
	if bd.CompilerOptions.OutputPath == "" {
		bd.CompilerOptions.OutputPath = path.Join(moduleRoot, OUTPUT_DIR_DEFAULT_NAME)
	}
	srcDir := bd.CompilerOptions.SourcePath
	buildDir := bd.CompilerOptions.OutputPath

	commands := util.NewStrArr()
	if bd.Includes != nil {
		commands.Add("-i")
		for k, v := range bd.Includes {
			commands.Add(k)
			commands.Add(v)
		}
	}
	if bd.CompilerOptions.Bin {
		commands.Add("--bin")
	}
	if bd.CompilerOptions.Abi {
		commands.Add("--abi")
	}
	if bd.CompilerOptions.Optimize {
		commands.Add("--optimize")
		if bd.CompilerOptions.OptimizeRuns != 0 {
			commands.Add("--optimize-runs")
			commands.Add(strconv.Itoa(int(bd.CompilerOptions.OptimizeRuns)))
		}
	}
	commands.Add("-o")
	commands.Add(buildDir)
	for i := 0; i < len(bd.Units); i++ {
		unit := bd.Units[i]
		thisCmd := *commands
		for j := 0; j < len(unit); j++ {
			cName := unit[j]
			cPath := path.Join(srcDir, cName+".sol")
			thisCmd = append(thisCmd, cPath)
		}
		cErr := this.solc.Compile(thisCmd, moduleRoot)
		if cErr != nil {
			return cErr
		}
	}
	return nil
}

func (this *Builder) Build(moduleRoot, buildFile string) error {

	if buildFile == "" {
		buildFile = BUILD_FILE_DEFAULT_NAME
	}
	buildFilePath := path.Join(moduleRoot, buildFile)
	bd, loadErr := LoadBuildData(buildFilePath)
	if loadErr != nil {
		return loadErr
	}
	return this.BuildFromBuildData(moduleRoot, bd)
}
