package main

import (
	"fmt"
	"test_project/dates"
)

func main() {
	days := 3
	fmt.Println("Ur appointment is in", days, "days")
	fmt.Println("with a follow-up in", days+dates.DaysInWeek, "days")
}
