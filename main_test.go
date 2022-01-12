package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunCommand(t *testing.T) {
	out := &bytes.Buffer{}
	exit := runCommand(out)
	if exit != exitOK {
		t.Errorf("expected exit with signal 0 but got %q", exit)
	}
	prefix := "cli playground for learning go:"
	if !strings.HasPrefix(out.String(), prefix) {
		t.Errorf("expected to start with %q but got %q", prefix, out.String())
	}
}
