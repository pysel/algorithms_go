package main

import "fmt"

func fibonacci_search(x int, arr []int) int {
	var fib_seq = []int{0, 1}
	var i = 0
	for i < len(arr) {
		i = calc_next_fib(fib_seq[len(fib_seq)-1], fib_seq[len(fib_seq)-2])
		fib_seq = append(fib_seq, i)
	}
	fmt.Println(fib_seq)
	return -1
}

func calc_next_fib(pr, cur int) int {
	fmt.Println(pr, cur)
	return pr + cur
}
func main() {
	var arr = []int{1, 3, 4, 5, 6, 7, 8, 10, 13, 16, 19, 20, 21}
	fmt.Println(fibonacci_search(9, arr))
}
