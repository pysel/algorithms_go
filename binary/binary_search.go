package main

// Binary Search is a searching algorithm used in a sorted array
//  by repeatedly dividing the search interval in half.
//   The idea of binary search is to use the information that
//    the array is sorted and reduce the time complexity
//     to O(Log n).

import "fmt"

//"fmt"

func binary_search(x, l, r int, arr []int) int {
	var mid int = l + (r-l)/2

	if r >= l {
		if arr[mid] == x {
			return mid
		} else if x < arr[mid] {
			return binary_search(x, l, mid-1, arr)
		} else {
			return binary_search(x, mid+1, r, arr)
		}
	}

	return -1

}

func main() {
	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(binary_search(2, 0, len(arr)-1, arr))
}
