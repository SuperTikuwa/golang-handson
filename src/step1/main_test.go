package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func Test_main(t *testing.T) {
	got := extractStdout(t, main)
	expect := "Hello,World!"

	if got != expect {
		t.Error("expected: ", expect, "\nactual: ", got)
	}
}

func extractStdout(t *testing.T, fnc func()) string {
	t.Helper()

	orgStdout := os.Stdout

	defer func() {
		os.Stdout = orgStdout
	}()

	r, w, _ := os.Pipe()
	os.Stdout = w
	main()
	w.Close()

	var buf bytes.Buffer

	if _, err := buf.ReadFrom(r); err != nil {
		t.Fatalf("failed to read buf: %v", err)
	}

	return strings.TrimRight(buf.String(), "\n")
}
