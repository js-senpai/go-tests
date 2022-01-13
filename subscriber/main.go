package main

import (
	"fmt"
	"test_project/github.com/headfirstgo/magazine"
)

func printInfo(s *magazine.Subscriber) {
	fmt.Println("Name:", s.Name)
	fmt.Println("Monthly rate:", s.Rate)
	fmt.Println("Active?", s.Active)
}
func defaultSubscriber(name string) *magazine.Subscriber {
	s := magazine.Subscriber{Name: name, Rate: 5.99, Active: true}
	return &s
}
func applyDiscount(s *magazine.Subscriber) {
	s.Rate = 4.99
}
func main() {
	employee := magazine.Employee{Name: "Joy Carr", Salary: 6000}
	employee.Street = "123 Oak St"
	employee.City = "Omaha"
	employee.State = "NE"
	employee.PostalCode = "68111"
	var s magazine.Subscriber
	applyDiscount(&s)
	fmt.Println(s.Rate)
	fmt.Println(employee.Name)
	fmt.Println(employee.Salary)
}
