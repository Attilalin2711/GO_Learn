package main

import "fmt"

func produce(out chan<- int) {
	// chan<- 表示只能傳送，限制權限也清楚表達意圖。
	for i := 1; i <= 5; i++ {
		out <- i
	}
	close(out) // 傳送端關閉，表示不會再有資料。
}

func square(in <-chan int, out chan<- int) {
	// range 會接收至 in 關閉且緩衝資料讀完。
	for n := range in {
		out <- n * n
	}
	close(out)
}

func main13() {
	numbers := make(chan int)    // 無緩衝：傳送與接收必須會合。
	squares := make(chan int, 2) // 有緩衝：可暫存兩筆資料。

	// pipeline：produce -> square -> main。
	go produce(numbers)
	go square(numbers, squares)
	for result := range squares {
		fmt.Println("平方：", result)
	}

	// 關閉且清空後會立即得到零值和 ok=false。
	value, ok := <-squares
	fmt.Printf("關閉後接收：value=%d, ok=%t\n", value, ok)
}
