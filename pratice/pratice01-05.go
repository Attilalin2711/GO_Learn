/*
練習題：員工狀態管理系統
*/
package main

type Status int

const (
	Idle Status = iota
	Working
	Vacation
	Resigned
)

func (s Status) String() string {
	switch s {
	case Idle:
		return "Idle"
	case Working:
		return "Working"
	case Vacation:
		return "Vacation"
	case Resigned:
		return "Resigned"
	default:
		return "Unknown"
	}
}

const pi = 3.1415

type Empolyee struct {
	Name   string
	Age    int
	Salary int
	Status Status
}

func NewEmployee(name string, age int, Salary int, status Status) Empolyee {
	return Empolyee{Name: name, Age: age, Salary: Salary, Status: status}
}

func RaiseSalaryByValue(e Empolyee, amount int) {
	e.Salary += amount
}

func RaiseSalaryByPointer(e *Empolyee, amount int) {
	e.Salary += amount
}

func (e Empolyee) FakeChangeStatus(status Status) {
	e.Status = status
}

func (e *Empolyee) RealChangeStatus(status Status) {
	e.Status = status
}

func TotalSalary(employees ...Empolyee) int {
	total := 0
	for _, n := range employees {
		total += n.Salary
	}
	return total
}

func DividSalarty(salary int, months int) (int, bool) {
	if months == 0 {
		return 0, false
	}
	return salary / months, true
}

func main1() {

}
