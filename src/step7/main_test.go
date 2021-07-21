// Step7
// ポインタを利用してみましょう．
// あなたの年齢を格納したint型の変数 age を main 関数の中で定義してください．
// その後，引数にageのアドレスを受け取ってageを+1する，birthDay関数を定義してください．
// ※birthDay関数は値をreturnしてはいけません．
package main

import "testing"

func Test_main(t *testing.T) {
	age := 0
	currentAge := age
	birthDay(&age)

	if age != currentAge+1 {
		t.Error("age is not updated")
	}
}
