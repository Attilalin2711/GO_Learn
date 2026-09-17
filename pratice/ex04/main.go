package main

/*
## 題組 4：二維資料搜尋與批次檢查（07、10）

**目標：** 分辨一般 break、帶標籤 break、帶標籤 continue 與 return。

**新版規則：** 二維資料內容與搜尋目標可以自行設計，例如座位、地圖或棋盤。核心是實際比較一般 break、標籤 break、標籤 continue 與 early return 的控制範圍。

資料：

```go
matrix := [][]int{
    {1, 2, 3},
    {4, 5, 6},
    {5, 8, 9},
}
```

要求：

1. 用巢狀迴圈搜尋 `5`，先用一般 `break`，記錄找到的座標。
2. 改成 `break Search`，只印出第一個符合的座標。
3. 實作 `Find(matrix [][]int, target int) (row, col int, found bool)`，用 early return 回傳第一個符合位置；找不到回傳 `-1, -1, false`。
4. 處理批次 `[][]int{{80, 90}, {70, -1, 100}, {}, {60, 75}}`：空批次跳過，任何一筆不在 0～100 就用帶標籤的 continue 跳過整批。僅在整批檢查通過後印出總分。

參考驗收：使用題目資料時，一般 break 會找到兩個 `5`，標籤 break 與 Find 只取第一個。自訂資料時仍須測試找到、找不到、空資料、不同列長度，以及整批跳過非法資料。

進階挑戰：用不同長度的列與空列驗證搜尋不會索引越界。

執行方式（專案根目錄）：go run ./pratice/ex04
作答方式：在下方新增所需型別、函式與 import，並完成 main 的 TODO。
*/

// TODO: 在此定義本題需要的型別與函式。

func main() {
	// TODO: 建立矩陣並比較兩種 break
	// TODO: 實作 Find 與找不到的情境
	// TODO: 使用帶標籤 continue 檢查整批資料
}
