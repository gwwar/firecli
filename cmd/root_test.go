package cmd

import "testing"

func TestGetRootCmd(t *testing.T) {
	command := GetRootCmd()
	use := command.Use
	expected := "firecli"
	if use != expected {
		t.Errorf("expected %q but got %q", expected, use)
	}
}
