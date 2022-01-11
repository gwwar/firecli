package cmd

import (
	"bytes"
	"testing"
)

func TestWriteWordShort(t *testing.T) {
	www := newWriter(10)
	out := &bytes.Buffer{}
	error := writeWord(out, www, "foo")
	if error != nil {
		t.Errorf("writeWord should not return error but got %q", error)
	}
	expected := ". foo"
	if out.String() != expected {
		t.Errorf("expected %q but got %q", expected, out.String())
	}
}

func TestWriteWordLongerThanLineLength(t *testing.T) {
	www := newWriter(10)
	out := &bytes.Buffer{}
	error := writeWord(out, www, "fantastical")
	if error != nil {
		t.Errorf("writeWord should not return error but got %q", error)
	}
	expected := ". fantastic-\n. al"
	if out.String() != expected {
		t.Errorf("expected %q but got %q", expected, out.String())
	}
}

func TestWriteWordExactlyOnLineWidth(t *testing.T) {
	www := newWriter(10)
	out := &bytes.Buffer{}
	error := writeWord(out, www, "0123456789")
	if error != nil {
		t.Errorf("writeWord should not return error but got %q", error)
	}
	expected := ". 0123456789"
	if out.String() != expected {
		t.Errorf("expected %q but got %q", expected, out.String())
	}
}

func TestWriteWordMultipleCalls(t *testing.T) {
	var err error
	var expected string
	www := newWriter(10)
	out := &bytes.Buffer{}
	err = writeWord(out, www, "Hello")
	if err != nil {
		t.Errorf("writeWord should not return error but got %q", err)
	}
	expected = ". Hello"
	if out.String() != expected {
		t.Errorf("expected %q but got %q", expected, out.String())
	}
	err = writeWord(out, www, "World!")
	if err != nil {
		t.Errorf("writeWord should not return error but got %q", err)
	}
	expected = ". Hello      .\n. World!"
	if out.String() != expected {
		t.Errorf("expected %q but got %q", expected, out.String())
	}
}

func TestLineReturn(t *testing.T) {
	www := newWriter(10)
	out := &bytes.Buffer{}
	err := writeWord(out, www, "1234567890")
	if err != nil {
		t.Errorf("writeWord should not return error but got %q", err)
	}
	expected := ". 1234567890"
	if out.String() != expected {
		t.Errorf("expected %q but got %q", expected, out.String())
	}
	err = lineReturn(out, www)
	if err != nil {
		t.Errorf("lineReturn should not return error but got %q", err)
	}
	expected = ". 1234567890 .\n"
	if out.String() != expected {
		t.Errorf("expected %q but got %q", expected, out.String())
	}
}

func TestLineReturnWithPadding(t *testing.T) {
	www := newWriter(10)
	out := &bytes.Buffer{}
	err := writeWord(out, www, "1234567")
	if err != nil {
		t.Errorf("writeWord should not return error but got %q", err)
	}
	expected := ". 1234567"
	if out.String() != expected {
		t.Errorf("expected %q but got %q", expected, out.String())
	}
	err = lineReturn(out, www)
	if err != nil {
		t.Errorf("lineReturn should not return error but got %q", err)
	}
	expected = ". 1234567    .\n"
	if out.String() != expected {
		t.Errorf("expected %q but got %q", expected, out.String())
	}
}

func TestCatsay(t *testing.T) {
	www := newWriter(30)
	out := &bytes.Buffer{}
	data := []byte(`Alice was beginning to get very tired of sitting by her sister on the bank`)
	err := catsay(out, bytes.NewReader(data), www, false)
	if err != nil {
		t.Errorf("catsay should not return error but got %q", err)
	}
	expected := "----------------------------------\n. Alice was beginning to get     .\n. very tired of sitting by her   .\n. sister on the bank             .\n"
	if out.String() != expected {
		t.Errorf("expected %q but got %q", expected, out.String())
	}
}
