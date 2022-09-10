package main

import "fmt"

func ternary_search(x, l, r int, arr []int) int {

	if r >= l {
		var mid1 = l + (r-l)/3
		var mid2 = r - (r-l)/3
		if x == arr[mid1] {
			return mid1
		}
		if x == arr[mid2] {
			return mid2
		}

		if x < arr[mid1] {
			return ternary_search(x, l, mid1-1, arr)
		} else if x > arr[mid2] {
			return ternary_search(x, mid2+1, r, arr)
		} else {
			return ternary_search(x, mid1+1, mid2-1, arr)
		}
	}
	return -1
}

func main() {
	var arr = []int{1, 3, 4, 5, 6, 7, 8}
	fmt.Println(ternary_search(11, 0, len(arr)-1, arr))
}
