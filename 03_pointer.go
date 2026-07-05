package main

import "fmt"

// 值傳遞：只會改到複製品
func changeByValue(x int) {
	x = 100
}

// 指針傳遞：可以改到原本的變數
func changeByPointer(x *int) {
	*x = 100
}

func main03() {
	a := 10

	fmt.Println("原本的 a:", a)

	// 取得 a 的記憶體位址
	p := &a

	fmt.Println("a 的地址:", p)
	fmt.Println("p 指向的值:", *p)

	// 透過指針修改 a(取值後修改)
	*p = 20
	fmt.Println("透過 *p 修改後的 a:", a)

	// 值傳遞，不會改到原本的 a
	changeByValue(a)
	fmt.Println("changeByValue 後的 a:", a)

	// 指針傳遞，會改到原本的 a
	changeByPointer(&a)
	fmt.Println("changeByPointer 後的 a:", a)

	/*
		使用 new 建立一個 int 指針
		q := new(int)等價於
		var a int
		q = &a
	*/
	q := new(int)

	fmt.Println("new(int) 建立後 q 指向的值:", *q)

	*q = 200

	fmt.Println("修改 q 指向的值後:", *q)
}
