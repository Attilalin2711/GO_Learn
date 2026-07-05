package main

import "fmt"

func main() {
	if point := 90; point < 60 {
		fmt.Println("F")
	} else if point >= 60 && point < 70 {
		fmt.Println("D")
	} else if point >= 70 && point < 80 {
		fmt.Println("C")
	} else if point >= 80 && point < 90 {
		fmt.Println("B")
	} else {
		fmt.Println("A")
	}

}
