package main

import (
	"regexp"
	"strings"
	"testing"

	"github.com/SuperTikuwa/testutil"
)

func Test_main(t *testing.T) {

	got := testutil.ExtractStdout(t, main)
	out := strings.Split(got, "\n")
	expect := "1 2 3"

	if !regexp.MustCompile(`[0-9 ]+`).Match([]byte(out[0])) {
		t.Error("expected: ", expect, "\nactual: ", got)
	}

	expect = "String"

	if !regexp.MustCompile(`.+`).Match([]byte(out[1])) {
		t.Error("expected: ", expect, "\nactual: ", got)
	}
}
