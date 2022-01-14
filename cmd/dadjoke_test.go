package cmd

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRandomDadJoke(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		joke := &Joke{
			ID:        "foobarbaz",
			Punchline: "Why are humans known to be extremely afraid of computers? Probably, because they byte!",
			Status:    200,
		}
		err := json.NewEncoder(w).Encode(joke)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
	}))
	defer server.Close()

	out := &bytes.Buffer{}
	err := randomDadJoke(server.URL, out)

	if err != nil {
		t.Errorf("dadjoke should not return error but got %q", err)
	}
	expected := "Why are humans known to be extremely afraid of computers? Probably, because they byte!\n"
	if out.String() != expected {
		t.Errorf("dadjoke should be successful but response was %q", out.String())
	}
}
