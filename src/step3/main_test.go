package main

import (
	"regexp"
	"testing"

	"github.com/SuperTikuwa/testutil"
)

func Test_main(t *testing.T) {
	out := testutil.ExtractStdout(t, main)
	reg := `[Hh]ello.+`
	if !regexp.MustCompile(reg).Match([]byte(out)) {
		t.Error("Invalid output: ", out)
	}

	names := []string{"Alice", "Bob", "Chris"}
	for _, name := range names {
		r := sayHello(name)
		reg = `[Hh]ello.` + name + `.*`
		if !regexp.MustCompile(reg).Match([]byte(r)) {
			t.Error("Invalid return: ", r)
		}
	}
}
