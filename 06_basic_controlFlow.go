package main

import "fmt"

func main06() {

	/*
		if 條件{
			條件成立動作
		}else if{
			條件成立動作
		}else{
			其餘條件成立動作
		}
	*/
	if score := 80; score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score >= 70 {
		fmt.Println("C")
	} else if score >= 60 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}
	fmt.Println("----")
	/*
		for是go唯一迴圈!
		->while必須用for模擬
	*/
	for i := 0; i < 5; i++ {
		println(i)
	}
	fmt.Println("----")
	/*
		使用for模擬while
	*/
	/*
		for a := 0; a < 5; {
			fmt.Println("a<5")
		}
	*/
	/*
		無限迴圈
	*/
	/*
		for {
			fmt.Println("loop")
		}
		1.當不知道何時結束時，就會用
		ex.網路取資料

	*/
	/*
				可搭配break continur使用
				for {
		    		data, err := conn.Read()

		    		if err != nil {
		        	break
		    		}

		    		process(data)
					}
	*/
	g := 0
	for {
		if g > 10 {
			break
		}
		fmt.Println(g)
		g++
	}
	fmt.Println("----")

	for h := 0; h < 10; h++ {
		if h == 2 {
			continue
		}
		fmt.Println(h)
	}

	fmt.Println("----")

	/*
		遇到 array、slice、map、string
		可用range走訪資料
		for index, value := range nums {
		}
		不需要index可用_替代
		for _, value := range nums {
		}
	*/
	nums := []int{10, 20, 30}
	for index, value := range nums {
		fmt.Printf("index=%d,value=%d\n", index, value)
	}

	fmt.Println("----")
	/*
		switch基本寫法
	*/
	day := 3
	switch day {
	case 1:
		fmt.Println("Mon")
	case 2:
		fmt.Println("Tue")
	case 3:
		fmt.Println("Wed")
	case 4:
		fmt.Println("Thu")
	case 5:
		fmt.Println("Fri")

	}
	fmt.Println("----")
	/*
		sitch可不需要變數,相當於if else
	*/
	score := 95
	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >= 70:
		fmt.Println("C")
	case score >= 60:
		fmt.Println("D")
	default:
		fmt.Println("F")
	}
	fmt.Println("----")
	/*
		判斷式可有多個值
	*/
	grade := "B"
	switch grade {
	case "A", "B":
		fmt.Println("High Level")
	case "C":
		fmt.Println("Middle Level")
	}
	fmt.Println("----")
	/*
		如果想繼續執行下一個 case，可以使用 fallthrough
	*/
	x := 1

	switch x {
	case 1:
		fmt.Println("One")
		fallthrough
	case 2:
		fmt.Println("Two")
	}
	fmt.Println("----")

}

//06 Flow Control 語法通式整理

/*if 寫法

if 宣告變數(可選); 條件 {
	條件成立時執行
} else if 條件 {
	前面條件不成立，且此條件成立時執行
} else {
	以上條件都不成立時執行
}

注意：
- 宣告變數可選
- 條件必須是 bool
- Go 的 if 不需要小括號
- { 必須跟 if 在同一行
*/

/*for 寫法

for 初始化; 條件; 更新 {
	重複執行的內容
}

例如：
for i := 0; i < 5; i++ {
	fmt.Println(i)
}
*/

/*for 模擬 while 寫法

初始化變數

for 條件 {
	重複執行的內容
	更新變數
}

例如：
i := 0
for i < 5 {
	fmt.Println(i)
	i++
}
*/

/* for無限迴圈寫法

for {
	重複執行的內容

	if 結束條件 {
		break
	}
}

常用於：
- 不知道何時結束的流程
- 讀取資料
- Server 等待連線
*/

/*break 寫法

for {
	if 結束條件 {
		break
	}
}

用途：
- 直接離開目前所在的 for
- 也可用在 switch
*/

/*continue 寫法

for 初始化; 條件; 更新 {
	if 跳過條件 {
		continue
	}

	本次迴圈要執行的內容
}

用途：
- 跳過本次迴圈
- 直接進入下一輪
*/

/*range 寫法

for index, value := range 資料 {
	使用 index 與 value
}

不需要 index：

for _, value := range 資料 {
	使用 value
}

不需要 value：

for index := range 資料 {
	使用 index
}

可用於：
- array
- slice
- map
- string
*/

/*switch value 寫法

switch 變數 {
case 值1:
	變數符合值1時執行
case 值2:
	變數符合值2時執行
default:
	都不符合時執行
}

注意：
- Go 的 switch 不需要 break
- 執行完符合的 case 後會自動離開 switch
*/

/*switch expression 寫法

switch {
case 條件1:
	條件1成立時執行
case 條件2:
	條件2成立時執行
default:
	都不成立時執行
}

用途：
- 可取代 if...else if...else
*/

/*switch 多個 case 值寫法

switch 變數 {
case 值1, 值2, 值3:
	符合其中一個值就執行
case 值4:
	符合值4時執行
default:
	都不符合時執行
}
*/

/*fallthrough 寫法

switch 變數 {
case 值1:
	執行 case 值1
	fallthrough
case 值2:
	繼續執行 case 值2
}

注意：
- fallthrough 會直接執行下一個 case
- 不會重新判斷下一個 case 條件
- 實務上不常用
*/
