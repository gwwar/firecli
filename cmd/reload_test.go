package cmd

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestReload(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.String() == "/-/reload" {
			w.WriteHeader(200)
		} else {
			w.WriteHeader(500)
		}
	}))
	defer server.Close()

	out := &bytes.Buffer{}
	err := reload(server.URL, out)

	if err != nil {
		t.Errorf("reload should not return error but got %q", err)
	}
	prefix := "Successfully reloaded prometheus configs"
	if !strings.HasPrefix(out.String(), prefix) {
		t.Errorf("reload should be successful but response was %q", out.String())
	}
}
