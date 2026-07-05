package main

import "fmt"

//較正式的寫法
type Status int

const (
	Idle Status = iota
	Running
	Stopped
	Error
)

func (s Status) Status() string {
	switch s {
	case Idle:
		return "idle"
	case Running:
		return "running"
	case Stopped:
		return "stopped"
	case Error:
		return "Error"
	default:
		return "unknown"
	}
}

func main5() {
	//常數宣告也可以自己宣告型別或讓GO自動判別
	const age int = 30
	const height = 175.5
	const name = "Leon"
	//多個const可寫成一組
	const (
		StatusOK            = 200
		StatusNotFound      = 404
		StatusInternalError = 500
	)
	//Go沒有enum,通常用iota產生連續數字
	//iota 可以理解成在 const (...) 裡面，每一行自動 +1 的計數器
	const (
		sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println(Monday)

	status := Running
	println((status.Status()))

}

/*
常量 = const
枚舉 = 用 const + iota 模擬
*/
