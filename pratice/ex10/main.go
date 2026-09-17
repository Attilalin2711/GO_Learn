package main

/*
## 題組 10：課程資料的 JSON 匯入匯出（04、08、10、16）

**目標：** 練習 struct tag、序列化、反序列化及資料驗證。

**新版規則：** JSON 模型可改成商品、電影、玩家或任意資料。核心是匯出欄位、JSON tag、omitempty、未匯出欄位、Marshal/Unmarshal、輸入驗證，以及觀察 `map[string]any` 的動態型別。

要求：

1. 建立 Course：Title string、Hours int、Tags []string、Teacher string，以及未匯出的 internal string。
2. JSON 欄位名稱使用 `title`、`hours`、`tags`、`teacher`；teacher 加上 omitempty。
3. 實作 `EncodeCourse(c Course) ([]byte, error)`，使用 MarshalIndent。
4. 實作 `DecodeCourse(data []byte) (Course, error)`，使用 Unmarshal；再驗證 Title 不可為空、Hours 必須大於 0。
5. 將 `{"title":"Go 入門","hours":12,"tags":["Go","後端"]}` 解碼後再編碼。
6. 另解碼 `{"students":25,"active":true}` 到 `map[string]any`，印出兩個值的型別。

驗收：

- 正常資料得到 Title 為 `Go 入門`、Hours 為 12、兩個 Tags。
- Teacher 為空時 JSON 沒有 teacher；設為 `Leon` 後會出現。
- internal 即使有值也不出現在 JSON 中。
- 缺少結尾括號、hours 為字串、title 空字串、hours 為 0，均回傳非 nil error。
- 動態解碼的 students 型別為 `float64`，active 為 `bool`。

進階挑戰：比較 Tags 為 nil slice 與空 slice 時的 JSON；使用 Decoder 的 DisallowUnknownFields 拒絕未知欄位。

執行方式（專案根目錄）：go run ./pratice/ex10
作答方式：在下方新增所需型別、函式與 import，並完成 main 的 TODO。
*/

// TODO: 在此定義本題需要的型別與函式。

func main() {
	// TODO: 建立 Course 與 JSON tags
	// TODO: 實作 EncodeCourse、DecodeCourse 及資料驗證
	// TODO: 驗證 omitempty、未匯出欄位及非法 JSON
	// TODO: 解碼動態 map 並觀察型別
}
