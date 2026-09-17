package main

/*
## 題組 7：商品庫存與資料副本（10、11）

**目標：** 分辨陣列複製、slice 共用資料與 map 的查詢更新。

**新版規則：** 可使用庫存、成績、購物車或自訂資料。數值與 key 不限；核心是證明陣列賦值會複製、slice 可能共用底層陣列、copy 可建立獨立副本，並正確處理 map 的新增、查詢、刪除與錯誤條件。

第一部分：先預測每一步資料，再寫程式驗證。

1. 建立 `original := [3]int{10, 20, 30}`，以賦值複製成 copied，修改 `copied[0] = 99`。
2. 建立 `selected := original[1:]`，修改 `selected[0] = 88`，印出 original、selected、len、cap。
3. 用 make 與 copy 建立 selected 的獨立副本，修改副本後確認 selected 不變。
4. 建立長度 0、容量 2 的 slice，append 三筆資料，逐次記錄 len 與 cap；不指定擴容後的精確容量。

驗收：步驟 1 後 original 仍為 `[10 20 30]`；步驟 2 後為 `[10 88 30]`，selected 的 len、cap 都為 2；最後 append 的 slice 長度為 3、容量至少為 3。

第二部分：用 `map[string]int{"apple": 10, "banana": 0}` 管理庫存。

1. 實作 `Restock(stock map[string]int, name string, amount int) error`，只接受非空名稱、正數數量與已初始化 map；新品可直接新增。
2. 實作 `Sell(stock map[string]int, name string, amount int) error`，拒絕不存在商品、非正數數量、庫存不足；失敗時不得修改庫存。
3. 使用 comma-ok 區分 banana 的零庫存與 orange 不存在。
4. 刪除 banana，並將剩餘商品 key 排序後列印。

驗收：apple 補貨 5、售出 3 後為 `12`；再售出 20 回傳錯誤且仍為 `12`。查詢 banana 初始為 `0, true`，orange 為 `0, false`，刪除後 banana 為 `0, false`。

思考：讀取 nil map 與寫入 nil map 的結果有何不同？將 map 傳入函式後修改項目，為何呼叫端能看到變化？

執行方式（專案根目錄）：go run ./pratice/ex07
作答方式：在下方新增所需型別、函式與 import，並完成 main 的 TODO。
*/

// TODO: 在此定義本題需要的型別與函式。

func main() {
	// TODO: 比較陣列複製、slice 共用與 copy
	// TODO: 觀察 append 的長度與容量
	// TODO: 實作 Restock、Sell 及錯誤處理
	// TODO: 區分零庫存與不存在商品並排序輸出
}
