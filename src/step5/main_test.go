// Step5
// Student構造体を定義して，そのメンバーを表示するプログラムを制作してください．
// 構造体のメンバ変数は，以下のように定義してください．(値は適当に決めてください)
// 	Name: "Taro",
// 	Age:  20,
// 	Number: "K99999"
package main

import (
	"regexp"
	"testing"

	"github.com/SuperTikuwa/testutil"
)

func Test_main(t *testing.T) {
	out := testutil.ExtractStdout(t, main)
	reg := `{.+ [0-9]+ [a-zA-Z][0-9]{5}}`

	if !regexp.MustCompile(reg).MatchString(out) {
		t.Errorf("Wrong answer %s \n Match %s", out, reg)
	}
}
