package main

import (
	"fmt"
	"os"
	"bytes"
	"github.com/androlo/sol-tester/tester"
	"github.com/codegangsta/cli"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/logger/glog"
	"io/ioutil"
	"path"
	"strings"
	"github.com/fatih/color"
)

var (
	app       *cli.App
	DebugFlag = cli.BoolFlag{
		Name:  "debug",
		Usage: "output full trace logs",
	}
	ForceJitFlag = cli.BoolFlag{
		Name:  "forcejit",
		Usage: "forces jit compilation",
	}
	DisableJitFlag = cli.BoolFlag{
		Name:  "nojit",
		Usage: "disabled jit compilation",
	}
	GasFlag = cli.StringFlag{
		Name:  "gas",
		Usage: "gas limit for the evm",
		Value: "10000000000",
	}
	PriceFlag = cli.StringFlag{
		Name:  "price",
		Usage: "price set for the evm",
		Value: "0",
	}
	ValueFlag = cli.StringFlag{
		Name:  "value",
		Usage: "value set for the evm",
		Value: "0",
	}
	CallerAddressFlag = cli.StringFlag{
		Name:  "callerAddress",
		Usage: "caller address for the evm",
		Value: "",
	}
	DirFlag = cli.StringFlag{
		Name:  "dir",
		Usage: "sets the working directory. Defaults to the current working directory.",
		Value: "",
	}
	SysStatFlag = cli.BoolFlag{
		Name:  "sysstat",
		Usage: "display system stats",
	}
	VerbosityFlag = cli.IntFlag{
		Name:  "verbosity",
		Usage: "sets the verbosity level",
	}
)

func init() {
	app = utils.NewApp("0.1.0", "go-solunit")
	app.Flags = []cli.Flag{
		DebugFlag,
		VerbosityFlag,
		ForceJitFlag,
		DisableJitFlag,
		SysStatFlag,
		GasFlag,
		PriceFlag,
		ValueFlag,
		DirFlag,
	}
	app.Action = run
}

func renderLogo() {
	logo := `
*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*
|                      __   __  __           _    __    |
|      _____  ____    / /  / / / /  ____    (_)  / /_   |
|     / ___/ / __ \  / /  / / / /  / __ \  / /  / __/   |
|    (__  ) / /_/ / / /  / /_/ /  / / / / / /  / /_     |
|   /____/  \____/ /_/   \____/  /_/ /_/ /_/   \__/     |
|                                                       |
|                 By: Andreas Olofsson                  |
|             e-mail: androlo1980@gmail.com             |
|                                                       |
*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*
`
	fmt.Printf("\n%s\n\n", logo)
}

func run(ctx *cli.Context) {
	glog.SetToStderr(true)
	glog.SetV(ctx.GlobalInt(VerbosityFlag.Name))
	renderLogo()
	dir := ctx.GlobalString(DirFlag.Name)

	if dir == "" {
		dir, _ = os.Getwd()
	}
	var testFiles []string
	if ctx.NArg() == 0 {
		// If no arguments, search the working directory for testfiles.
		files, rdErr := ioutil.ReadDir(dir)
		if rdErr != nil {
			panic(rdErr)
		}
		testFiles = make([]string, 0)
		for _, file := range files {

			if !file.IsDir() && strings.Index(file.Name(), "Test.bin") > 0 {

				name := file.Name()[0 : len(file.Name())-4]
				_, sErr := os.Stat(name + ".abi")
				if sErr != nil {
					continue
				}
				testFiles = append(testFiles, name)
			}
		}
	} else {
		testFiles = ctx.Args()
	}
	if len(testFiles) == 0 {
		fmt.Println("No tests found. Aborting.")
		return
	} else {
		fmt.Printf("Tests: %v\n", testFiles)
	}
	testContracts := make([]*tester.TestContract, 0)
	for _, cName := range testFiles {
		binName := cName + ".bin"
		abiName := cName + ".abi"

		binBts, binErr := ioutil.ReadFile(path.Join(dir, binName))
		if binErr != nil {
			panic(fmt.Errorf("Can't find bin-file '%s' in working directory: %s", binName, dir))
		}
		bytecode := string(binBts)
		abiBts, abiErr := ioutil.ReadFile(path.Join(dir, abiName))
		if abiErr != nil {
			panic(fmt.Errorf("Can't find abi-file '%s' in working directory: %s", binName, dir))
		}
		abiRdr := bytes.NewReader(abiBts)
		jsonAbi, jsonErr := abi.JSON(abiRdr)
		if jsonErr != nil {
			panic(fmt.Errorf("Failed to convert abi file to abi object: " + jsonErr.Error()))
		}
		testContracts = append(testContracts, tester.NewTestContract(cName, bytecode, jsonAbi))
	}

	gas := ctx.GlobalString(GasFlag.Name)
	price := ctx.GlobalString(PriceFlag.Name)
	value := ctx.GlobalString(ValueFlag.Name)
	debug := ctx.GlobalBool(DebugFlag.Name)
	forceJit := ctx.GlobalBool(ForceJitFlag.Name)
	enableJit := ctx.GlobalBool(DisableJitFlag.Name)
	callerAddress := ctx.GlobalString(CallerAddressFlag.Name)

	testOptions := tester.NewTestOptions(gas, price, value, debug, forceJit, enableJit, callerAddress, dir)

	testRunner, trErr := tester.NewTestRunner(testContracts, testOptions)
	if trErr != nil {
		fmt.Println(trErr)
		return
	}
	_, runErr := testRunner.Run()
	if runErr != nil {
		fmt.Println(runErr)
		return
	}
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		color.White("\n")
		os.Exit(1)
	}
}