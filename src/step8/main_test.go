// Step8
// Step6で制作したStudent構造体に，以下のようなメソッドを追加してください．
// 	- getName() : 名前を返すメソッド
// 	- getAge() : 年齢を返すメソッド
// 	- getNumber() : 学籍番号を返すメソッド
// 	- selfIntroduction() : 自己紹介文を返すメソッド
package main

import (
	"strconv"
	"strings"
	"testing"
)

func Test_main(t *testing.T) {

	m := student{
		name:   "Taro",
		age:    20,
		number: "2012345",
	}
	intro := m.selfIntroduction()
	if !strings.Contains(intro, m.name) {
		t.Error("name is not included")
	}
	if !strings.Contains(intro, strconv.Itoa(m.age)) {
		t.Error("age is not included")
	}
	if !strings.Contains(intro, m.number) {
		t.Error("number is not included")
	}

	primaryAge := m.getAge()
	m.birthDay()
	if primaryAge+1 != m.getAge() {
		t.Error("age is not changed")
	}

}
