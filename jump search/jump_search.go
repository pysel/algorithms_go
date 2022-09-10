package main

import (
	"fmt"
	"math"
)

func jump_search(x int, arr []int) int {
	var jump = int(math.Sqrt(float64(x)))
	for i := jump; i < len(arr); {
		if arr[i] > x {
			for j := i - jump; j < i; j++ {
				if arr[j] == x {
					return j
				}
			}
		}
		if i+jump > len(arr) {
			for j := i; j < len(arr); j++ {
				if arr[j] == x {
					return j
				}
			}
		}
		i += jump
	}
	return -1
}
func main() {
	var arr = []int{1, 3, 4, 5, 6, 7, 8, 10, 13, 16, 19, 20, 21}
	fmt.Println(jump_search(9, arr))
}
