package main

import (
	"bytes"
	"testing"
)

func TestRunCommand(t *testing.T) {
	out := &bytes.Buffer{}
	exit := runCommand(out)
	if exit != exitOK {
		t.Errorf("expected exit with signal 0 but got %q", exit)
	}
	expected := "cli playground for learning go: it will probably have some silly commands\n\nUsage:\n  firecli [command]\n\nAvailable Commands:\n  catsay      A speaking cat\n  completion  Generate the autocompletion script for the specified shell\n  help        Help about any command\n\nFlags:\n  -h, --help     help for firecli\n  -t, --toggle   Help message for toggle\n\nUse \"firecli [command] --help\" for more information about a command.\n"
	if out.String() != expected {
		t.Errorf("expected %q but got %q", expected, out.String())
	}
}
