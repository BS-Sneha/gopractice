package main

import "fmt"

type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	empId    int
	basicpay int
	pf       int
}

type Contract struct {
	empId    int
	basicpay int
}

func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

func (c Contract) CalculateSalary() int {
	return c.basicpay
}

func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}

	fmt.Printf("Total Expenses %d", expense)
}

func main() {
	emp1 := Permanent{
		empId:    123,
		basicpay: 8000,
		pf:       1200,
	}
	emp2 := Permanent{
		empId:    456,
		basicpay: 9000,
		pf:       1110,
	}

	emp3 := Contract{
		empId:    369,
		basicpay: 4850,
	}
	employees := []SalaryCalculator{emp1, emp2, emp3}
	totalExpense(employees)
}
