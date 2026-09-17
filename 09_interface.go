package main

import "fmt"

// Animal 定義行為。Go 不必寫 implements；擁有全部方法就自動實作介面。
type Animal interface {
	Speak() string
	Move() string
}

type Dog struct{ Name string }
type Cat struct{ Name string }

func (d Dog) Speak() string { return "汪汪！" }
func (d Dog) Move() string  { return d.Name + " 正在跑步" }
func (c Cat) Speak() string { return "喵～" }
func (c Cat) Move() string  { return c.Name + " 輕巧地走路" }

// Introduce 依賴共同能力而非具體型別，這就是介面的多型。
func Introduce(a Animal) {
	fmt.Printf("叫聲：%s；動作：%s\n", a.Speak(), a.Move())
}

func main9() {
	animals := []Animal{Dog{Name: "Lucky"}, Cat{Name: "Chichi"}}
	for _, animal := range animals {
		Introduce(animal)
	}

	// comma-ok 型別斷言在型別不符時不會 panic。
	if dog, ok := animals[0].(Dog); ok {
		fmt.Println("第一隻動物是狗：", dog.Name)
	}

	// type switch 適合判斷多種底層具體型別。
	for _, animal := range animals {
		switch v := animal.(type) {
		case Dog:
			fmt.Println(v.Name, "需要每天散步")
		case Cat:
			fmt.Println(v.Name, "喜歡待在高處")
		}
	}
}
