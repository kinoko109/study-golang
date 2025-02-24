package main

import (
	"fmt"
	"os"
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

	returnedFunc := ReturnFunc()
	returnedFunc()

	CallbackFunc(func() {
		fmt.Println("im a callback function!!")
	})

	closureFunc := Closure()
	fmt.Println(closureFunc("hello"))
	fmt.Println(closureFunc("my"))
	fmt.Println(closureFunc("name"))

	generatorIntegers := intergers()
	fmt.Println(generatorIntegers())
	fmt.Println(generatorIntegers())
	fmt.Println(generatorIntegers())
	fmt.Println(generatorIntegers())

	generatorIntegers2 := intergers()
	fmt.Println(generatorIntegers2())
	fmt.Println(generatorIntegers2())
	fmt.Println(generatorIntegers2())

	// if
	condition := 0
	if(condition == 1) {
		fmt.Println("one")
	} else if condition == 2 {
		fmt.Println("two")
	} else {
		fmt.Println("other")
	}

	if conditionA := 100; conditionA == 100 {
		fmt.Println("100です")
	}

	conditionB := 0
	if conditionB := 100; true {
		fmt.Println(conditionB)
	}
	fmt.Println(conditionB)

	// エラーハンドリング
	var status string = "error"
	errorValue, err := strconv.Atoi(status)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("errorValue = %T\n", errorValue)

	// for
	point := 0
	for point < 10 {
		fmt.Println(point)
		point++
	}

	arr := [3]int{1, 2, 3}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	for i, v := range arr {
		fmt.Println(i, v)
	}

	sl := []string{"python", "php", "go"}
	for i, v := range sl {
		fmt.Println(i, v)
	}

	// map
	map1 := map[string]int {
		"apple": 100,
		"banana": 500,
	}
	for k, v := range map1 {
		fmt.Println(k, v)
	}

	// switch
	condition1 := 1
	switch condition1 {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("defalut")
	}

	switch condition2 := 100; condition2 {
	case 100:
		fmt.Println("100")
	default:
		fmt.Println("default")
	}

	switch {
	case condition1 == 1:
		fmt.Println("1")
	default:
		fmt.Println("default")
	}

	// ラベル付きfor
	// Loop:
	// 	for {
	// 		for {
	// 			for {
	// 				fmt.Println("START")
	// 				break
	// 			}
	// 			fmt.Println("実行したくない")
	// 		}
	// 		fmt.Println("実行したくない")
	// 	}
	// 	fmt.Println("END")
	Loop:
	for i := 0; i < 3; i++ {
		for j := 1; j < 3; j++ {
			if j > 1 {
				continue Loop // jが3を超えるまで、一番はじめのforに戻る
			}
			fmt.Println(i, j, i*j)
		}
		fmt.Println("実行したくない")
	}

	// TestDefer()
	// defer func() {
	// 	fmt.Println("1")
	// 	fmt.Println("2")
	// 	fmt.Println("3")
	// }()
	// RunDefer()

	file, fileError := os.Create("test.txt")
	if fileError != nil {
		fmt.Println("fileError", fileError)
	}
	defer file.Close()
	file.Write([]byte("Hello"))

	// panic
	// defer func() {
	// 	if x:= recover(); x != nil {
	// 		fmt.Println(x)
	// 	}
	// }()
	// panic("runtime error")
	// fmt.Println("start")

	// goroutin
	// go sub()
	// for {
	// 	fmt.Println("main loop")
	// 	time.Sleep(200 * time.Millisecond)
	// }

	// make
	slice1 := []int{100, 200}
	fmt.Println(slice1)

	slice1 = append(slice1, 1000, 2000)
	fmt.Println(slice1)

	slice2 := make([]int, 5, 10)
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	// 容量を超えた数をappendしたりすると、自動的的に初期の容量の倍が確保されてしまう
	slice2 = append(slice2, 1,2,3,4,5,6,7,8)
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	// copy
	slice3 := []int{100, 200}
	slice2 = slice3

	// 参照渡しになるので、slice3も書き換わってしまう
	slice2[0] = 999
	fmt.Println(slice3)
//
	slice4 := []int{1,2,3,4,5}
	slice2 = make([]int, 5, 10)
	n := copy(slice2, slice4)
	fmt.Println(n, slice4)

}
// 並行処理
func sub() {
	for {
		fmt.Println("sub loop")
		time.Sleep(100 * time.Millisecond)
	}
}

// defer
func TestDefer() {
	defer fmt.Println("EMD")
	fmt.Println("START")
}

func RunDefer() {
	defer fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")
}

func Closure() func(string)string {
	var store string
	return func(next string)string {
		stringValue := store
		store = next
		return stringValue
	}
}

// 関数を引数にとる関数
func CallbackFunc(f func()) {
	f()
}

// 関数を返す関数
func ReturnFunc()func() {
	return func() {
		fmt.Println("im function!!")
	}
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

// ジェネレーター
func intergers() func()int {
	i := 0
	return func()int {
		i++
		return i
	}
}
