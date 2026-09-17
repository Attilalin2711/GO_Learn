package main

/*
## 題組 5：安全除法與工作清理（08）

**目標：** 區分一般錯誤、延後清理與 panic 復原。

**新版規則：** 可改成付款、檔案、任務或其他可能失敗的操作。核心是一般可預期錯誤使用 error、資源清理使用 defer、真正不可恢復的流程才示範 panic/recover；至少測試成功與失敗各一次。

要求：

1. 實作 `SafeDivide(a, b int) (int, error)`，除數為 0 時回傳非 nil error，否則回傳整數商與 nil。
2. 呼叫端分別處理 `10/2`、`10/0`、`-9/2`，錯誤時不可把結果當成成功值使用。
3. 實作 `RunJob(fail bool) error`：先註冊印出 `cleanup A` 的 defer，再註冊 `cleanup B`，接著印出 `start`。fail 為 true 時回傳 error，否則印出 `done` 並回傳 nil。
4. 實作 `RunProtectedJob()`，在同一個函式的 deferred function 中直接呼叫 recover；印出 `before panic` 後觸發 `panic("job failed")`，捕捉後印出錯誤。呼叫端再印出 `next job`。
5. 實作 `ReadText(path string) (string, error)`：用 `os.Open` 開檔、成功後立即註冊 `defer file.Close()`，使用 `io.ReadAll` 讀取並檢查 error。

驗收：

- 除法結果為 `5`、非 nil error、`-4`。
- RunJob(false) 順序為 `start → done → cleanup B → cleanup A`。
- RunJob(true) 仍執行兩個 cleanup，且沒有 `done`。
- panic 後的同函式一般敘述不會繼續執行，但呼叫端會印出 `next job`。
- 內容為 `hello` 的檔案回傳 `hello`；不存在的路徑回傳非 nil error。

思考：除以零的輸入檢查為什麼適合回傳 error？recover 能不能直接捕捉另一個 goroutine 的 panic？

執行方式（專案根目錄）：go run ./pratice/ex05
作答方式：在下方新增所需型別、函式與 import，並完成 main 的 TODO。
*/

// TODO: 在此定義本題需要的型別與函式。

func main() {
	// TODO: 實作 SafeDivide 並處理成功與錯誤
	// TODO: 實作 RunJob 觀察 defer 順序
	// TODO: 實作 RunProtectedJob 並觀察呼叫端流程
	// TODO: 實作 ReadText 並測試正常與不存在路徑
}
