package tasks

import (
	"fmt"
	"strconv"
)

func Solution2() {

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

	age := "123abc"
	result := age != ""
	fmt.Println(result)

	flag := 1
	var boolFlag bool
	if flag == 1 {
		boolFlag = true
	} else {
		boolFlag = false
	}
	fmt.Println(boolFlag)

	str_one := "Privet"
	str_two := ""

	// fmt.Println(toBool(str_one))
	// fmt.Println(toBool(str_two))
	toBool(str_one)
	toBool(str_two)

	a := 1
	b := 0
	fmt.Println(toBoolFrom_0_1(a))
	fmt.Println(toBoolFrom_0_1(b))
	str := strconv.FormatBool(toBoolFrom_0_1(b))
	fmt.Printf("From Boolean to string: %s, check type: %T", str, str)
}

func toBoolFrom_0_1(num int) bool {
	var booleanType bool
	switch num {
	case 1:
		booleanType = true
	case 0:
		booleanType = false
	default:
		fmt.Printf("Unknow selector: %d", num)
	}
	return booleanType
}

// func toBool(str string) bool {
// 	if str != "" {
// 		return true
// 	} else {
// 		return false
// 	}
// }

func toBool(str string) {
	if str != "" {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
