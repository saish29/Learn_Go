package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 5, 4}
	slice := make([]int, 4, 5)

	nums[0] = 1
	nums[1] = 2
	nums[3] = 3

	slice = append(slice, 22)
	nums = append(nums, 3)
	fmt.Println(slice, nums)

}
