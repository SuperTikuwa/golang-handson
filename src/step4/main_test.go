// Step4
// 引数が3の倍数ならFizz,5の倍数ならBuzz,
// 3の倍数かつ5の倍数ならFizzBuzz,
// どれでもない場合は引数をそのまま出力するFizzBuzz関数を制作し，
// 引数に1~100を与えて実行してください．
package main

import (
	"strconv"
	"strings"
	"testing"

	"github.com/SuperTikuwa/testutil"
)

func Test_main(t *testing.T) {
	out := testutil.ExtractStdout(t, main)
	arr := strings.Split(out, "\n")

	for i, v := range arr {
		if fizzBuzz(i+1) != v {
			t.Errorf("Wrong answer %d:%s \n", i+1, v)
		}
	}
}

// Return Fizz or Buzz or FizzBuzz or the original number
func fizzBuzz(n int) string {
	if n%3 == 0 && n%5 == 0 {
		return "FizzBuzz"
	} else if n%3 == 0 {
		return "Fizz"
	} else if n%5 == 0 {
		return "Buzz"
	}
	return strconv.Itoa(n)
}
