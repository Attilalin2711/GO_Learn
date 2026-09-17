package main

/*
## 題組 11：課程查詢 HTTP API（08、11、16、17）

**目標：** 把課程資料透過 HTTP 回傳，整合路由、查詢參數、狀態碼、JSON 與 client。

**新版規則：** API 資源可改成商品、書籍、電影或自訂資料。路徑、id 和內容可自行命名；核心是 GET 查詢、JSON 回應、400/404/405 狀態、Allow header、httptest server、client timeout 與關閉 response body。handler 不應修改作為資料來源的 map。

資料：以唯讀 `map[string]Course` 保存兩筆課程，key 為 `go-basic`、`go-concurrency`，分別為 `Go 入門 / 12 小時`、`Go 並行 / 8 小時`。

要求：

1. 使用 NewServeMux 註冊 `/course`，只接受 GET；其他 method 回傳 405 並設定 `Allow: GET`。
2. 讀取 query parameter `id`。缺少或空值回傳 400；找不到課程回傳 404；成功回傳 200 與 Course JSON。
3. 所有回應都設定 `Content-Type: application/json; charset=utf-8`，錯誤格式固定為 `{"error":"訊息"}`。請自行編碼 JSON 錯誤，不直接以純文字作為錯誤 body。
4. 在寫入 body 前設定 header 與狀態碼，並處理 JSON 編碼錯誤。
5. 使用 httptest.NewServer 啟動測試伺服器，註冊 defer Close。
6. 建立 Timeout 為 2 秒的 http.Client，呼叫 API，檢查 request error、status、讀取 body 的 error，並關閉每個 response.Body。連續請求可拆成輔助函式，使 defer 在每次請求處理結束時執行。

驗收案例：

| 請求 | 狀態碼 | 重點 |
| --- | --- | --- |
| GET /course?id=go-basic | 200 | title 為 Go 入門，hours 為 12 |
| GET /course?id=go-concurrency | 200 | title 為 Go 並行，hours 為 8 |
| GET /course | 400 | 可解碼的 error JSON |
| GET /course?id=missing | 404 | 可解碼的 error JSON |
| POST /course?id=go-basic | 405 | Allow 為 GET，body 是 error JSON |

思考：收到 HTTP 404 時，client 的 err 是否一定非 nil？設定 header 為什麼要早於寫入 body？

進階挑戰：加入 `GET /courses`，以課程 id 排序後回傳課程陣列；使用 httptest.NewRecorder 驗證 handler，不必啟動實際監聽埠。

執行方式（專案根目錄）：go run ./pratice/ex11
作答方式：在下方新增所需型別、函式與 import，並完成 main 的 TODO。
*/

// TODO: 在此定義本題需要的型別與函式。

func main() {
	// TODO: 在本目錄建立或沿用題組 10 的 Course 定義
	// TODO: 建立唯讀課程資料與 HTTP handler
	// TODO: 處理 200、400、404、405 及 JSON 回應
	// TODO: 使用 httptest server 與有 timeout 的 client
	// TODO: 逐一驗證請求並關閉 response.Body
}
