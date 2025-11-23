package main

import "fmt"

func main() {

	// === task 1 ===
	// var num int32
	// var float_num float64
	// var boolean bool
	// var str string
	// var empty struct{}

	// var ptr *int32
	// ptr = &num
	// var ptr2 *float64
	// ptr2 = &float_num
	// var ptr3 *bool
	// ptr3 = &boolean
	// var ptr4 *string
	// ptr4 = &str
	// var ptr5 *struct{}
	// ptr5 = &empty

	// fmt.Printf("Type of variable: %T\n", ptr)
	// fmt.Printf("Type of variable: %T\n", ptr2)
	// fmt.Printf("Type of variable: %T\n", ptr3)
	// fmt.Printf("Type of variable: %T\n", ptr4)
	// fmt.Printf("Type of variable: %T\n", ptr5)

	// === taks 2 ===
	// age := "23"
	// foo := "23abs"

	// num, err := strconv.Atoi(age)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(num)
	// fmt.Println(num + 1)

	// var onlyNum string
	// for _, char := range foo {
	// 	if !unicode.IsDigit(char) {
	// 		break
	// 	}
	// 	onlyNum += string(char)
	// }
	// num2, err := strconv.Atoi(onlyNum)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("Only digit from foo %d, with type: %T", num2, num2)

	// age := "123abc"
	// result := age != ""
	// fmt.Println(result)

	// flag := 1
	// var boolFlag bool
	// if flag == 1 {
	// 	boolFlag = true
	// } else {
	// 	boolFlag = false
	// }
	// fmt.Println(boolFlag)

	// str_one := "Privet"
	// str_two := ""

	// fmt.Println(toBool(str_one))
	// fmt.Println(toBool(str_two))
	// toBool(str_one)
	// toBool(str_two)

	// a := 1
	// b := 0
	// fmt.Println(toBoolFrom_0_1(a))
	// fmt.Println(toBoolFrom_0_1(b))
	// str := strconv.FormatBool(toBoolFrom_0_1(b))
	// fmt.Printf("From Boolean to string: %s, check type: %T", str, str)

	// === task 3 ===
	age := 36.6
	temperature := 25

	// temp := age
	// age = float64(temperature)
	// temperature = int(temp)

	age, temperature = float64(temperature), int(age)

	fmt.Printf("age is: %v, temperature is: %v\n", age, temperature)

	// str_age := fmt.Sprintf("%v", age)
	// fmt.Printf("String is: %s, with type: %T\n", str_age, str_age)
	// str_temperature := fmt.Sprintf("%v", temperature)
	// fmt.Printf("String is: %s, with type: %T\n", str_temperature, str_temperature)

	// var tmp string
	// tmp = str_age
	// // fmt.Println(tmp)
	// str_age = str_temperature
	// str_temperature = tmp

	// fmt.Println(str_age)
	// fmt.Println(str_temperature)

	// age, err := strconv.ParseFloat(str_age, 64)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("New var age is: %v\n", age)

	// // ===== Error below =====
	// temperature, err = strconv.Atoi(str_temperature)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("New var temp is: %v\n", temperature)

}

// func toBoolFrom_0_1(num int) bool {
// 	var booleanType bool
// 	switch num {
// 	case 1:
// 		booleanType = true
// 	case 0:
// 		booleanType = false
// 	default:
// 		fmt.Printf("Unknow selector: %d", num)
// 	}
// 	return booleanType
// }

// func toBool(str string) bool {
// 	if str != "" {
// 		return true
// 	} else {
// 		return false
// 	}
// }

// func toBool(str string) {
// 	if str != "" {
// 		fmt.Println(true)
// 	} else {
// 		fmt.Println(false)
// 	}
// }
