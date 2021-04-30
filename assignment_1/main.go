package main

import "fmt"

func main() {
	numbers := newList()

	for i, n := range numbers {
		if n % 2 == 0 {
			fmt.Println(i," is even")
		} else {
		fmt.Println(i, " is odd")
		}
	}

}

func newList() []int {
	nums := make([]int, 11)
	for i := range nums {
		nums[i] = i
	}

	return nums
}