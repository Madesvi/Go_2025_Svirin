package task5

import (
	"fmt"
	"math"
)

func Solution() {
	a, b := 0, 0
	fmt.Print("Введите первое число:")
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Print("Введите второе число:")
	_, err = fmt.Scan(&b)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	// fmt.Println(a, b)
	//	Сумма
	sum := a + b
	fmt.Printf("Сумма введенных чисел: %d + %d = %d\n", a, b, sum)
	//	Разность
	diff := a - b
	fmt.Printf("Разность введенных чисел: %d - %d = %d\n", a, b, diff)
	//	Частное
	div := float64(a) / float64(b)
	fmt.Printf("Частное введенных чисел: %d / %d = %.2f\n", a, b, div)
	// Произведение
	mult := a * b
	fmt.Printf("Произведение введенных чисел: %d * %d = %d\n", a, b, mult)
	//	Целочисленное деление
	div2 := a / b
	fmt.Printf("Целочисленное деление введенных чисел: %d / %d = %d\n", a, b, div2)
	//	С остатком
	remainder := a % b
	fmt.Printf("Деление с остатком введенных чисел: %d / %d = %d\n", a, b, remainder)
	//	Возведение в степень
	pow := math.Pow(float64(a), float64(b))
	fmt.Printf("Возведение в степень введенных чисел: %d в степени %d = %.2f\n", a, b, pow)

}
