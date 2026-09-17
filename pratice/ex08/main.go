package main

/*
## 題組 8：並行工作與安全計數（12～14）

**目標：** 分清楚等待完成、保護共享資料與傳遞資料的責任。

**新版規則：** 工作者可以處理下載、訂單、計算或自訂任務。工作數量可變，但必須能由輸入推算正確總數；核心是 WaitGroup 等待、Mutex 保護共享狀態、channel 傳遞資料，以及由正確的一方關閉 channel。

第一部分：工作者與計數器。

1. 啟動 3 個 goroutine，每個印出自己的編號及兩筆任務。
2. 使用 WaitGroup，在啟動 goroutine 前 Add，在工作函式內 defer Done，全部結束後印出 `全部完成`。
3. 建立包含 Mutex 與 value 的 SafeCounter，以 pointer receiver 實作 Increment、Value，讀寫都使用同一把鎖。
4. 啟動 10 個 goroutine，每個遞增 1000 次；等待後讀取總值。

驗收：共 6 筆工作訊息，`全部完成` 最後出現；計數必須為 `10000`。不可用固定 Sleep 代替等待完成，也不要求不同 goroutine 的輸出順序固定。

第二部分：平方資料管線。

1. 產生器將 1～5 傳到 channel，再關閉它。
2. 平方處理函式用 `range` 讀取輸入並輸出平方，輸入結束後關閉輸出。
3. 函式參數使用 `chan<- int` 與 `<-chan int` 表達方向。
4. 主程式讀取結果並計算總和；分別用無緩衝與容量 2 的 channel 執行。
5. 讀完已關閉的輸出 channel 後，再用 comma-ok 接收一次並列印。

驗收：結果依序為 `1, 4, 9, 16, 25`，總和 `55`；最後接收為 `0, false`，兩種容量都能正常結束。

進階挑戰：改成 3 個平方工作者共用輸入與輸出。使用 WaitGroup，由統一的協調者在所有工作者完成後關閉輸出。驗收五個結果各出現一次、總和仍為 55，順序不限定。

檢查方式：環境支援 race detector 時，可執行 `go run -race ./pratice/ex08`。思考：只有 WaitGroup 為什麼不足以保護 `value++`？

執行方式（專案根目錄）：go run ./pratice/ex08
作答方式：在下方新增所需型別、函式與 import，並完成 main 的 TODO。
*/

// TODO: 在此定義本題需要的型別與函式。

func main() {
	// TODO: 啟動工作者並使用 WaitGroup 等待
	// TODO: 建立 SafeCounter 並驗證計數
	// TODO: 實作單向 channel 的產生器與平方管線
	// TODO: 比較 channel 容量並觀察關閉後接收
	// TODO: 進階：增加多個工作者並統一關閉輸出
}
