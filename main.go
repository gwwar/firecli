package main

import (
	"github.com/gwwar/firecli/cmd"
	"os"
)

type exitCode int

const (
	exitOK     exitCode = 0
	exitError  exitCode = 1
	exitCancel exitCode = 2
	exitAuth   exitCode = 4
)

func runCommand() exitCode {
	err := cmd.RootCmd.Execute()
	if ( err != nil ) {
		return exitError
	}
	return exitOK
}

func main() {

	code := runCommand()
	os.Exit(int(code))
}
