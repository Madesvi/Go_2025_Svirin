package task9

import "fmt"

func Solution() {
	// Уравнение
	// A * x + B = 0
	// x = - (B / A)
	// A > 0
	var a int64
	var b int64
	var x int64

	// Цикл ввода с проверкой
	for {
		fmt.Print("Введите коэффициент A больше 0: ")
		_, err := fmt.Scan(&a)
		if err != nil || a <= 0 {
			fmt.Println("Ошибка: Введите А больше 0")
			continue
		}
		fmt.Print("Введите коэффициент B: ")
		_, err = fmt.Scan(&b)
		if err != nil {
			fmt.Printf("Ошибка: %s\n", err)
			continue
		}
		break
	}
	fmt.Printf("Введенные значения: A = %d, B = %d\n", a, b)

	x = -(b / a)
	fmt.Printf("x = %d\n", x)

	// Проверка
	check := (a*x)+b == 0
	fmt.Println("=== Проверка ===")
	if check {
		fmt.Println("Уравнение решено верно")
	} else {
		fmt.Println("Уравнение решено не верно")
	}

}
