package main

import "fmt"

func main10() {
	// 陣列長度是型別的一部分；陣列賦值會複製全部元素。
	scores := [3]int{80, 90, 100}
	copied := scores
	copied[0] = 60
	fmt.Println("原陣列：", scores)
	fmt.Println("複製後：", copied)

	// slice 描述底層陣列的一段範圍，包含指標、長度與容量。
	selected := scores[1:]
	selected[0] = 95 // 共用底層陣列，所以 scores[1] 也被修改。
	fmt.Printf("切片：%v，len=%d，cap=%d\n", selected, len(selected), cap(selected))
	fmt.Println("被切片修改的陣列：", scores)

	// make 建立長度 0、容量 3 的切片；append 加入元素。
	numbers := make([]int, 0, 3)
	numbers = append(numbers, 10, 20, 30)
	fmt.Printf("append 後：%v，len=%d，cap=%d\n", numbers, len(numbers), cap(numbers))

	// 容量不足時 append 會配置新底層陣列，容量增長值不應寫死依賴。
	numbers = append(numbers, 40)
	fmt.Printf("擴容後：%v，len=%d，cap=%d\n", numbers, len(numbers), cap(numbers))

	// copy 可建立互不影響的副本；複製數量是兩邊長度的較小值。
	independent := make([]int, len(numbers))
	copy(independent, numbers)
	independent[0] = 999
	fmt.Println("原切片：", numbers)
	fmt.Println("獨立副本：", independent)
}
