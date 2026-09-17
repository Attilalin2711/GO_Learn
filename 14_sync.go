package main

import (
	"fmt"
	"sync"
)

// 將資料與保護它的 mutex 放在一起，可降低漏加鎖的風險。
type SafeCounter struct {
	mu    sync.Mutex
	value int
}

func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock() // 即使提早 return 也會解鎖。
	c.value++
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main14() {
	var counter SafeCounter
	var wg sync.WaitGroup
	// 100 個 goroutine 同時累加；無鎖會有 data race 並可能遺失更新。
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}
	wg.Wait()
	fmt.Println("計數結果：", counter.Value()) // 必定是 100000。

	// WaitGroup 等工作完成；Mutex 保護共享資料。可用 go run -race 檢查競態。
}
