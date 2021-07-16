package main

import (
	"regexp"
	"testing"

	"github.com/SuperTikuwa/testutil"
)

func Test_main(t *testing.T) {

	got := testutil.ExtractStdout(t, main)
	expect := "Hello,World!"
	reg := `[Hh][Ee][Ll][Ll][Oo][,]?[ ]?[Ww][Oo][Rr][Ll][Dd][.!]?`

	if !regexp.MustCompile(reg).Match([]byte(got)) {
		t.Error("expected: ", expect, "\nactual: ", got)
	}
}
