package main

import "fmt"

type Animal interface {
	speak()
	move()
}

type Dog struct {
	name string
}
type Cat struct {
	name string
}

type butterfly struct {
	name string
}

func (d Dog) speak() {
	fmt.Println("WOOF!")
}

func (d Dog) move() {
	fmt.Println("run")
}

func (d Cat) speak() {
	fmt.Println("Meow")
}

func (d Cat) move() {
	fmt.Println("run")
}

func (d butterfly) move() {
	fmt.Println("flap")
}

func CallyourPet(p Animal) {
	p.speak()
	p.move()
}

func main() {
	lucky := Dog{name: "lucky"}
	chichi := Cat{name: "chichi"}
	bebe := butterfly{name: "bebe"}
	_ = bebe
	CallyourPet(lucky)
	CallyourPet(chichi)
	//CallyourPet(bebe)

}
