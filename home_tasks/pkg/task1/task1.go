package task1

import "fmt"

func Solution() {
	//=== task 1 ===
	var num int32
	var float_num float64
	var boolean bool
	var str string
	var empty struct{}

	var ptr *int32
	ptr = &num
	var ptr2 *float64
	ptr2 = &float_num
	var ptr3 *bool
	ptr3 = &boolean
	var ptr4 *string
	ptr4 = &str
	var ptr5 *struct{}
	ptr5 = &empty

	fmt.Printf("Type of variable: %T\n", ptr)
	fmt.Printf("Type of variable: %T\n", ptr2)
	fmt.Printf("Type of variable: %T\n", ptr3)
	fmt.Printf("Type of variable: %T\n", ptr4)
	fmt.Printf("Type of variable: %T\n", ptr5)
}
