package main

import "fmt"

func main01() {
	//明確宣告
	var age int = 30
	//自動判別
	var name = "Leon"
	//短變數宣告-->宣告一個新變數,讓GO自動判別型別,只能在function內使用
	height := 175.5
	//印出不換行
	fmt.Print(name)
	//印出自動換行
	fmt.Println(age)
	//格式化輸出
	fmt.Printf("%s age is %d,height is %f\n", name, age, height)
	//GO編譯時會自動加;  if跟{需要放在同行
	if age >= 18 {
		fmt.Println("adult")
	} else {
		fmt.Println("minor")
	}

	for i := 0; i < 5; i++ {
		fmt.Println("i=", i)
	}

	score := 85

	//C++ 的 switch 主要是拿來比對固定值,Go的switch 可以不指定變數，直接判斷多個條件
	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >= 70:
		fmt.Println("C")
	default:
		fmt.Println("D")
	}
}

/*
重點
1.變數宣告方式與輸出方式
2.基礎型別
3.基礎語法if,for,switch
*/
