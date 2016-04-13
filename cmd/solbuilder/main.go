package main

import (
	"fmt"
	"os"

	"github.com/androlo/sol-tester/builder"
	"github.com/codegangsta/cli"
	"github.com/ethereum/go-ethereum/cmd/utils"
)

var (
	app       *cli.App
	DebugFlag = cli.BoolFlag{
		Name:  "debug",
		Usage: "additional output",
	}
	DirFlag = cli.StringFlag{
		Name:  "dir",
		Usage: "sets the working directory. Defaults to the current working directory.",
		Value: "",
	}
	BuildFileFlag = cli.StringFlag{
		Name:  "buildfile",
		Usage: "The name of the buildfile. Defaults to 'build.json'.",
		Value: "build.json",
	}
)

func init() {
	app = utils.NewApp("0.1.0", "go-solbuilder")
	app.Flags = []cli.Flag{
		DebugFlag,
		DirFlag,
		BuildFileFlag,
	}
	app.Action = run
}

func renderLogo() {
	logo := `
*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*
|                _ ____        _ _     _                |
|               | |  _ \      (_) |   | |               |
|      ___  ___ | | |_) |_   _ _| | __| | ___ _ __      |
|     / __|/ _ \| |  _ <| | | | | |/ _' |/ _ \ '__|     |
|     \__ \ (_) | | |_) | |_| | | | (_| |  __/ |        |
|      ___/\___/|_|____/ \__,_|_|_|\__,_|\___|_|        |
|                                                       |
|                 By: Andreas Olofsson                  |
|             e-mail: androlo1980@gmail.com             |
|                                                       |
*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*
`
	fmt.Printf("\n%s\n\n", logo)
}

func run(ctx *cli.Context) {
	renderLogo()
	dir := ctx.GlobalString(DirFlag.Name)

	if len(dir) == 0 {
		dir, _ = os.Getwd()
	}

	buildfile := ctx.GlobalString(BuildFileFlag.Name)
	fmt.Println(dir)
	fmt.Println(buildfile)
	bd := builder.NewBuilder()
	bErr := bd.Build(dir, buildfile)
	if bErr != nil {
		panic(bErr)
	}
	fmt.Println("Done")
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
