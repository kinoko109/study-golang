package main

import (
	"fmt"
	"strconv"
	"time"
)

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}

	// 定数
	const P1 = 3.14 // 頭文字を大文字にするとパブリック変数になる

	const (
		Domain = "test.co.jp"
		SiteName = "test"
	)

	const (
		first = iota // iotaは連番を自動採番して返す
		second
		third
	)

func main() {
	fmt.Println("Hellow World")
	fmt.Print(time.Now())

	// 明示的な変数定義

	var i int = 100
	fmt.Println(i)

	var s string = "文字列です"
	fmt.Println(s)

	//　複数定義
	var t, f bool = true, false
	fmt.Println(t, f)

	var(
		i2 int = 200
		s2 string = "文字列です2"
	)
	fmt.Println(i2, s2)

	// 初期値を代入しない場合は、自動的にその型に合った初期値が自動で定義される
	var i3 int
	var s3 string
	fmt.Println(i3, s3) // 0, ""

	i3 = 10000
	s3 = "3423"
	fmt.Println(i3, s3)

	// 暗黙的な定義 - 変数定義と代入を同時に行う
	i4 := 4000
	s4 := "dfjaldsjkfa"
	fmt.Println(i4, s4)

	// i4はint型なので、エラーになる
	// i4 = "dafafdsf"
	// fmt.Println(i4)

	outer()

	// int型
	var i64 int64 = 10000
	fmt.Println(i64)

	// 違うint型同士はエラーになる
	// fmt.Println(i + i64)

	// 型変換を行うことで計算できる
	fmt.Println(int64(i) + i64)
	// 型を表示
	fmt.Printf("%T\n", int64(i))

	// string
	fmt.Println(`test
	test
	`)

	// byte型
	byteA := []byte{72, 73}
	fmt.Println(byteA)

	fmt.Println(string(byteA))

	c := []byte(byteA)
	fmt.Println(c)

	// 配列 - Goの配列型は要素数を変更できない。要素数を変える場合は、スライス型を使用する。
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)

	// index2番目は未定義なので、空文字になる
	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)

	// スプレッド
	arr4 := [...]string{"A", "B"}
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4)

	fmt.Println(arr1[0])
	fmt.Println(arr2[1])

	// 要素変更
	fmt.Println(arr1[0])
	arr1[0] = 592375927
	fmt.Println(arr1[0])

	//　要素数
	fmt.Println(len(arr2))

	// interface & nil
	// interfaceの初期値はｎilで、どんな型とも互換性がある
	var interfaceA interface{}
	fmt.Println(interfaceA)
	interfaceA = 1
	fmt.Println(interfaceA)
	interfaceA = "aaaa"
	fmt.Println(interfaceA)
	interfaceA = [1]int{1}
	fmt.Println(interfaceA)

	// 型変換
	var stringVar string = "1000"
	fmt.Printf("stringVar = %T\n", stringVar)

	numberVar, _ := strconv.Atoi(stringVar)
	fmt.Println(numberVar)
	fmt.Printf("numberVar = %T\n", numberVar)

	var h string = "Hello"
	b := []byte(h)
	fmt.Println(b)

	h2 := string(b)
	fmt.Println(h2)

	fmt.Println(P1)
	fmt.Println(Domain)
	fmt.Println(SiteName)
	fmt.Println(first, second, third)

	plusValue := Plus(1, 2)
	fmt.Println(plusValue)

	divValue1, divValu2 := Div(10, 2)
	fmt.Println(divValue1, divValu2)

	doubleValue := Double(2)
	fmt.Println(doubleValue)

	// 無名関数
	funcMumei := func(x, y int)int {
		return x + y
	}
	mumei := funcMumei(1, 1)
	fmt.Println(mumei)

	// 即時実行
	funcMumei2 := func(x, y int)int {
		return x + y
	}(1,100)
	fmt.Println(funcMumei2)
}

// 関数
func Plus(x, y int)int {
	return x + y
}

func Div(x, y int)(int, int) {
	a := x / y
	b := x % y

	return a, b
}

// 返り値の型に変数名指定すると、returnで省略できる
func Double(price int)(result int) {
	result = price * 2
	return
}
