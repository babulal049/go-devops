package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i uint64
	fmt.Print("Enter the number:")
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		fmt.Println("Error: Please enter a valid integer.")
		return
	}
	numberStr := strconv.FormatUint(i, 10)
	length := len(numberStr)
	fmt.Println("Length of entered integer", length)
	var sumofCubes uint64
	var value uint64
	value = i
	for value > 0 {
		digit := value % 10
		cubes := digit * digit * digit
		sumofCubes += cubes
		value /= 10
	}
	fmt.Println("Armstrong value of entered number:", sumofCubes)
	if i == sumofCubes {
		fmt.Println("Entered number is Armstrong number")
	} else {
		fmt.Println("Entered number is not Armstrong number")
	}
}
