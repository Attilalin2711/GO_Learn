package main

import (
	"encoding/json"
	"fmt"
)

// tag 設定 JSON 名稱；omitempty 在零值時省略欄位。只有匯出欄位會被處理。
type Course struct {
	Title    string   `json:"title"`
	Hours    int      `json:"hours"`
	Tags     []string `json:"tags"`
	Teacher  string   `json:"teacher,omitempty"`
	internal string
}

func main16() {
	course := Course{Title: "Go 入門", Hours: 12, Tags: []string{"Go", "後端"}, internal: "內部用"}
	// MarshalIndent 把 Go 值編碼成易讀 JSON。
	data, err := json.MarshalIndent(course, "", "  ")
	if err != nil {
		fmt.Println("編碼失敗：", err)
		return
	}
	fmt.Println(string(data))

	input := `{"title":"Go 並行程式設計","hours":8,"tags":["goroutine","channel"]}`
	var decoded Course
	// 必須傳指標，Unmarshal 才能寫入 decoded。
	if err := json.Unmarshal([]byte(input), &decoded); err != nil {
		fmt.Println("解碼失敗：", err)
		return
	}
	fmt.Printf("解碼結果：%+v\n", decoded)

	// 不確定結構可用 map；JSON 數字預設為 float64，已知格式時應優先用 struct。
	var dynamic map[string]any
	if err := json.Unmarshal([]byte(`{"active":true,"students":25}`), &dynamic); err != nil {
		fmt.Println("動態解碼失敗：", err)
		return
	}
	fmt.Printf("動態資料：active=%v, students=%v (%T)\n",
		dynamic["active"], dynamic["students"], dynamic["students"])
}
