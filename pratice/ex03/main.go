package main

/*
## 題組 3：成績分類與流程追蹤（06、07）

**目標：** 練習 if、switch、迴圈、continue、break 與 early return。

**新版規則：** 可自行設計成績、溫度、遊戲排名或風險等級。分類名稱與測試資料可替換；核心是完成範圍驗證、if 與 switch 兩種分類、迴圈統計、continue、break 與多值 case。兩種分類函式請使用不同名稱，例如 `GradeByIf`、`GradeBySwitch`。

要求：

1. 實作 `Grade(score int) string`：不在 0～100 回傳 `invalid`；90 以上 A、80～89 B、70～79 C、60～69 D、其餘 F。先用 if，再改寫成不帶運算式的 switch。
2. 處理 `[]int{95, -1, 82, 60, 59, 101, 100}`，用 `continue` 跳過非法分數，印出每筆有效分數與等級，統計有效人數及及格人數。
3. 用只有條件的 `for` 計算 1～10 的總和。
4. 用無限 `for`，累加 1、2、3……，總和首次達到或超過 20 時 `break`，印出最後加入的數字與總和。
5. 使用一個 case 多值，將 A、B 分為 `high`，C、D 分為 `pass`，F 分為 `fail`，其他為 `invalid`。

參考驗收：若使用題目提供的資料，結果應為有效人數 `5`、及格人數 `4`、1～10 總和 `55`，累加達到 20 時為 `21`。自訂資料時，請自行寫出並核對預期結果。

閱讀題：先預測以下輸出，再執行確認，解釋第二個 case 是否會重新判斷條件。

```go
score := 95
switch {
case score >= 90:
    fmt.Println("A")
    fallthrough
case score < 60:
    fmt.Println("F")
default:
    fmt.Println("other")
}
```

進階挑戰：寫一段 `goto End` 跳過中間輸出的合法範例，再改用 if 表達相同行為。說明為何 goto 不能跳過在目標位置仍有效的變數宣告。

執行方式（專案根目錄）：go run ./pratice/ex03
作答方式：在下方新增所需型別、函式與 import，並完成 main 的 TODO。
*/

// TODO: 在此定義本題需要的型別與函式。

func main() {
	// TODO: 實作 Grade 的 if 與 switch 版本
	// TODO: 跳過非法分數並統計人數
	// TODO: 完成條件迴圈、無限迴圈與多值 case
	// TODO: 預測 fallthrough 並完成 goto 挑戰
}
