package builder

import (
	"bytes"
	"fmt"
	"os/exec"
)

type Solc struct{}

func (this *Solc) Compile(args []string, cwd string) error {
	return this.solc(args, cwd)
}

func (this *Solc) Version() error {
	return this.solc([]string{"--version"}, "")
}

func (this *Solc) solc(args []string, cwd string) error {
	fmt.Printf("Solc args: %v\n", args)
	cmd := exec.Command("solc", args...)
	if cwd != "" {
		cmd.Dir = cwd
	}
	var out bytes.Buffer
	var outErr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &outErr
	err := cmd.Run()
	if out.Len() > 0 {
		fmt.Printf("%s\n", out.String())
	}
	if outErr.Len() > 0 {
		fmt.Printf("%s\n", outErr.String())
		return fmt.Errorf("Error when running Solc: %s\n", outErr.String())
	}
	if err != nil {
		return err
	}
	return nil
}
