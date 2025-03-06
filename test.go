package main

import "fmt"

func main() {

	var floatNum int = 32767
	var intNum int = 221
	// var res float32 = floatNum + float32(intNum)
	fmt.Println(floatNum % intNum)

	var str string = `Hello`
	fmt.Println(len(str))

	var bool_val bool = false
	fmt.Println((bool_val))

	txt1, txt2 := 10.3332, 45.33
	fmt.Println(txt1, txt2)
	fmt.Println(txt2)

	var const_val = 3424
	const_val = const_val + 1
	fmt.Println(const_val + 1)
	// printMe("Hell no")
	res2, res3 := intDiv(8, 3)
	fmt.Println(res2, res3)
}

func printMe(print_val string) {
	fmt.Println(print_val)
}

func intDiv(num int, denom int) (int, int) {
	return num / denom, num % denom
}
