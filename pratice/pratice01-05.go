/*
綜合練習題：員工狀態管理系統

請手寫一個 Go 程式，完成以下需求。

題目需求

請建立一個 Employee 結構體，表示員工資料。

1. 宣告常數與狀態

請用 type、const、iota 建立員工狀態：

type Status int

狀態包含：

Idle
Working
Vacation
Resigned

並寫一個 method：

func (s Status) String() string

需求：

狀態	回傳字串
Idle	"idle"
Working	"working"
Vacation	"vacation"
Resigned	"resigned"
其他	"unknown"
2. 建立 struct

請建立 Employee struct，欄位如下：

Name   string
Age    int
Salary int
Status Status
3. 建立 function

請寫一個 function：

func NewEmployee(name string, age int, salary int, status Status) Employee

功能：回傳一個新的 Employee。

4. 寫一般 function：值傳遞

請寫一個 function：

func RaiseSalaryByValue(e Employee, amount int)

功能：讓 e.Salary 增加 amount。

但注意：
這個 function 使用 值傳遞，所以不應該改到原本的員工資料。

5. 寫一般 function：指針傳遞

請寫一個 function：

func RaiseSalaryByPointer(e *Employee, amount int)

功能：讓原本的 Employee.Salary 增加 amount。

6. 寫 method：value receiver

請幫 Employee 寫一個 method：

func (e Employee) FakeChangeStatus(status Status)

功能：修改 e.Status。

但注意：
這個 method 使用 value receiver，所以不應該改到原本的員工狀態。

7. 寫 method：pointer receiver

請幫 Employee 寫一個 method：

func (e *Employee) RealChangeStatus(status Status)

功能：真的修改原本的 Employee.Status。

8. 寫不定長參數 function

請寫一個 function：

func TotalSalary(employees ...Employee) int

功能：計算多位員工的總薪資。

例如：

TotalSalary(e1, e2, e3)

回傳三個人的 Salary 總和。

9. 寫多回傳值 function

請寫一個 function：

func DivideSalary(salary int, months int) (int, bool)

功能：計算每個月平均薪資。

需求：

如果 months == 0：

return 0, false

否則：

return salary / months, true
10. main 裡面測試

請在 main() 裡完成以下事情：

建立一個員工：
emp := NewEmployee("Leon", 30, 50000, Idle)
印出初始資料。
呼叫：
RaiseSalaryByValue(emp, 5000)

再印出 emp.Salary，觀察有沒有改變。

呼叫：
RaiseSalaryByPointer(&emp, 5000)

再印出 emp.Salary，觀察有沒有改變。

呼叫：
emp.FakeChangeStatus(Working)

再印出 emp.Status.String()。

呼叫：
emp.RealChangeStatus(Working)

再印出 emp.Status.String()。

建立另外兩個員工，計算三個人的總薪資。
使用 DivideSalary(emp.Salary, 12)，並用 if 判斷 ok 是否為 true。
*/
package main

type Employee struct {
}
