package main

import "fmt"

func checkScore(score int) string {
	if score < 0 || score > 100 {
		return "invalid score"
	}

	if score >= 60 {
		return "pass"
	}

	return "fail"
}

func main07() {
	/*
		一般 break 只能跳出目前所在的最內層 for
		以下以搜尋矩陣（2D Array）做為使用範例
	*/
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {

			if matrix[i][j] == 5 {
				fmt.Printf("Found %d at (%d, %d)\n", matrix[i][j], i, j)
				break
			}

			fmt.Printf("Checking (%d,%d) = %d\n", i, j, matrix[i][j])
		}
	}

	fmt.Println("Search End")
	fmt.Println("----")
	/*
		使用label+break跳出頂層for
		1. Label 是替某個 for、switch 或 select 命名；break Label 就是直接結束那個被命名的控制結構，而不是跳到 Label 那一行
		2. Label可以是任意合法的go名稱
	*/
Search:
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {

			if matrix[i][j] == 5 {
				fmt.Printf("Found %d at (%d, %d)\n", matrix[i][j], i, j)
				break Search
			}

			fmt.Printf("Checking (%d,%d) = %d\n", i, j, matrix[i][j])
		}
	}

	fmt.Println("Search End")
	fmt.Println("----")

	/*
		label+goto可直接跳到指定 Label
	*/
	fmt.Println("\n== goto ==")

	x := 3

	if x < 5 {
		goto SmallNumber
	}

	fmt.Println("x >= 5")
	goto EndGoto

SmallNumber:
	fmt.Println("x < 5")

EndGoto:
	fmt.Println("goto section end")

	// 3. return / Early Return
	fmt.Println("\n== return / Early Return ==")

	fmt.Println(checkScore(85))
	fmt.Println(checkScore(40))
	fmt.Println(checkScore(120))
}

//07 Advanced Flow Control 語法通式整理

/*Label + break 寫法

LabelName:
for 初始化; 條件; 更新 {
	for 初始化; 條件; 更新 {
		if 結束條件 {
			break LabelName
		}
	}
}

用途：
- 跳出指定 Label 標記的 for / switch / select
- 常用於多層迴圈提前結束
*/

/*Label + continue 寫法

LabelName:
for 初始化; 條件; 更新 {
	for 初始化; 條件; 更新 {
		if 跳過條件 {
			continue LabelName
		}
	}
}

用途：
- 跳過 Label 標記迴圈的本次迭代
- 直接進入外層迴圈的下一輪
*/

/*goto 寫法

goto LabelName

// 中間程式會被略過

LabelName:
	跳到這裡後繼續執行

注意：
- goto 是直接跳到指定 Label
- 實務上不常用
- 不建議拿來取代 if / for
*/

/*goto 常見結構

if 條件 {
	goto LabelName
}

一般流程

LabelName:
	特殊流程或結束流程
*/

/*goto 限制

goto 不可以跳過變數宣告

錯誤範例：

goto End

x := 10

End:
fmt.Println(x)

原因：
- 如果允許跳過 x := 10
- 到 End 時 x 可能還沒有被建立
*/

/*fallthrough 寫法

switch 變數 {
case 值1:
	執行 case 值1
	fallthrough
case 值2:
	繼續執行 case 值2
default:
	都不符合時執行
}

注意：
- fallthrough 只能用在 switch 的 case 裡
- fallthrough 會直接執行下一個 case
- 不會重新判斷下一個 case 條件
*/

/*return 寫法

func 函式名稱(參數) 回傳型別 {
	if 結束條件 {
		return 回傳值
	}

	return 回傳值
}

用途：
- 結束目前 function
- 回傳結果
*/

/*Early Return 寫法

func 函式名稱(參數) 回傳型別 {
	if 錯誤條件 {
		return 錯誤結果
	}

	if 特殊條件 {
		return 特殊結果
	}

	return 正常結果
}

用途：
- 先處理錯誤或特殊情況
- 避免 if else 巢狀太深
- Go 很常用這種寫法
*/
