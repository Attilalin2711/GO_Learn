package main

import (
	"fmt"
	"sort"
)

func main11() {
	// map 儲存 key-value，key 必須是可比較的型別。
	prices := map[string]int{"蘋果": 30, "香蕉": 20}
	prices["芭樂"] = 40 // 新增
	prices["蘋果"] = 35 // 更新

	// comma-ok 可區分「不存在」與「存在但值是零值」。
	if price, exists := prices["香蕉"]; exists {
		fmt.Println("香蕉價格：", price)
	}
	if _, exists := prices["西瓜"]; !exists {
		fmt.Println("沒有西瓜的價格資料")
	}
	delete(prices, "香蕉")

	// map 迭代順序沒有保證；需要穩定輸出時先排序 key。
	keys := make([]string, 0, len(prices))
	for name := range prices {
		keys = append(keys, name)
	}
	sort.Strings(keys)
	for _, name := range keys {
		fmt.Printf("%s：%d 元\n", name, prices[name])
	}

	// map 是參考型別，函式內的修改可由呼叫端看到。
	applyDiscount(prices, 5)
	fmt.Println("折扣後：", prices)
}

func applyDiscount(prices map[string]int, amount int) {
	for name, price := range prices {
		prices[name] = price - amount
	}
}
