package main

import "fmt"

// 沒有回傳值的 function
func sayHello(name string) {
	fmt.Println("Hello,", name)
}

// 一般回傳值
func add(a, b int) int {
	return a + b
}

// 多回傳值
func divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}

	return a / b, true
}

// 不定長參數
func sum(nums ...int) int {
	total := 0

	for _, n := range nums {
		total += n
	}

	return total
}

// 命名回傳值-->宣告回傳值變數
func rectangle(width, height int) (area int) {
	area = width * height
	return
}

// 多個命名回傳值
func calc(a, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return
}

func main02() {
	sayHello("Leon")

	result := add(10, 5)
	fmt.Println("add =", result)

	divResult, ok := divide(10, 2)
	if ok {
		fmt.Println("divide =", divResult)
	} else {
		fmt.Println("cannot divide by zero")
	}

	fmt.Println("sum =", sum(1, 2, 3, 4, 5))

	area := rectangle(10, 5)
	fmt.Println("rectangle area =", area)

	s, d := calc(10, 3)
	fmt.Println("calc sum =", s)
	fmt.Println("calc diff =", d)

	// 匿名函式：把 function 存到變數裡
	multiply := func(a, b int) int {
		return a * b
	}

	fmt.Println("multiply =", multiply(3, 4))

	// 匿名函式：立即執行
	message := func(name string) string {
		return "Hi, " + name
	}("Leon")

	fmt.Println(message)
}

/*基本寫法
func 函式名稱(參數名稱 參數型別) 回傳型別 {

    // 函式內容

    return 回傳值

}
參數跟回傳值數量可自訂
*/

/*匿名函數
func (參數)回傳型別{
	return 回傳值
}
*/

/*不定長參數
func 函式名稱(參數名稱 ...參數型別) 回傳型別 {
}

如果有多個參數,不定長參數要放最後

func 函式名稱(參數名稱,參數名稱,參數名稱 ...參數型別) 回傳型別 {
}
*/
