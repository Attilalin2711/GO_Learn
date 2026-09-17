package main

import (
	"fmt"
	"sync"
	"time"
)

func main12() {
	// goroutine 是 Go runtime 管理的輕量並行工作單位。
	var wg sync.WaitGroup
	for workerID := 1; workerID <= 3; workerID++ {
		wg.Add(1) // 啟動前登記，避免主程式過早結束等待。
		go func(id int) {
			defer wg.Done()
			for job := 1; job <= 2; job++ {
				fmt.Printf("工作者 %d 處理任務 %d\n", id, job)
				time.Sleep(20 * time.Millisecond)
			}
		}(workerID) // 明確傳入當次值。
	}
	wg.Wait()
	fmt.Println("全部工作完成")

	// 每次輸出順序可能不同。若 goroutine 共享可變資料，還需同步機制。
}
