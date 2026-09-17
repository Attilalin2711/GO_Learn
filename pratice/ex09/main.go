package main

/*
## 題組 9：等待結果、逾時與非阻塞操作（13、15）

**目標：** 用 select 處理多來源結果，並讓逾時後的工作能正常結束。

**新版規則：** 可模擬 API、感測器、賽跑或其他多來源工作。訊息及延遲時間可自行設定；核心是 select 等待多來源、timeout、default 非阻塞操作，以及逾時後不遺留永遠阻塞的 goroutine。

要求：

1. 實作 `WaitResult(fast, slow <-chan string, timeout time.Duration) (string, bool)`，用 select 等待任一結果或 `time.After(timeout)`；收到結果回傳訊息與 true，逾時回傳空字串與 false。
2. 每個來源本題只傳一筆，不須 close。兩個來源各使用容量 1 的 channel，避免函式先返回後，尚未被讀取的傳送永久阻塞。
3. 情境 A：fast 預先放入 `fast done`，slow 保持空且不關閉，timeout 設為 1 秒。
4. 情境 B：兩個來源皆保持空且不關閉，timeout 設為 20 毫秒。
5. 情境 C：兩個 goroutine 分別延遲一段時間再傳送，呼叫 WaitResult，並用 WaitGroup 確認兩個工作者最後都能退出。不以毫秒級執行先後作為唯一驗收標準。
6. 另建立容量 1 的 queue，使用 `select/default` 嘗試接收，依序驗證空佇列、有一筆 42、讀完後又空的行為。

驗收：A 為 `fast done, true`；B 為 `"", false`；非阻塞讀取依序為 `目前沒有資料`、`42`、`目前沒有資料`。若兩個來源同時就緒，不要求 select 優先選哪一個。

進階挑戰：改成無緩衝結果 channel，增加 `done` channel，讓每個工作者用 select 在「傳送結果」與「收到停止通知」之間選擇。主程式返回時關閉 done，並確認所有工作者退出。

執行方式（專案根目錄）：go run ./pratice/ex09
作答方式：在下方新增所需型別、函式與 import，並完成 main 的 TODO。
*/

// TODO: 在此定義本題需要的型別與函式。

func main() {
	// TODO: 實作 WaitResult
	// TODO: 驗證立即結果、逾時與延遲工作情境
	// TODO: 以 WaitGroup 確認工作者退出
	// TODO: 使用 select/default 完成非阻塞接收
	// TODO: 進階：以 done 通知無緩衝傳送端退出
}
