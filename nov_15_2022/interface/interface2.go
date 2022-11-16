package main

import "fmt"

type salaryCalculator interface {
	calculateSalary() int
}

type Permanent struct {
	employeId int
	basicPay  int
	pf        int
}

type Contract struct {
	employeId int
	basicPay  int
}

func (p Permanent) calculateSalary() int {
	return p.basicPay + p.pf
}

func (c Contract) calculateSalary() int {
	return c.basicPay
}

func totalExpense(s []salaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.calculateSalary()
	}
	fmt.Printf("Total Expense Per Month $%d", expense)
}

func main() {
	pemp1 := Permanent{
		employeId: 1,
		basicPay:  5000,
		pf:        20,
	}

	pemp2 := Permanent{
		employeId: 2,
		basicPay:  6000,
		pf:        30,
	}
	cemp1 := Contract{
		employeId: 3,
		basicPay:  3000,
	}
	employees := []salaryCalculator{pemp1, pemp2, cemp1}
	fmt.Println(employees)
}
