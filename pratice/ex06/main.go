package main

/*
## 題組 6：動物照護中心（04、09）

**目標：** 使用介面統一處理不同型別，練習型別斷言與 type switch。

**新版規則：** 不必使用動物，可以改成交通工具、通知方式、付款方式或任意具有共同能力的型別。核心是定義介面、讓至少三種型別實作它、以介面 slice 操作，並示範安全型別斷言與 type switch。

要求：

1. 定義 `Animal` 介面，包含 `Speak() string`、`Move() string`。
2. 建立 Dog、Cat，皆有 Name 欄位並以 value receiver 實作介面。
3. 實作 `Describe(a Animal) string`，將 Speak 與 Move 結果組合成介紹。
4. 建立包含一隻 Dog 與一隻 Cat 的 `[]Animal` 並逐一介紹。
5. 使用 comma-ok 將第一個元素斷言成 Dog，再把 Cat 嘗試斷言成 Dog，失敗時印出提示。
6. 用 type switch 輸出照護需求：Dog 為 `散步`、Cat 為 `攀爬`、其他為 `一般照護`。新增 Bird 實作介面，確認走 default。

驗收：三種動物都能傳給 Describe；錯誤斷言不會 panic；新增 Bird 不需要修改 Describe。

進階挑戰：將 Dog 的 Move 改成 pointer receiver，觀察 `Dog{}` 與 `&Dog{}` 哪一個能指定給 Animal，解釋原因。

執行方式（專案根目錄）：go run ./pratice/ex06
作答方式：在下方新增所需型別、函式與 import，並完成 main 的 TODO。
*/

// TODO: 在此定義本題需要的型別與函式。

func main() {
	// TODO: 定義 Animal、Dog、Cat 與各方法
	// TODO: 實作 Describe 並逐一介紹
	// TODO: 用 comma-ok 安全斷言型別
	// TODO: 新增 Bird 並完成 type switch
}
