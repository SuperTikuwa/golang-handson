// Step6
// student構造体を定義して，そのメンバーを表示するプログラムを制作してください．
// 構造体のメンバ変数は，以下のように定義してください．(値は適当に決めてください)
// 	name: "Taro",
// 	age:  20,
// 	number: "K99999"
package main

import (
	"regexp"
	"strings"
	"testing"

	"github.com/SuperTikuwa/testutil"
)

func Test_main(t *testing.T) {
	out := testutil.ExtractStdout(t, main)

	if !strings.Contains(out, "name") || !strings.Contains(out, "age") || !strings.Contains(out, "number") {
		t.Errorf("Wrong output\n")
	}

	taro := student{
		name:   "Taro",
		age:    20,
		number: "K99999",
	}

	reg := `[a-zA-Z]+`
	if !regexp.MustCompile(reg).MatchString(taro.name) {
		t.Errorf("Wrong output\n")
	}

	if taro.age < 0 {
		t.Errorf("Wrong output\n")
	}

	reg = `[a-zA-Z0-9]+`
	if !regexp.MustCompile(reg).MatchString(taro.number) {
		t.Errorf("Wrong output\n")
	}

}
