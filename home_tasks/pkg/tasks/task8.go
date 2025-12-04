package tasks

import "fmt"

func SolutionTasks8() {
	// var a int64 = 1221

	// div := a / 100
	// fmt.Printf("div: %d\n", div) // 12
	// remainder := (a / 100) % 10
	// fmt.Printf("rem: %d\n", remainder) // 2

	var n int64

	for {
		fmt.Print("Введите число: ")
		_, err := fmt.Scan(&n)
		if err != nil || n < 1000 || n > 9999 {
			fmt.Println("Ошибка: введите четырехзначное число")
			continue
		}
		break
	}

	digits := [4]int64{
		n / 1000,
		(n / 100) % 10,
		(n / 10) % 10,
		n % 10,
	}
	fmt.Println(digits) //[1 2 2 1]

	if digits[0] == digits[3] && digits[1] == digits[2] {
		fmt.Printf("Введенное число палиндром: %t\n", true)
	} else {
		fmt.Printf("Введенное число не палиндром: %t\n", false)
	}
}
