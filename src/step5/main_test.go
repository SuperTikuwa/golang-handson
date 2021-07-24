// Step5
// Collectionを使ってみましょう
// まず，配列とスライスの違いを理解しましょう．
// その後，スライスの先頭の要素と最後の要素を消去する removeFirstLast(strSlice []string)[]string 関数を定義しましょう．
// 更に，mapを利用してみましょう．
// 人物名をmapのキーとして，その人物の年齢を値として持つmapを定義しましょう．
// そして，mapを受け取って，すべての人物の名前と年齢を連結した文字列の配列を返す getAllNameAndAge(strSlice [string]int)[]string を定義しましょう．
// (例) "Taro":12, "Jiro":19, "Saburo":18 のようなmapを受けとったとき，
// 		"Jiro 19", "Saburo 18", "Taro 12" のような文字列の配列を返すようにしてください．
package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func Test_main(t *testing.T) {
	// removeFirstLast関数のテスト
	strSlice := []string{"a", "b", "c", "d", "e"}
	strSlice2 := removeFirstLast(strSlice)
	if strSlice2[0] != "b" || strSlice2[len(strSlice2)-1] != "d" {
		t.Error("removeFirstLast関数のテストでエラーが発生しました")
	}
	// getAllNameAndAge関数のテスト
	// 人物名をキーに持ち，年齢を値に持つmapを定義する
	nameAndAge := map[string]int{
		"Alice":   30,
		"Bob":     40,
		"Charlie": 50,
		"David":   20,
		"Eve":     40,
		"Fred":    10,
		"Ginny":   50,
		"Harriet": 20,
		"Ileana":  40,
		"Joseph":  10,
		"Kincaid": 30,
		"Larry":   50,
	}
	// getAllNameAndAge関数を実行して，その結果を文字列の配列として返す
	strSlice3 := getAllNameAndAgeOrder(nameAndAge)
	for k, v := range nameAndAge {
		for i, s := range strSlice3 {
			if strings.Contains(s, k) && strings.Contains(s, strconv.Itoa(v)) {
				break
			}

			if i == len(strSlice3)-1 && (!strings.Contains(s, k) || !strings.Contains(s, strconv.Itoa(v))) {
				t.Error("getAllNameAndAge関数のテストでエラーが発生しました", k, v, i, s)
			}
		}
	}
	fmt.Println(strSlice3)
	for i := range strSlice3 {
		if len(strSlice3)-1 == i {
			break
		}

		if strSlice3[i] > strSlice3[i+1] {
			t.Error("getAllNameAndAge関数のテストでエラーが発生しました")
			break
		}
	}

}
