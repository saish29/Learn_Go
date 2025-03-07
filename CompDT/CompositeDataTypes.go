package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 5, 4}
	slice := make([]int, 4, 5)

	nums[0] = 1
	nums[1] = 2
	nums[3] = 3

	dict := map[int]string{
		1: "Lost",
		2: "Echo",
		3: "Linkin",
	}

	slice = append(slice, 22)
	nums = append(nums, 3)
	fmt.Println(slice, nums)
	fmt.Println(dict)
	fmt.Println(dict[1])

	if val, exists := dict[4]; exists {
		fmt.Println(val)
	} else {
		fmt.Println("Not Found")
	}

	//x := 2
	if x := 22; x > 5 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

	// Iterating over a dict
	for key, value := range dict {
		fmt.Printf("Key is %d and value is %s\n", key, value)
	}

	// Iterating oer an array
	for index, val := range nums {
		fmt.Printf("Index is %d and value is %d\n", index, val)
	}

	// Traditional for loop - C type like

	for i := 1; i <= 5; i++ {

		fmt.Printf("Value is %d\n", i)
	}

	// No whle loop in go - use traditional for loop
	j := 1

	for j <= 10 {
		fmt.Printf("value of j is %d\n", j)
		j *= 2
	}

	// Infinite loop iteration

	k := 1

	for {
		fmt.Printf("Current Val : %d\n", k)
		k++

		if k >= 15 {
			break
		}
	}

}
