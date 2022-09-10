package main

import "fmt"

func kadanes(array []int) int { // max sub sum
	var currSum = 0
	var maxSum = -1 // min int
	for i := 0; i < len(array); i++ {
		currSum += array[i]

		if currSum > maxSum {
			maxSum = currSum
		}
		if currSum < 0 {
			currSum = 0
		}
	}
	return maxSum
}
func main() {
	var array = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 3, 1, 5, 1}
	fmt.Println(kadanes(array))
}
