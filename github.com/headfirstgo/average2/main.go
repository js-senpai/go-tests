package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func average(numbers ...float64) float64 {
	var sum float64 = 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func main() {
	arguments := os.Args[1:]
	var numbers []float64
	for _, argumenet := range arguments {
		number, err := strconv.ParseFloat(argumenet, 64)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}
	// sampleCount := float64(len(arguments))
	// fmt.Printf("Average: %0.2f\n", sum/float64(sampleCount))
	fmt.Printf("Average: %0.2f\n", average(numbers...))
}
