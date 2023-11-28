package main

import (
	"fmt"
	"math"
)

func main() {
	var i uint64
	fmt.Print("Enter the number:")
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		fmt.Println("Error: Please enter a valid integer.")
		return
	}
	if isArmstrong(i) {
		fmt.Println(i, "is an Armstrong number.")
	} else {
		fmt.Println(i, "is not an Armstrong number.")
	}

}

func isArmstrong(i uint64) bool {
	var sumofCubes float64
	var value uint64
	value = i
	for value > 0 {
		digit := value % 10
		cubes := math.Pow(float64(digit),3)
		sumofCubes += cubes
		value /= 10
	}
	return i == uint64(sumofCubes)
}
