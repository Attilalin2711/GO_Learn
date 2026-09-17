package main

/*
## 題組 2：圖書館借閱管理（01～05）

目標：整合函式、指標、struct、method、const、iota、可變參數與多回傳值。

情境：
圖書館需要管理書籍的借閱狀態與可借天數。請建立 Book 型別及相關操作，
比較值傳遞和指標對原始資料的影響。

第一部分：狀態與資料模型

1. 宣告 `type BookStatus int`。
2. 使用 `const` 和 `iota` 建立四種狀態：Available、Borrowed、Reserved、Lost。
3. 為 BookStatus 實作 `String() string` method，分別回傳小寫狀態文字；
   未知的狀態值回傳 `unknown`。
4. 建立 Book struct，至少包含 Title string、Pages int、BorrowDays int、
   Status BookStatus。
5. 實作建構函式，接收各欄位資料並回傳 Book。

第二部分：值與指標

6. 實作 `ExtendByValue`，接收 Book 與增加天數。它會修改收到的副本，
   呼叫後原書籍的 BorrowDays 必須保持不變。
7. 實作 `ExtendByPointer`，接收 *Book 與增加天數。它必須修改原書籍；
   收到 nil、增加天數小於等於 0 時直接返回。
8. 為 Book 實作 value receiver 的 `PreviewBorrow`，嘗試把狀態改為 Borrowed。
9. 為 *Book 實作 pointer receiver 的 `Borrow`，真正把狀態改為 Borrowed；
   receiver 為 nil 時直接返回。
10. 在 main 中比較兩種延長方式及兩種 method 對原資料造成的差異。

第三部分：函式練習

11. 實作可變參數函式 `TotalPages(books ...Book) int`，計算總頁數；
    沒有參數時回傳 0。
12. 實作 `PagesPerDay(pages, days int) (int, bool)`；頁數為負或天數小於
    等於 0 時回傳 `0, false`，否則回傳整數除法結果及 true。
13. 實作具名回傳值函式，接收兩本書的頁數並回傳總和與頁數差。
14. 使用匿名函式計算逾期費用，例如「逾期天數 × 每日費用」。
15. 使用 `new(Book)` 建立另一本書，設定欄位後印出內容。

自行驗收：

- 四個狀態由 0 起依序遞增，任意未知狀態會得到 `unknown`。
- value 版本執行後，原本的 BorrowDays 與 Status 不變。
- pointer 版本執行後，原本的 BorrowDays 與 Status 改變。
- `ExtendByPointer(nil, 7)` 與 nil receiver 都不會 panic。
- `TotalPages()` 回傳 0；傳入多本書時回傳正確總頁數。
- PagesPerDay 至少測試一個合法案例與兩種非法案例。

進階挑戰：

- Book 加入 MaxBorrowDays，延長後不可超過上限。
- 將 Borrow 改成回傳 bool，已借出或遺失的書籍不得再次借閱。

思考題：

1. value receiver 為什麼無法改變原本的 Book？
2. 指標本身是否也會在傳入函式時被複製？
3. 為什麼操作指標前要檢查 nil？

執行方式（專案根目錄）：go run ./pratice/ex02
*/

// TODO: 在此定義 BookStatus、狀態常數、Book、methods 與輔助函式。

func main() {
	// TODO: 建立至少兩本書並印出初始狀態。

	// TODO: 比較 value 與 pointer 操作。

	// TODO: 驗證 nil 防護與非法增加天數。

	// TODO: 測試總頁數、每日頁數及具名回傳值函式。

	// TODO: 完成匿名函式與 new(Book) 練習。
}
