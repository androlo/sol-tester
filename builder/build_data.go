package builder

import (
	"encoding/json"
	"io/ioutil"
)

// Struct representation of a build.json file.
type BuildData struct {
	Module          string            `json:module`
	Units           [][]string        `json:contracts`
	CompilerOptions *CompilerOptions  `json:options`
	Includes        map[string]string `json:includes`
}

type CompilerOptions struct {
	SourcePath   string `json:sourcePath`
	OutputPath   string `json:outputPath`
	Bin          bool   `json:bin`
	Abi          bool   `json:abi`
	Optimize     bool   `json:optimize`
	OptimizeRuns uint32 `json:optimizeRuns`
}

func LoadBuildData(path string) (*BuildData, error) {
	data, readErr := ioutil.ReadFile(path)
	if readErr != nil {
		return nil, readErr
	}
	bd := &BuildData{}
	umErr := json.Unmarshal(data, bd)
	if umErr != nil {
		return nil, umErr
	}
	return bd, nil
}
