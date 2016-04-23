package builder

import (
	"fmt"
	"github.com/androlo/sol-tester/util"
	"os"
	"path"
	"strconv"
)

const BUILD_FILE_DEFAULT_NAME = "build.json"

const SOURCE_DIR_DEFAULT_NAME = "src"
const TEST_DIR_DEFAULT_NAME = "test"

var outputDirDefaultName = path.Join("build", "release")

var testOutputDirDefaultName = path.Join("build", "test")

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
	if (bd.Units == nil || len(bd.Units) == 0) && (bd.Tests == nil || len(bd.Tests) == 0) {
		return fmt.Errorf("No compilation units or tests.")
	}
	if bd.CompilerOptions == nil {
		bd.CompilerOptions = &CompilerOptions{}
	}
	if bd.CompilerOptions.SourcePath == "" {
		bd.CompilerOptions.SourcePath = path.Join(moduleRoot, SOURCE_DIR_DEFAULT_NAME)
	}
	if bd.CompilerOptions.OutputPath == "" {
		bd.CompilerOptions.OutputPath = path.Join(moduleRoot, outputDirDefaultName)
	}
	srcDir := bd.CompilerOptions.SourcePath
	buildDir := bd.CompilerOptions.OutputPath
	raErr := os.RemoveAll(buildDir)
	if raErr != nil {
		return raErr
	}

	commands := util.NewStrArr()
	commands.Add("-i")
	commands.Add("=.")
	if bd.Includes != nil {
		for k, v := range bd.Includes {
			commands.Add(k + "=" + v)
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
	// Normal build
	if bd.Units != nil && len(bd.Units) != 0 {
		relCommands := *commands
		relCommands = append(relCommands, buildDir)
		for i := 0; i < len(bd.Units); i++ {
			unit := bd.Units[i]
			for j := 0; j < len(unit); j++ {
				cName := unit[j]
				cPath := path.Join(srcDir, cName+".sol")
				relCommands = append(relCommands, cPath)
			}
			cErr := this.solc.Compile(relCommands, moduleRoot)
			if cErr != nil {
				return cErr
			}
		}
	}
	// Tests
	if bd.Tests != nil && len(bd.Tests) != 0 {
		var testDir string
		if bd.CompilerOptions.TestPath == "" {
			testDir = path.Join(moduleRoot, TEST_DIR_DEFAULT_NAME)
		} else {
			testDir = bd.CompilerOptions.TestPath
		}
		if bd.CompilerOptions.TestPath == "" {
			bd.CompilerOptions.TestPath = path.Join(moduleRoot, TEST_DIR_DEFAULT_NAME)
		}
		var testBuildDir string
		if bd.CompilerOptions.TestOutputPath == "" {
			testBuildDir = path.Join(moduleRoot, testOutputDirDefaultName)
		} else {
			testBuildDir = bd.CompilerOptions.TestOutputPath
		}

		if bd.Tests != nil && len(bd.Tests) > 0 {
			testCommands := *commands
			testCommands = append(testCommands, testBuildDir)
			for i := 0; i < len(bd.Tests); i++ {
				test := bd.Tests[i]
				tPath := path.Join(testDir, test+".sol")
				testCommands = append(testCommands, tPath)
			}
			tErr := this.solc.Compile(testCommands, moduleRoot)
			if tErr != nil {
				return tErr
			}
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
