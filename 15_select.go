package main

import (
	"fmt"
	"time"
)

func main15() {
	fast := make(chan string)
	slow := make(chan string)
	go func() {
		time.Sleep(20 * time.Millisecond)
		fast <- "快速任務完成"
	}()
	go func() {
		time.Sleep(60 * time.Millisecond)
		slow <- "慢速任務完成"
	}()

	// select 等待多個 channel；兩個 case 同時可執行時會擇一。
	for i := 0; i < 2; i++ {
		select {
		case message := <-fast:
			fmt.Println(message)
		case message := <-slow:
			fmt.Println(message)
		case <-time.After(200 * time.Millisecond):
			fmt.Println("等待逾時")
			return
		}
	}

	// default 使 select 在沒有 channel 就緒時立刻繼續（非阻塞操作）。
	queue := make(chan int, 1)
	queue <- 42
	select {
	case value := <-queue:
		fmt.Println("立即收到：", value)
	default:
		fmt.Println("目前沒有資料")
	}
}
