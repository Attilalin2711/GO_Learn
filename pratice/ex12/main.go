package main

/*
## 最後綜合挑戰：並行課程查詢報表

完成題組 8～11 後再做。沿用課程 API，同時查詢 `go-basic`、`go-concurrency`、`missing`，整合以下要求：

**新版規則：** 可沿用自己在題組 11 設計的任何 API。查詢 id、報表欄位與資源名稱都可自訂；核心是每個請求由 goroutine 執行、以 channel 回傳結果、由協調者關閉 channel、區分成功與失敗、固定排序並輸出 JSON 報表。

1. 每個查詢由一個 goroutine 執行，透過 channel 回傳包含 ID、Course、Err 的結果。
2. 用 WaitGroup 等待工作者，統一關閉結果 channel；主程式用 range 彙整。
3. 將非 200 回應轉成可辨識的錯誤，錯誤不得被當成零值課程納入統計。
4. 由接收端統一建立成功清單與失敗清單，兩份清單皆按 ID 排序，再輸出 JSON 報表。
5. 每次 HTTP 請求都使用有限 timeout 並關閉 response.Body。

驗收：成功 2 筆、失敗 1 筆，成功課程總時數 `20`；missing 列於失敗清單。即使各次回應的完成順序不同，最終清單順序仍固定，程式能正常結束。

自我檢查：能否說明每個 channel 由誰關閉、每個 goroutine 如何退出，以及哪些資料只有單一 goroutine 會修改？

執行方式（專案根目錄）：go run ./pratice/ex12
作答方式：在下方新增所需型別、函式與 import，並完成 main 的 TODO。
*/

// TODO: 在此定義本題需要的型別與函式。

func main() {
	// TODO: 將題組 11 所需型別與服務函式整理到本目錄，僅保留一個 main
	// TODO: 建立含 ID、Course、Err 的結果型別
	// TODO: 並行查詢三個課程並統一關閉 channel
	// TODO: 分類成功失敗、排序並輸出 JSON 報表
	// TODO: 驗證總時數、timeout 與資源清理
}
