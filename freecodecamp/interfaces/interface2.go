package main

import "fmt"

type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() int {
	return c.hourlyPay * c.hoursPerYear
}

type fullTime struct {
	name   string
	salary int
}

func (ft fullTime) getSalary() int {
	return ft.salary
}

func (ft fullTime) getName() string {
	return ft.name
}

func main() {
	ft := fullTime{name: "Bob", salary: 7300}
	ct := contractor{name: "Jill", hourlyPay: 872, hoursPerYear: 982}
	fmt.Println(employee(ft).getName())
	fmt.Println(employee(ft).getSalary())
	fmt.Println("============================")
	fmt.Println(employee(ct).getName())
	fmt.Println(employee(ct).getSalary())
}
