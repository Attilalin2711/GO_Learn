package main

import "fmt"

// 定義一個結構體
type Person struct {
	Name string
	Age  int
}

/*
1.Go 沒有 class，但 struct 可以receiver綁定 method。
2.receiver通常放在 func name前面,它決定這個方法屬於哪個型別,通常有以下特性:
	a.可自訂名稱,名稱只是在function內部使用的變數名稱
	b.可使用value receiver：收到的是副本，改不到原本的 Person
	c.可使用pointer receiver:指向原本物件的指針,可以改到原本的
*/
func (p Person) FakeBirthday() {
	p.Age++
	fmt.Println("FakeBirthday 裡面的 Age:", p.Age)
}

// pointer receiver：收到的是位址，可以修改原本的 Person
func (p *Person) RealBirthday() {
	p.Age++
	fmt.Println("RealBirthday 裡面的 Age:", p.Age)
}

// 一般函式：用指針修改 struct
func ChangeName(p *Person, newName string) {
	p.Name = newName
}

func main04() {
	// 建立一個 Person
	person := Person{
		Name: "Leon",
		Age:  30,
	}

	fmt.Println("初始資料:", person)

	// 呼叫 value receiver
	person.FakeBirthday()
	fmt.Println("FakeBirthday 後:", person)

	// 呼叫 pointer receiver
	person.RealBirthday()
	fmt.Println("RealBirthday 後:", person)

	// 把 person 的位址傳進函式
	ChangeName(&person, "Lin")

	fmt.Println("ChangeName 後:", person)
}

/*
其實go語言的完整函式結構
func (可選的receiver) 函式名稱(輸入參數) (輸出結果) {
    函式內容
}

一般 function 沒有 receiver，method 才有 receiver

*/
