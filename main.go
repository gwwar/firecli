package main

import (
	"github.com/gwwar/firecli/cmd"
	"io"
	"os"
)

type exitCode int

const (
	exitOK    exitCode = 0
	exitError exitCode = 1
)

func runCommand(out io.Writer) exitCode {
	rootCommand := cmd.GetRootCmd()
	rootCommand.SetOut(out)
	err := rootCommand.Execute()
	if err != nil {
		return exitError
	}
	return exitOK
}

func main() {
	code := runCommand(os.Stdout)
	os.Exit(int(code))
}
